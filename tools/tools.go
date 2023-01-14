package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func FixSymbol(code string) string {
	return strings.Replace(code, "\"", "\\\"", -1)
}

func GenerateCommand(code, filetype string) string {
	return fmt.Sprintf("echo \"%s\" > file.%s && sh compiler.sh", code, filetype)
}
