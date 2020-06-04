package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
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
	up, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(up.Scheme, "http") {
		return "", errors.New("Invalid URL")
	}

	resp, err := http.Get(u)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not get url: %v\n", err))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not get url: %v\n", err))
	}

	r := regexp.MustCompile(`meta name="twitter:player" content="(https://embed.redtube.com/\?id=\d+)"`)
	res := r.FindSubmatch(body)

	if len(res) < 2 {
		return "", errors.New(fmt.Sprintf("Could not find player url: %v\n", err))
	}

	return string(res[1]), nil
}

func get(id string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get(id)
	if err != nil {
		log.Printf("ERROR: id %s get error\n", id)
		return
	}
	defer res.Body.Close()

	fr := regexp.MustCompile(`flashvars_vid\d* = ({.*});`)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERROR: id %s body get error\n", id)
		return
	}

	fres := fr.FindSubmatch(body)
	if len(fres) < 2 {
		log.Printf("ERROR: id %s flashvars not found\n", id)
		return
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
