package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"log"
)

type proj struct {
	title       string
	description string
}

func ScrapProjTest() proj {
	doc, err := goquery.NewDocument("https://github.com/karan/Projects/blob/master/README.md")
	if err != nil {
		log.Fatal(err)
	}

	projs := []proj{}
	i := random(7, 95)

	doc.Find(".markdown-body p").Each(func(i int, s *goquery.Selection) {
		p := proj{}
		//		data["title"] = s.Find("strong").Text()
		//		data["desc"] = s.Text()
		p.title = s.Find("strong").Text()
		p.description = s.Text()
		projs = append(projs, p)
	})

	return projs[i]
}

func mainStruct() {

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/suggest", func(ren render.Render) {
		ren.HTML(200, "index", ScrapProjTest())
	})

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Run()
}
