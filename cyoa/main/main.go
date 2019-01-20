package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var tpl = template.Must(template.ParseFiles("arc.html"))

type Story map[string]Arc

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

var s Story

func LoadFile(filename string) []byte {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	return data
}

func ParseJSONStory(data []byte) Story {

	var result Story
	err := json.Unmarshal(data, &result)

	if err != nil {
		fmt.Printf("%s", err)
	}

	return result
}

func main() {

	var storyFilename = flag.String("file", "story.json", "the JSON file with the cyoa story")

	flag.Parse()

	data := LoadFile(*storyFilename)
	s = ParseJSONStory(data)

	// http.HandleFunc("/", welcomeHandler)
	// http.HandleFunc("/arc/", arcHandler)

	pathFn := func(r *http.Request) string {
		path := r.URL.Path
		if path == "/story" || path == "/story/" {
			return "/story/intro"
		}

		return path[len("/story/"):]
	}

	h := NewHandler(s, WithPathFn(pathFn))
	//h := NewHandler(s, WithTemplate(nil))

	log.Fatal(http.ListenAndServe(":8081 ", h))
}

func NewHandler(s Story, opt ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn}

	for _, opt := range opt {
		opt(&h)
	}

	return h
}

type handler struct {
	s        Story
	t        *template.Template
	pathFunc func(r *http.Request) string
}

type HandlerOption func(h *handler)

func defaultPathFn(r *http.Request) string {
	path := r.URL.Path
	fmt.Println(path)
	if path == "" || path == "/" {
		return "intro"
	}

	return path[1:]
}

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFunc = fn
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := h.pathFunc(r)
	fmt.Println(path)
	if v, ok := s[path]; ok {
		err := h.t.Execute(w, v)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Page not found", http.StatusNotFound)
	}

}

func arcHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("arc.html"))

	path := r.URL.Path

	key := strings.SplitAfter(path, "/arc/")

	if v, ok := s[key[1]]; ok {
		err := t.Execute(w, v)
		if err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Story not found")
	}

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("welcome.html"))

	data := struct {
		Title string
	}{
		Title: "Welcome page",
	}

	t.Execute(w, data)
}
