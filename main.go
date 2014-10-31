package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"log"
	"math/rand"
	"time"
)

var data = make(map[string]interface{})

func ScrapProj() map[string]interface{} {
	doc, err := goquery.NewDocument("https://github.com/karan/Projects/blob/master/README.md")
	if err != nil {
		log.Fatal(err)
	}

	i := random(7, 95)

	doc.Find(".markdown-body p").Eq(i).Each(func(i int, s *goquery.Selection) {
		data["title"] = s.Find("strong").Text()
		data["desc"] = s.Text()
	})

	return data
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/suggest", func(ren render.Render) {
		ren.HTML(200, "index", ScrapProj())
	})

	m.Run()
}
