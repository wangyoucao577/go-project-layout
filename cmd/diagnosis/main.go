package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/wangyoucao577/go-project-layout/diagnosis"
)

func main() {
	flag.Parse() // parse flags for glog
	defer glog.Flush()

	mux := http.NewServeMux()

	//Sample Request: http://localhost:8000/diagnosis?diagnosis=ping
	mux.HandleFunc("/diagnosis", func(w http.ResponseWriter, req *http.Request) {
		item := req.URL.Query().Get("diagnosis")
		if item == "ping" {
			if di := diagnosis.New(); di != nil {
				di.RemoteAddr = req.RemoteAddr // dynamic update for each request
				fmt.Fprintf(w, "%s", di)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no diagnosis command %s", req.URL)
	})

	glog.Info("Listen on :8000")
	glog.Fatal(http.ListenAndServe(":8000", mux))
}
