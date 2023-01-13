package tools

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetFile(file string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", file, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyText)
}
