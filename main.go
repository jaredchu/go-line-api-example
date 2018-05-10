package main

import (
	"net/http"
	"html/template"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"fmt"
	"io"
	"io/ioutil"
)

var memCacheKey = "last_request"

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/webhook", webhook)
	http.HandleFunc("/webhook/list", webhookList)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}

func webhook(w http.ResponseWriter, r *http.Request) {
	data := getBodyString(r)
	ctx := appengine.NewContext(r)
	item := &memcache.Item{
		Key:   memCacheKey,
		Value: data,
	}
	if err := memcache.Set(ctx, item); err != nil {
		fmt.Fprint(w, err)
		return
	}

	io.WriteString(w, string(data))
}

func webhookList(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if item, err := memcache.Get(ctx, memCacheKey); err == memcache.ErrCacheMiss {
		io.WriteString(w, "No item")
	} else {
		io.WriteString(w, string(item.Value[:]))
	}
}

func getBodyString(r *http.Request) []byte {
	var bodyBuffer []byte
	if r.Body != nil {
		bodyBuffer, _ = ioutil.ReadAll(r.Body)
	}

	return bodyBuffer
}
