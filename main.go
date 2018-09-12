package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/tfrench15/redis-project/proxy"
)

var (
	redisAddr = flag.String("redis", "localhost:6379", "address of the backing Redis")
	proxyAddr = flag.String("proxy", "localhost:8080", "HTTP address of the proxy")
	expiry    = flag.Int("expiry", 10, "time to live for a key in the cache")
	capacity  = flag.Int("capacity", 10, "size of the cache")
)

func main() {
	flag.Parse()

	redisClient, err := proxy.NewRedisClient(*redisAddr)
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	cache := proxy.NewCache(*capacity, time.Duration(*expiry)*time.Second)
	redisProxy := proxy.NewRedisProxy(cache, redisClient)
	err = http.ListenAndServe(*proxyAddr, redisProxy)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
