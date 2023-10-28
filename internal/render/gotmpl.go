package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type GoRenderer struct {
	*Render	
}

func NewGoRenderer(r *Render) Renderer {
	return &GoRenderer{
		Render: r,
	}
}

func (g *GoRenderer)  Page(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", g.RootPath, templateName))
	if err != nil {
		log.Println(err)
		return err
	}
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	td = g.CheckAuth(td, r)
	err = tmpl.Execute(w, &td)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}