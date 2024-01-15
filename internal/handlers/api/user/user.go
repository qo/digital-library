package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/logger"
	"github.com/qo/digital-library/internal/storage/author"
	"github.com/qo/digital-library/internal/storage/book"
	"github.com/qo/digital-library/internal/storage/book_review"
	"github.com/qo/digital-library/internal/storage/user"
)

type userStorage interface {
	PostUser(*user.User) error
	GetUser(id int) (*user.User, error)
	PutUser(user *user.User) error
	DeleteUser(id int) error
	GetUserFavoriteBooks(id int) ([]book.Book, error)
	GetUserFavoriteAuthors(id int) ([]author.Author, error)
	GetUserBookReviews(id int) ([]book_review.BookReview, error)
}

type userHandler struct {
	logger.Logger
	userStorage
}

func New(log logger.Logger, us userStorage) *userHandler {
	return &userHandler{
		log,
		us,
	}
}

type postRequest = user.User

type postResponse struct {
	Error string `json:"error,omitempty"`
}

func (uh *userHandler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't post user"

		rd, we := json.NewDecoder(r.Body), json.NewEncoder(w)

		var req postRequest

		err := rd.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(postResponse{
				Error: "invalid request",
			})
			uh.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		err = uh.PostUser(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(postResponse{
				Error: "db error",
			})
			uh.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		uh.Debug("post user success")

		w.WriteHeader(http.StatusCreated)

		we.Encode(postResponse{})
	}
}

type getResponse struct {
	Error string `json:"error,omitempty"`
	user.User
}

func (uh *userHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get user"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(getResponse{
				Error: "user id is not a number",
			})
			uh.Error(fmt.Sprintf("%s: user id is not a number: %s", errMsg, err))
			return
		}

		user, err := uh.GetUser(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(getResponse{
				Error: "db error",
			})
			uh.Error(fmt.Sprintf("%s: user with %d id doesn't exist: %s", errMsg, id, err))
			return
		}

		uh.Debug("get user success", "user", user)

		w.WriteHeader(http.StatusOK)

		// can't mix field:value and value syntax
		we.Encode(getResponse{
			"",
			*user,
		})
	}
}

type putRequest = user.User

type putResponse struct {
	Error string `json:"error,omitempty"`
}

func (uh *userHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't put user"

		rd, we := json.NewDecoder(r.Body), json.NewEncoder(w)

		var req putRequest

		err := rd.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(putResponse{
				Error: "invalid request",
			})
			uh.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		uh.Debug("request parsed", "req", req)

		err = uh.PutUser(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(putResponse{
				Error: "db error",
			})
			uh.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		uh.Debug("put user success")

		w.WriteHeader(http.StatusOK)

		we.Encode(putResponse{})
	}
}

type deleteResponse struct {
	Error string `json:"error,omitempty"`
}

func (uh *userHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't delete user"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(deleteResponse{
				Error: "user id is not a number",
			})
			uh.Error(fmt.Sprintf("%s: user id is not a number: %s", errMsg, err))
			return
		}

		err = uh.DeleteUser(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(deleteResponse{
				Error: "db error",
			})
			uh.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		uh.Debug("delete user success")

		w.WriteHeader(http.StatusOK)

		we.Encode(deleteResponse{})
	}
}

type getFavoriteBooksResponse struct {
	Error string `json:"error,omitempty"`
	Books []book.Book
}

func (uh *userHandler) GetFavoriteBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get favorite books"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(deleteResponse{
				Error: "user id is not a number",
			})
			uh.Error(fmt.Sprintf("%s: user id is not a number: %s", errMsg, err))
			return
		}

		books, err := uh.GetUserFavoriteBooks(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(getFavoriteBooksResponse{
				Error: "db error",
			})
			uh.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		w.WriteHeader(http.StatusOK)
		we.Encode(getFavoriteBooksResponse{
			Books: books,
		})
		uh.Info("favorite books fetched")
	}
}
