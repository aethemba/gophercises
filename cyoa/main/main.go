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

func LoadJSONStory(data []byte) Story {

	var result Story
	err := json.Unmarshal(data, &result)

	if err != nil {
		fmt.Printf("%s", err)
	}

	return result
}

func main() {

	var storyFilename = flag.String("filename", "story.json", "path to json file with the story")

	flag.Parse()

	data := LoadFile(*storyFilename)
	s = LoadJSONStory(data)

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/arc/", arcHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func arcHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("arc.html"))

	path := r.URL.Path

	key := strings.SplitAfter(path, "/arc/")

	data := struct {
		Title   string
		Story   string
		Options []Option
	}{
		Title: "Arc page",
	}

	if v, ok := s[key[1]]; ok {
		data.Story = v.Story[0]
		data.Options = v.Options
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Story not found")
		return
	}

	t.Execute(w, data)
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
