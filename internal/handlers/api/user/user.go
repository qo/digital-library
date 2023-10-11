package user

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/qo/digital-library/internal/storage"
)

type UserStorage interface {
	GetUser(id int) (*storage.User, error)
	PutUser(user *storage.User) (int, error)
	DeleteUser(id int) (int, error)
}

type GetResponse struct {
	Status     int    `json:"status"` // Error, Ok
	Error      string `json:"error,omitempty"`
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

func Get(log *slog.Logger, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get user"

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			render.JSON(w, r, GetResponse{
				Status:     http.StatusInternalServerError,
				Error:      "user id is not a number",
				Id:         0,
				FirstName:  "",
				SecondName: "",
			})
			log.Error(fmt.Sprintf("%s: user id is not a number: %s", errMsg, err))
			return
		}

		user, err := us.GetUser(id)
		// TODO: check type of error
		if err != nil {
			render.JSON(w, r, GetResponse{
				Status:     http.StatusInternalServerError,
				Error:      "db error",
				Id:         0,
				FirstName:  "",
				SecondName: "",
			})
			log.Error(fmt.Sprintf("%s: user with %d id doesn't exist: %s", errMsg, id, err))
			return
		}

		log.Debug("get user success", "user", user)

		render.JSON(w, r, GetResponse{
			Status:     http.StatusOK,
			Id:         user.Id,
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
		})
	}
}

type PutRequest struct {
	Id         int    `json:"id"          validate:"required"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type PutResponse struct {
	Status int    `json:"status"` // Error, Ok
	Error  string `json:"error,omitempty"`
	Id     int    `json:"id"`
}

func Put(log *slog.Logger, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't put user"

		var req PutRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.JSON(w, r, PutResponse{
				Status: http.StatusInternalServerError,
				Error:  "invalid request",
				Id:     0,
			})
			log.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		log.Debug("request parsed", "req", req)

		user := storage.User{
			Id:         req.Id,
			FirstName:  req.FirstName,
			SecondName: req.SecondName,
		}

		id, err := us.PutUser(&user)
		// TODO: check type of error
		if err != nil {
			render.JSON(w, r, PutResponse{
				Status: http.StatusInternalServerError,
				Error:  "db error",
				Id:     id,
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("put user success", "id", id)

		render.JSON(w, r, PutResponse{
			Status: http.StatusOK,
			Id:     id,
		})
	}
}

type DeleteResponse struct {
	Status int    `json:"status"` // Error, Ok
	Error  string `json:"error,omitempty"`
	Id     int    `json:"id"`
}

func Delete(log *slog.Logger, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't delete user"

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			render.JSON(w, r, DeleteResponse{
				Status: http.StatusInternalServerError,
				Error:  "user id is not a number",
				Id:     0,
			})
			log.Error(fmt.Sprintf("%s: user id is not a number: %s", errMsg, err))
			return
		}

		deletedId, err := us.DeleteUser(id)
		// TODO: check type of error
		if err != nil {
			render.JSON(w, r, DeleteResponse{
				Status: http.StatusInternalServerError,
				Error:  "db error",
				Id:     0,
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("delete user success", "deletedId", deletedId)

		render.JSON(w, r, DeleteResponse{
			Status: http.StatusOK,
			Id:     id,
		})
	}
}
