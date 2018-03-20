package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func nodes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("nodes")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{
		"nodes": {
		  "CVp5QJV9RcayNf4ivywb5w": {
			"ip": "172.17.0.2",
			"version": "6.2.2",
			"http": {
			  "publish_address": "172.17.0.2:9200"
			}
		  }
		}
	  }`)
}
func nodesLocal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("nodes_local")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{}`)
}

func slash(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{
		"name": "CVp5QJV",
		"cluster_name": "docker-cluster",
		"cluster_uuid": "dYsZenEOT_CNHasr1ch4IQ",
		"version": {
		  "number": "6.2.2",
		  "build_hash": "10b1edd",
		  "build_date": "2018-02-16T19:01:30.685723Z",
		  "build_snapshot": false,
		  "lucene_version": "7.2.1",
		  "minimum_wire_compatibility_version": "5.6.0",
		  "minimum_index_compatibility_version": "5.0.0"
		},
		"tagline": "You Know, for Search"
	  }`)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.Write(os.Stdout)

	// r.ParseForm()       // parse arguments, you have to call this by yourself
	// fmt.Println(r.Form) // print form information in server side
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	//fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/", slash)
	http.HandleFunc("/_nodes", nodes)
	http.HandleFunc("/_nodes/_local", nodesLocal)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
