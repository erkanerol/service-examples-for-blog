package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://service-middle:8081")
		if err != nil {
			fmt.Fprintf(w,"%s", err.Error())
		}else  {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(w,"%s", err.Error())
			}else {
				fmt.Fprintf(w, "Response from service middle:'%s'",body)
			}
		}

	})

	http.ListenAndServe(":8082", nil)
}
