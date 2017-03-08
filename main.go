package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/romainmenke/simple-example/public"
)

func main() {

	http.Handle("/",
		http.HandlerFunc(handler),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Listening on port : " + port)
	err := http.ListenAndServe(":"+port, nil)
	log.Println(err)

}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Expose-Headers", "Content-Type, Content-Length, Content-Encoding, Vary")

	path := r.URL.Path
	if path == "/" || path == "" {
		path = "/index"
	}

	if strings.Contains(path, "gzip") {
		w.Header().Add("Vary", "Accept-Encoding")
		w.Header().Set("Content-Encoding", "gzip")
	}

	if strings.Contains(path, "/assets") {
		http.FileServer(
			&assetfs.AssetFS{Asset: public.Asset, AssetDir: public.AssetDir, AssetInfo: public.AssetInfo, Prefix: ""},
		).ServeHTTP(w, r)
		return
	}

	path = "assets/html/gzip" + path + ".min.html"
	w.Header().Add("Vary", "Accept-Encoding")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	page, err := public.Asset(path)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(page)

}
