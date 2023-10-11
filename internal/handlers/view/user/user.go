package user

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"path"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/handlers/api/user"
)

const (
	internalServerErrorCode = http.StatusInternalServerError
)

func Get(log *slog.Logger, proto string, host string, port int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get user"

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			const msg = "user id is not a number"
			http.Error(w, msg, internalServerErrorCode)
			log.Warn(fmt.Sprintf("%s: %s", errMsg, msg), "err", err)
			return
		}

		response, err := http.Get(fmt.Sprintf("%s://%s:%d/api/user/%d", proto, host, port, id))
		if err != nil {
			const msg = "get request to api failed"
			http.Error(w, msg, internalServerErrorCode)
			log.Warn(fmt.Sprintf("%s: %s", errMsg, msg), "err", err)
			return
		}

		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			const msg = "api response body couldn't be read"
			http.Error(w, msg, internalServerErrorCode)
			log.Error(fmt.Sprintf("%s: %s", errMsg, msg), "err", err)
			return
		}

		var gr user.GetResponse
		err = json.Unmarshal(bytes, &gr)
		if err != nil {
			const msg = "api response couldn't be unmarshalled"
			http.Error(w, msg, internalServerErrorCode)
			log.Error(fmt.Sprintf("%s: %s", errMsg, msg), "err", err)
			return
		}

		if gr.Status != http.StatusOK {
			http.Error(w, gr.Error, gr.Status)
			log.Warn(fmt.Sprintf("request to api returned not-ok status"), "response", gr)
			return
		}

		tp := path.Join("internal", "views", "user", "user.tmpl")

		tmpl, err := template.ParseFiles(tp)
		if err != nil {
			const msg = "can't parse html template"
			http.Error(w, msg, internalServerErrorCode)
			log.Error(fmt.Sprintf("%s: %s", errMsg, msg), "template path", tp)
			return
		}

		err = tmpl.Execute(w, gr)
		if err != nil {
			const msg = "can't execute html template"
			http.Error(w, msg, internalServerErrorCode)
			log.Error(fmt.Sprintf("%s: %s", errMsg, msg), "template path", tp, "template", tmpl)
		}

		log.Info("user view rendered")
	}
}
