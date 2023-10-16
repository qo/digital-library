package book

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/storage"
)

type BookStorage interface {
	PostBook(*storage.Book) error
	GetBook(id int) (*storage.Book, error)
	PutBook(book *storage.Book) error
	DeleteBook(id int) error
}

type PostRequest = storage.Book

type PostResponse struct {
	Error string `json:"error,omitempty"`
}

func Post(log *slog.Logger, bs BookStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't post book"

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

		err = bs.PostBook(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(PostResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("post book success")

		w.WriteHeader(http.StatusCreated)

		we.Encode(PostResponse{})
	}
}

type GetResponse struct {
	Error string `json:"error,omitempty"`
	storage.Book
}

func Get(log *slog.Logger, bs BookStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get book"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(GetResponse{
				Error: "book id is not a number",
			})
			log.Error(fmt.Sprintf("%s: book id is not a number: %s", errMsg, err))
			return
		}

		book, err := bs.GetBook(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(GetResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: book with %d id doesn't exist: %s", errMsg, id, err))
			return
		}

		log.Debug("get book success", "book", book)

		w.WriteHeader(http.StatusOK)

		we.Encode(GetResponse{
			"",
			*book,
		})
	}
}

type PutRequest = storage.Book

type PutResponse struct {
	Error string `json:"error,omitempty"`
}

func Put(log *slog.Logger, bs BookStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't put book"

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

		err = bs.PutBook(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(PutResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("put book success")

		w.WriteHeader(http.StatusOK)

		we.Encode(PutResponse{})
	}
}

type DeleteResponse struct {
	Error string `json:"error,omitempty"`
}

func Delete(log *slog.Logger, bs BookStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't delete book"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(DeleteResponse{
				Error: "book id is not a number",
			})
			log.Error(fmt.Sprintf("%s: book id is not a number: %s", errMsg, err))
			return
		}

		err = bs.DeleteBook(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(DeleteResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("delete book success")

		w.WriteHeader(http.StatusOK)

		we.Encode(DeleteResponse{})
	}
}
