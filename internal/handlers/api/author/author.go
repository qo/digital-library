package author

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/logger"
	"github.com/qo/digital-library/internal/storage/author"
)

type authorStorage interface {
	GetAuthor(id int) (*author.Author, error)
	PostAuthor(*author.Author) error
	PutAuthor(author *author.Author) error
	DeleteAuthor(id int) error
}

type authorHandler struct {
	logger.Logger
	authorStorage
}

func New(log logger.Logger, as authorStorage) *authorHandler {
	return &authorHandler{
		log,
		as,
	}
}

type postRequest = author.Author

type postResponse struct {
	Error string `json:"error,omitempty"`
}

func (ah authorHandler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't post author"

		rd, we := json.NewDecoder(r.Body), json.NewEncoder(w)

		var req postRequest

		err := rd.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(postResponse{
				Error: "invalid request",
			})
			ah.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		err = ah.PostAuthor(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(postResponse{
				Error: "db error",
			})
			ah.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		ah.Debug("post author success")

		w.WriteHeader(http.StatusCreated)

		we.Encode(postResponse{})
	}
}

type getResponse struct {
	Error string `json:"error,omitempty"`
	author.Author
}

func (ah authorHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get author"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(getResponse{
				Error: "author id is not a number",
			})
			ah.Error(fmt.Sprintf("%s: author id is not a number: %s", errMsg, err))
			return
		}

		author, err := ah.GetAuthor(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(getResponse{
				Error: "db error",
			})
			ah.Error(fmt.Sprintf("%s: author with %d id doesn't exist: %s", errMsg, id, err))
			return
		}

		ah.Debug("get author success", "author", author)

		w.WriteHeader(http.StatusOK)

		we.Encode(getResponse{
			"",
			*author,
		})
	}
}

type putRequest = author.Author

type putResponse struct {
	Error string `json:"error,omitempty"`
}

func (ah authorHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't put author"

		rd, we := json.NewDecoder(r.Body), json.NewEncoder(w)

		var req putRequest

		err := rd.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(putResponse{
				Error: "invalid request",
			})
			ah.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		ah.Debug("request parsed", "req", req)

		err = ah.PutAuthor(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(putResponse{
				Error: "db error",
			})
			ah.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		ah.Debug("put author success")

		w.WriteHeader(http.StatusOK)

		we.Encode(putResponse{})
	}
}

type deleteResponse struct {
	Error string `json:"error,omitempty"`
}

func (ah authorHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't delete author"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(deleteResponse{
				Error: "author id is not a number",
			})
			ah.Error(fmt.Sprintf("%s: author id is not a number: %s", errMsg, err))
			return
		}

		err = ah.DeleteAuthor(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(deleteResponse{
				Error: "db error",
			})
			ah.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		ah.Debug("delete author success")

		w.WriteHeader(http.StatusOK)

		we.Encode(deleteResponse{})
	}
}
