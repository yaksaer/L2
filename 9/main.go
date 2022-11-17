package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Input site")
		return
	}
	site := os.Args[len(os.Args)-1]

	Download(site)
}

func Download(site string) error {
	if len(site) < 8 {
		fmt.Println("Write http protocol")
		return errors.New("Need http protocol")
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	timeout := 5 * time.Second
	client := &http.Client{Transport: tr,
		Timeout: timeout}
	req, err := http.NewRequest("GET", site, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	req.Header.Add("Accept", "text/html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	nm := make([]rune, 0, 5)

	if site[7:8] == "/" {
		for _, v := range site[8:] {
			if string(v) == "/" {
				break
			}
			nm = append(nm, v)

		}

		os.Mkdir(string(nm), 0777)
	} else {
		for _, v := range site[7:] {
			if string(v) == "/" {
				break
			}
			nm = append(nm, v)

		}
		os.Mkdir(string(nm), 0777)
	}
	mainFile, err := os.Create(string(nm) + "/index.html")
	data, err := io.ReadAll(resp.Body)
	mainFile.Write(data)
	reg := regexp.MustCompile(`(http|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	result := reg.FindAll(data, -1)
	urls := make([]string, 0, len(result))
	for i := 0; i < len(result); i++ {
		urls = append(urls, string(result[i]))
	}

	for i, url := range urls {
		num := strconv.Itoa(i)
		downUrls(url, string(nm), num)

	}
	return nil
}

func downUrls(url, path string, i string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	timeout := 5 * time.Second
	client := &http.Client{Transport: tr,
		Timeout: timeout}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "text/html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	File, _ := os.Create(path + "/index" + i + ".html")
	data, _ := io.ReadAll(resp.Body)
	File.Write(data)
}
