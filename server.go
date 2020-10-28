package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tabvn/goweb/db"
	"github.com/tabvn/goweb/render"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	r := httprouter.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		layout := &render.Layout{Content: render.Template(&render.Page{
			Template: "page-front.html",
			Title:    "Home page",
			Data: map[string]interface{}{
				"nodes": db.Nodes,
			},
		}), Meta: map[string]interface{}{"title": "This is page title"}}
		layout.Render(w, r, p)
	})

	r.GET("/node/:slug", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		slug := p.ByName("slug")
		node := db.FindBySlug(slug)
		layout := &render.Layout{Content: render.Template(&render.Page{
			Title: node.Title,
			Template: "page-node.html",
			Data: map[string]interface{}{
				"node": node,
			},
		}), Meta: map[string]interface{}{"title": node.Title}}
		layout.Render(w, r, p)
	})
	log.Println(fmt.Sprintf("server is running at: http://127.0.0.1%v", port))
	log.Fatal(http.ListenAndServe(port, r))
}
