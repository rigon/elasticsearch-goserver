package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func slash(w http.ResponseWriter, r *http.Request) {
	r.Write(os.Stdout)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
func nodes(w http.ResponseWriter, r *http.Request) {
	r.Write(os.Stdout)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, `{"nodes":{"GpFynDbQST-OLBNmIlRMQA":{"ip":"172.23.0.2","version":"6.2.3","http":{"publish_address":"172.23.0.2:9200"}}}}`)
}
func nodesLocal(w http.ResponseWriter, r *http.Request) {
	r.Write(os.Stdout)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, `{}`)
}
func kibanaMappings(w http.ResponseWriter, r *http.Request) {
	r.Write(os.Stdout)
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, `{"error":{"root_cause":[{"type":"index_not_found_exception","reason":"no such index","resource.type":"index_or_alias","resource.id":".kibana","index_uuid":"_na_","index":".kibana"}],"type":"index_not_found_exception","reason":"no such index","resource.type":"index_or_alias","resource.id":".kibana","index_uuid":"_na_","index":".kibana"},"status":404}`)
}
func kibanaConfig(w http.ResponseWriter, r *http.Request) {
	r.Write(os.Stdout)
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, `{"error":{"root_cause":[{"type":"index_not_found_exception","reason":"no such index","resource.type":"index_expression","resource.id":".kibana","index_uuid":"_na_","index":".kibana"}],"type":"index_not_found_exception","reason":"no such index","resource.type":"index_expression","resource.id":".kibana","index_uuid":"_na_","index":".kibana"},"status":404}`)
}
func kibanaSearch(w http.ResponseWriter, r *http.Request) {
	r.Write(os.Stdout)
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, `{"error":{"root_cause":[{"type":"index_not_found_exception","reason":"no such index","resource.type":"index_or_alias","resource.id":".kibana","index_uuid":"_na_","index":".kibana"}],"type":"index_not_found_exception","reason":"no such index","resource.type":"index_or_alias","resource.id":".kibana","index_uuid":"_na_","index":".kibana"},"status":404}`)
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
	http.HandleFunc("/.kibana/_mappings", kibanaMappings)
	http.HandleFunc("/.kibana/doc/config%3A6.2.3", kibanaConfig)
	http.HandleFunc("/.kibana/_search", kibanaConfig)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
