package render

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path"
)

type Layout struct {
	Content template.HTML
	Meta    map[string]interface{}
}

func (l *Layout) Render(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	template.Must(template.ParseFiles(path.Join("templates", "layout.html"))).Execute(w, l)
}

func (l *Layout) GetMeta(key string) interface{} {
	return l.Meta[key]
}
