package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/proxy"
)

const api = "https://api.ipify.org/?format=json"

func test() {
	sox, err := proxy.SOCKS5("tcp", "172.105.33.85:7497", nil, proxy.Direct)
	if err != nil {
		log.Fatalf("Failed to get dial %+v", err)
	}

	client := &http.Client{
		Transport: &http.Transport{Dial: sox.Dial},
	}

	resp, err := client.Get(api)
	if err != nil {
		log.Fatalf("Call to %s failed with %+v", api, err)
	}


	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body %+v", err)
	}

	body := string(bodyBytes)

	log.Println(body)
} 
