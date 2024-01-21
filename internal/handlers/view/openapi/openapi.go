package openapi

import (
	"fmt"
	"github.com/qo/digital-library/internal/logger"
	"html/template"
	"net/http"
	"os"
	"path"
)

const (
	internalServerErrorCode = http.StatusInternalServerError
)

type openapiHandler struct {
	logger.Logger
}

func New(log logger.Logger) *openapiHandler {
	return &openapiHandler{
		log,
	}
}

func (oh openapiHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get swagger"

		tp := path.Join("internal", "views", "swagger", "swagger.tmpl")

		tmpl, err := template.ParseFiles(tp)
		if err != nil {
			msg := fmt.Sprintf("%s: can't parse html template: %s", errMsg, err)
			http.Error(w, msg, internalServerErrorCode)
			oh.Error(msg, "template path", tp)
			return
		}

		jp := path.Join("docs", "swagger", "openapi.json")

		json, err := os.ReadFile(jp)
		if err != nil {
			msg := fmt.Sprintf("%s: can't read json file containing swagger spec: %s", errMsg, err)
			http.Error(w, msg, internalServerErrorCode)
			oh.Error(msg, "json path", jp)
			return
		}

		err = tmpl.Execute(w, string(json))
		if err != nil {
			msg := fmt.Sprintf("%s: can't execute html template: %s", errMsg, err)
			http.Error(w, msg, internalServerErrorCode)
			oh.Error(msg, "template path", tp, "template", tmpl, "json path", jp, "json", json)
		}

		oh.Info("swagger view rendered")
	}
}
