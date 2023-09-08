package main

import (
	"fmt"
	"gocache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Mutian": "118",
	"Jak":    "5189",
	"Sam":    "5267",
}

func main() {
	gocache.NewGroup("scores", 2<<10, gocache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := gocache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
