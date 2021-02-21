package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Timestamp from back:%d",time.Now().Unix())
	})

	http.ListenAndServe(":8080", nil)
}
