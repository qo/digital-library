package user

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/go-chi/chi/v5"
	user_handler "github.com/qo/digital-library/internal/handlers/api/user"
	"github.com/qo/digital-library/internal/logger"
	user_storage "github.com/qo/digital-library/internal/storage/user"
)

const (
	internalServerErrorCode = http.StatusInternalServerError
)

type userStorage interface {
	GetUser(id int) (*user_storage.User, error)
}

type userHandler struct {
	logger.Logger
	userStorage
}

func NewUserHandler(log logger.Logger, st userStorage) *userHandler {
	return &userHandler{
		log,
		st,
	}
}

func (uh *userHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get user"

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			const msg = "user id is not a number"
			http.Error(w, msg, internalServerErrorCode)
			uh.Warn(fmt.Sprintf("%s: %s", errMsg, msg), "err", err)
			return
		}

		response, err := user_handler.New(uh.Logger, uh.userStorage).GetUser(id)
		if err != nil {
			const msg = "get request to api failed"
			http.Error(w, msg, internalServerErrorCode)
			uh.Warn(fmt.Sprintf("%s: %s", errMsg, msg), "err", err)
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

		if response.StatusCode != http.StatusOK {
			http.Error(w, gr.Error, response.StatusCode)
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
