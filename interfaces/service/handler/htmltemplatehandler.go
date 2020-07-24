package handler

import (
	"context"
	"html/template"
	"io"

	log "github.com/howood/kangaroochat/infrastructure/logger"
	"github.com/labstack/echo/v4"
)

const (
	TemplateFirstCheckDir  = "/etc/templates/"
	TemplateSecondCheckDir = "./templates/"
	TemplateThirdCheckDir  = "/go/templates/"
)

type HtmlTemplate struct {
	Templates *template.Template
}

func (t *HtmlTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func LoadTemplate(pattern string) (*template.Template, error) {
	var template *template.Template
	var err error
	ctx := context.Background()
	if template, err = template.ParseGlob(TemplateFirstCheckDir + pattern); err != nil {
		log.Debug(ctx, err)
		if template, err = template.ParseGlob(TemplateSecondCheckDir + pattern); err != nil {
			log.Debug(ctx, err)
			if template, err = template.ParseGlob(TemplateThirdCheckDir + pattern); err != nil {
				log.Debug(ctx, err)
				return template, err
			}
		}
	}
	return template, nil
}
