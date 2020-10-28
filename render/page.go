package render

import (
	"bytes"
	"html/template"
	"path"
)

type Page struct {
	Template string
	Title    string
	Data     map[string]interface{}
}

func (p *Page) GetData(key string) interface{}{
	return p.Data[key]
}

func Template(page *Page) template.HTML {
	var data bytes.Buffer
	template.Must(template.ParseFiles(path.Join("templates", page.Template))).Execute(&data, page)
	return template.HTML(data.String())
}
