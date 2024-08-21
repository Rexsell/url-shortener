package main

import "url-shortener/internal/server"

func main() {
	server := server.New(":8080")
	if err := server.Start(); err != nil {
		return
	}
}
