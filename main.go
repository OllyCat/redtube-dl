package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/OllyCat/ariago"
	"github.com/valyala/fasthttp"
)

var (
	r      *regexp.Regexp
	fr     *regexp.Regexp
	client *fasthttp.Client
)

func init() {
	r = regexp.MustCompile(`meta name="twitter:player" content="(https://embed.redtube.com/\?id=\d+)"`)
	fr = regexp.MustCompile(`flashvars_vid\d* = ({.*});`)
	client = &fasthttp.Client{}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		url := scanner.Text()
		id, err := parce(url)
		if err != nil {
			continue
		}
		wg.Add(1)
		go get(id, &wg)
	}
	wg.Wait()
}

func parce(u string) (string, error) {
	up, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(up.Scheme, "http") {
		return "", errors.New("Invalid URL")
	}

	_, body, err := client.Get(nil, u)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not get url: %v\n", err))
	}

	res := r.FindSubmatch(body)

	if len(res) < 2 {
		return "", fmt.Errorf("Could not find player url: %w\n", err)
	}

	return string(res[1]), nil
}

func get(id string, wg *sync.WaitGroup) {
	defer wg.Done()

	var fres [][]byte

	// 10 повторов для защиты от скрипта "Loading..."
	var i int
	for {
		_, body, err := client.Get(nil, id)
		if err != nil {
			log.Printf("ERROR: id %s get error\n", id)
			return
		}

		// если патерн найдет - прерываем цыкл
		fres = fr.FindSubmatch(body)
		if len(fres) > 1 {
			break
		}
		// если нет - случайная задержка и повтор
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		i++
		// если повторов было больше 10, то прерываемся с ошибкой
		if i >= 10 {
			log.Printf("ERROR: id %s flashvars not found.\n", id)
			return
		}
	}

	js := JsonScript{}
	json.Unmarshal(fres[1], &js)

	qMax := 0
	var vUrl string

	for _, m := range js.MediaDefinitions {
		q, err := strconv.Atoi(m.Quality)
		if err != nil {
			continue
		}

		if q > qMax {
			qMax = q
			vUrl = m.VideoURL
		}
	}

	log.Printf("Quality: %d\nVideo URL: %s\nTitle: %s\n", qMax, vUrl, js.VideoTitle)

	ariago.Aria(vUrl, js.VideoTitle+".mp4")
	return
}
