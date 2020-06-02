package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/OllyCat/ariago"
)

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
	uParced, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(uParced.Scheme, "http") {
		return "", errors.New("Invalid URL")
	}

	return strings.Trim(uParced.Path, `/`), nil
}

func get(id string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get("https://embed.redtube.com/?id=" + id)
	if err != nil {
		log.Printf("ERROR: id %s get error\n", id)
		return
	}

	flashReg := regexp.MustCompile(`flashvars_vid\d* = ({.*});`)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERROR: id %s body get error\n", id)
		return
	}

	res.Body.Close()

	flashRes := flashReg.FindSubmatch(body)
	if len(flashRes) < 2 {
		log.Printf("ERROR: id %s flashvars not found\n", id)
		return
	}

	js := JsonScript{}
	json.Unmarshal(flashRes[1], &js)

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
