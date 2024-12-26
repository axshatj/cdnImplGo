package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type LoadBalancer struct {
	consistentHash *ConsistentHash
	httpClient     *http.Client
}

func NewLoadBalancer(edgeServers []string) *LoadBalancer {
	return &LoadBalancer{
		consistentHash: NewConsistentHash(edgeServers, 100),
		httpClient:     &http.Client{},
	}
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	server := lb.consistentHash.GetNode(path)

	if server == "" {
		http.Error(w, "No available edge servers", http.StatusServiceUnavailable)
		return
	}

	fmt.Printf("Routing request for %s to %s\n", path, server)
	resp, err := lb.httpClient.Get(server + path)
	if err != nil {
		http.Error(w, "Error fetching from edge server", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading edge server response", http.StatusInternalServerError)
		return
	}

	w.Write(content)
}

func (lb *LoadBalancer) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, lb)
}

