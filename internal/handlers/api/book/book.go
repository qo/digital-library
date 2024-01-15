package book

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/logger"
	"github.com/qo/digital-library/internal/storage/book"
)

type bookStorage interface {
	PostBook(*book.Book) error
	GetBook(id int) (*book.Book, error)
	PutBook(book *book.Book) error
	DeleteBook(id int) error
}

type bookHandler struct {
	logger.Logger
	bookStorage
}

func New(log logger.Logger, bs bookStorage) *bookHandler {
	return &bookHandler{
		log,
		bs,
	}
}

type postRequest = book.Book

type postResponse struct {
	Error string `json:"error,omitempty"`
}

func (bh *bookHandler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't post book"

		rd, we := json.NewDecoder(r.Body), json.NewEncoder(w)

		var req postRequest

		err := rd.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(postResponse{
				Error: "invalid request",
			})
			bh.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		err = bh.PostBook(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(postResponse{
				Error: "db error",
			})
			bh.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		bh.Debug("post book success")

		w.WriteHeader(http.StatusCreated)

		we.Encode(postResponse{})
	}
}

type getResponse struct {
	Error string `json:"error,omitempty"`
	book.Book
}

func (bh *bookHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get book"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(getResponse{
				Error: "book id is not a number",
			})
			bh.Error(fmt.Sprintf("%s: book id is not a number: %s", errMsg, err))
			return
		}

		book, err := bh.GetBook(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(getResponse{
				Error: "db error",
			})
			bh.Error(fmt.Sprintf("%s: book with %d id doesn't exist: %s", errMsg, id, err))
			return
		}

		bh.Debug("get book success", "book", book)

		w.WriteHeader(http.StatusOK)

		we.Encode(getResponse{
			"",
			*book,
		})
	}
}

type putRequest = book.Book

type putResponse struct {
	Error string `json:"error,omitempty"`
}

func (bh *bookHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't put book"

		rd, we := json.NewDecoder(r.Body), json.NewEncoder(w)

		var req putRequest

		err := rd.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(putResponse{
				Error: "invalid request",
			})
			bh.Error(fmt.Sprintf("%s: request not parsed: %s", errMsg, err), "req", req)
			return
		}

		bh.Debug("request parsed", "req", req)

		err = bh.PutBook(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(putResponse{
				Error: "db error",
			})
			bh.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		bh.Debug("put book success")

		w.WriteHeader(http.StatusOK)

		we.Encode(putResponse{})
	}
}

type deleteResponse struct {
	Error string `json:"error,omitempty"`
}

func (bh *bookHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't delete book"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(deleteResponse{
				Error: "book id is not a number",
			})
			bh.Error(fmt.Sprintf("%s: book id is not a number: %s", errMsg, err))
			return
		}

		err = bh.DeleteBook(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(deleteResponse{
				Error: "db error",
			})
			bh.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		bh.Debug("delete book success")

		w.WriteHeader(http.StatusOK)

		we.Encode(deleteResponse{})
	}
}
