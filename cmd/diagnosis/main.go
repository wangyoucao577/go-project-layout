package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wangyoucao577/go-project-layout/diagnosis"
)

func main() {

	mux := http.NewServeMux()

	//Sample Request: http://localhost:8000/diagnosis?diagnosis=ping
	mux.HandleFunc("/diagnosis", func(w http.ResponseWriter, req *http.Request) {
		item := req.URL.Query().Get("diagnosis")
		if item == "ping" {
			diagnosisInfo := diagnosis.New()
			diagnosisInfo.RemoteAddr = req.RemoteAddr // dynamic update for each request
			fmt.Fprintf(w, "%s", diagnosisInfo)
			return
		}

		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no diagnosis command %s", req.URL)
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}
