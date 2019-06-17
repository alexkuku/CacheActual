package main

import (
	"CacheActual/cache"
	"CacheActual/controller"
)

func main() {
	c := cache.New("inmemory")
	controller.New(c).Listen()
}
