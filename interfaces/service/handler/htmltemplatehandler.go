package handler

import (
	"context"
	"html/template"
	"io"

	log "github.com/howood/kangaroochat/infrastructure/logger"
	"github.com/labstack/echo/v4"
)

const (
	templateFirstCheckDir  = "/etc/templates/"
	templateSecondCheckDir = "./templates/"
	templateThirdCheckDir  = "/go/templates/"
)

//HTMLTemplate is HTMLTemplate struct
type HTMLTemplate struct {
	Templates *template.Template
}

//Render is render html
func (t *HTMLTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

//LoadTemplate load html directories
func LoadTemplate(pattern string) (*template.Template, error) {
	var template *template.Template
	var err error
	ctx := context.Background()
	if template, err = template.ParseGlob(templateFirstCheckDir + pattern); err != nil {
		log.Debug(ctx, err)
		if template, err = template.ParseGlob(templateSecondCheckDir + pattern); err != nil {
			log.Debug(ctx, err)
			if template, err = template.ParseGlob(templateThirdCheckDir + pattern); err != nil {
				log.Debug(ctx, err)
				return template, err
			}
		}
	}
	return template, nil
}
