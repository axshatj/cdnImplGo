package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	cdnURL     string
	httpClient *http.Client
}

func NewClient(cdnURL string) *Client {
	return &Client{
		cdnURL:     cdnURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) Run() {
	files := []string{"image1.jpg", "image2.jpg", "text1.txt"}

	for _, file := range files {
		fmt.Printf("Requesting %s...\n", file)
		resp, err := c.httpClient.Get(c.cdnURL + "/" + file)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", file, err)
			continue
		}
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response for %s: %v\n", file, err)
			continue
		}

		fmt.Printf("Received content for %s: %s\n\n", file, string(content))
	}
}

