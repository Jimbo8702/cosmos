package render

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

type JetRenderer struct {
	*Render
	views 		*jet.Set
}

func NewJetRenderer(r *Render, rootPath string) Renderer {
	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)), 
		jet.InDevelopmentMode(),
	)
	
	return &JetRenderer{
		Render: r,
		views: views,
	}
}

func (c *JetRenderer) Page(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	
	td = c.CheckAuth(td, r)

	t, err := c.views.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}
	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}