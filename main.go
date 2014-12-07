package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"log"
	"math/rand"
	"os"

	"time"
)

var data = make(map[string]interface{})

func ScrapProj(i int) {
	doc, err := goquery.NewDocument("https://github.com/karan/Projects/blob/master/README.md")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".markdown-body p").Each(func(i int, s *goquery.Selection) {
		txt := s.Text() + "\n"
		f, err := os.OpenFile("db.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		if _, err = f.WriteString(txt); err != nil {
			panic(err)
		}
	})
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/suggest", func(ren render.Render) {
		// i := random(7, 95)
		// ren.HTML(200, "index", ScrapProj(i))
	})

	m.Get("/", func(ren render.Render) {
		ren.HTML(200, "index", nil)
	})

	// m.Run()
	i := random(7, 95)
	ScrapProj(i)
}
