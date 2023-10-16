package user

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/storage"
)

type UserStorage interface {
	PostUser(*storage.User) error
	GetUser(id int) (*storage.User, error)
	PutUser(user *storage.User) error
	DeleteUser(id int) error
}

type PostRequest = storage.User

type PostResponse struct {
	Error string `json:"error,omitempty"`
}

func Post(log *slog.Logger, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't post user"

		rd, we := json.NewDecoder(r.Body), json.NewEncoder(w)

		var req PostRequest

		err := rd.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(PostResponse{
				Error: "invalid request",
			})
			log.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		err = us.PostUser(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(PostResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("post user success")

		w.WriteHeader(http.StatusCreated)

		we.Encode(PostResponse{})
	}
}

type GetResponse struct {
	Error      string `json:"error,omitempty"`
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

func Get(log *slog.Logger, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get user"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(GetResponse{
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
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(GetResponse{
				Error:      "db error",
				Id:         0,
				FirstName:  "",
				SecondName: "",
			})
			log.Error(fmt.Sprintf("%s: user with %d id doesn't exist: %s", errMsg, id, err))
			return
		}

		log.Debug("get user success", "user", user)

		w.WriteHeader(http.StatusOK)

		we.Encode(GetResponse{
			Id:         user.Id,
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
		})
	}
}

type PutRequest = storage.User

type PutResponse struct {
	Error string `json:"error,omitempty"`
}

func Put(log *slog.Logger, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't put user"

		rd, we := json.NewDecoder(r.Body), json.NewEncoder(w)

		var req PutRequest

		err := rd.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(PutResponse{
				Error: "invalid request",
			})
			log.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		log.Debug("request parsed", "req", req)

		err = us.PutUser(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(PutResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("put user success")

		w.WriteHeader(http.StatusOK)

		we.Encode(PutResponse{})
	}
}

type DeleteResponse struct {
	Error string `json:"error,omitempty"`
}

func Delete(log *slog.Logger, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't delete user"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(DeleteResponse{
				Error: "user id is not a number",
			})
			log.Error(fmt.Sprintf("%s: user id is not a number: %s", errMsg, err))
			return
		}

		err = us.DeleteUser(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(DeleteResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("delete user success")

		w.WriteHeader(http.StatusOK)

		we.Encode(DeleteResponse{})
	}
}
