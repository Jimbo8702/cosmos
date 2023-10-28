package render

import (
	"Jimbo8702/randomThoughts/cosmos/config"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type Renderer interface {
	Page(http.ResponseWriter, *http.Request, string, interface{}, interface{}) error
}

type TemplateData struct {
	CSRFToken 		string
	Port 			string
	ServerName 		string
	IsAuthenticated bool
	Secure 			bool
	IntMap 			map[string]int
	StringMap 		map[string]string
	FloatMap 		map[string]float32
	Data 			map[string]interface{}
}

type Render struct {
	RootPath 	string
	Secure 	 	bool
	Port 	 	string
	ServerName 	string
	Session 	*scs.SessionManager
}

func (c *Render) CheckAuth(td *TemplateData, r *http.Request) *TemplateData {
	if c.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	return td
}

func New(sess *scs.SessionManager, config *config.Config) *Render {
	return &Render{
		RootPath: config.RootPath,
		Port: config.Port,
		Session: sess,
	}
}


