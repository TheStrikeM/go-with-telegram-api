package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	token = "6450711663:AAHI8yujft5FHzKmLqSG6g68S3jvs3N4OoM"
	host  = "api.telegram.org"
)

type Response struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func parseJson(text string) (*Response, error) {
	var parsedJson Response
	if err := json.Unmarshal([]byte(text), &parsedJson); err != nil {
		return nil, err
	}
	return &parsedJson, nil
}

func main() {
	//u := url.URL{
	//	Scheme: "https",
	//	Host:   host,
	//	Path:   path.Join("bot"+token, "getMe"),
	//}
	//fmt.Println(u.String())
	//
	//req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//client := http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func() { _ = resp.Body.Close() }()
	//
	//body, _ := io.ReadAll(resp.Body)
	//fmt.Printf(string(body) + "\n")
	//
	//result, _ := parseJson(string(body))
	//fmt.Printf(result.Result.Username + "\n")
	for {
		time.Sleep(2 * time.Second)
		u := url.URL{
			Scheme: "https",
			Host:   host,
			Path:   path.Join("bot"+token, "getUpdates"),
		}
		req, _ := http.NewRequest(http.MethodGet, u.String(), nil)

		client := http.Client{}
		resp, _ := client.Do(req)
		defer func() { _ = resp.Body.Close() }()

		body, _ := io.ReadAll(resp.Body)
		fmt.Printf(string(body) + "\n")
	}
}
