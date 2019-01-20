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

	h := NewHandler(s)

	log.Fatal(http.ListenAndServe(":8081", h))
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//t := template.Must(template.New("arc").Parse("arc.html"))
	t := template.Must(template.ParseFiles("arc.html"))

	path := r.URL.Path
	fmt.Println(path)
	if path == "" || path == "/" {
		fmt.Println("found")
		t = template.Must(template.ParseFiles("welcome.html"))
		t.Execute(w, struct{ Title string }{Title: "Welcome"})
		return
	}

	key := strings.SplitAfter(path, "/arc/")

	if len(key) < 2 {
		return
	}

	if v, ok := s[key[1]]; ok {
		err := t.Execute(w, v)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Story not found")
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
