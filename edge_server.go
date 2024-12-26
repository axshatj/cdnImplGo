package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type EdgeServer struct {
	cache       *LRUCache
	originURL   string
	httpClient  *http.Client
}

func NewEdgeServer(originURL string) *EdgeServer {
	return &EdgeServer{
		cache:      NewLRUCache(100),
		originURL:  originURL,
		httpClient: &http.Client{},
	}
}

func (s *EdgeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if content, ok := s.cache.Get(path); ok {
		fmt.Printf("Serving %s from edge cache\n", path)
		w.Write([]byte(content))
		return
	}

	resp, err := s.httpClient.Get(s.originURL + path)
	if err != nil {
		http.Error(w, "Error fetching from origin", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading origin response", http.StatusInternalServerError)
		return
	}

	s.cache.Put(path, string(content))
	fmt.Printf("Caching %s on edge server\n", path)
	w.Write(content)
}

func (s *EdgeServer) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}

