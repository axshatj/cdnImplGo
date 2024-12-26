package main

import (
	"fmt"
	"net/http"
)

type OriginServer struct {
	content map[string]string
}

func NewOriginServer() *OriginServer {
	return &OriginServer{
		content: map[string]string{
			"/image1.jpg": "This is image 1 content",
			"/image2.jpg": "This is image 2 content",
			"/text1.txt":  "This is text file 1 content",
		},
	}
}

func (s *OriginServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if content, ok := s.content[path]; ok {
		fmt.Printf("Serving %s from origin server\n", path)
		w.Write([]byte(content))
	} else {
		http.NotFound(w, r)
	}
}

func (s *OriginServer) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}

