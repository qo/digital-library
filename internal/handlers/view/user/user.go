package user

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/handlers/api/user"
	"github.com/qo/digital-library/internal/logger"
)

const (
	internalServerErrorCode = http.StatusInternalServerError
)

type userHandler struct {
	logger.Logger
	user.UserStorage
}

func New(log logger.Logger, st user.UserStorage) *userHandler {
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

		handler := user.New(uh.Logger, uh.UserStorage)
		user, err := handler.GetUser(id)
		if err != nil {
			const msg = "api request couldn't be done"
			http.Error(w, msg, internalServerErrorCode)
			uh.Error(fmt.Sprintf("%s: %s", errMsg, msg), "err", err)
			return
		}

		tp := path.Join("internal", "views", "user", "user.tmpl")

		tmpl, err := template.ParseFiles(tp)
		if err != nil {
			const msg = "can't parse html template"
			http.Error(w, msg, internalServerErrorCode)
			uh.Error(fmt.Sprintf("%s: %s", errMsg, msg), "template path", tp)
			return
		}

		err = tmpl.Execute(w, user)
		if err != nil {
			const msg = "can't execute html template"
			http.Error(w, msg, internalServerErrorCode)
			uh.Error(fmt.Sprintf("%s: %s", errMsg, msg), "template path", tp, "template", tmpl)
		}

		uh.Info("user view rendered")
	}
}
