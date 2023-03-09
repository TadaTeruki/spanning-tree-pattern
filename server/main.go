package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, http.FileServer(http.Dir("./")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
