package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://service-back:8080")
		if err != nil {
			fmt.Fprintf(w,"%s", err.Error())
		}else  {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(w,"%s", err.Error())
			}else {
				fmt.Fprintf(w, "Response from service back:'%s'",body)
			}
		}

	})

	http.ListenAndServe(":8081", nil)
}
