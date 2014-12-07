package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	s "strings"
	"time"
)

var data = make(map[string]interface{})

const filename = "db.txt"

type Proj struct {
	Title       string
	Description string
}

func ScrapProj() {
	fmt.Println("scraping ..")
	doc, err := goquery.NewDocument("https://github.com/karan/Projects/blob/master/README.md")
	if err != nil {
		log.Fatal(err)
	}

	if !isFileExist(filename) {
		// if the file doesn't exist, create a new one
		w, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer w.Close()

		// write empty one first
		err = ioutil.WriteFile(filename, []byte(""), 0644)

		doc.Find(".markdown-body p").Slice(7, 95).Each(func(i int, s *goquery.Selection) {
			txt := s.Text() + "\n"
			f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}

			defer f.Close()

			if _, err = f.WriteString(txt); err != nil {
				panic(err)
			}
		})
	}
}

func ReadProjFromFile() []string {
	fmt.Println("reading from file ..")
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return s.Split(string(b), "\n")
}

func GetRandomProj(projs []string) map[string]interface{} {
	i := random(0, len(projs))
	project := projs[i]

	data["title"] = s.Split(project, " - ")[0]
	data["desc"] = s.Split(project, " - ")[1]
	return data
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {

	// If the file doesn't exist, scrap it
	if !isFileExist(filename) {
		ScrapProj()
	}

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/suggest", func(ren render.Render) {
		ren.HTML(200, "index", GetRandomProj(ReadProjFromFile()))
	})

	m.Get("/", func(ren render.Render) {
		ren.HTML(200, "index", nil)
	})

	m.Run()
}

// Check if file already exists or not
func isFileExist(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}
