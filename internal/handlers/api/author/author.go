package author

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qo/digital-library/internal/storage"
)

type AuthorStorage interface {
	PostAuthor(*storage.Author) error
	GetAuthor(id int) (*storage.Author, error)
	PutAuthor(author *storage.Author) error
	DeleteAuthor(id int) error
}

type PostRequest = storage.Author

type PostResponse struct {
	Error string `json:"error,omitempty"`
}

func Post(log *slog.Logger, as AuthorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't post author"

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

		err = as.PostAuthor(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(PostResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("post author success")

		w.WriteHeader(http.StatusCreated)

		we.Encode(PostResponse{})
	}
}

type GetResponse struct {
	Error string `json:"error,omitempty"`
	storage.Author
}

func Get(log *slog.Logger, as AuthorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't get author"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(GetResponse{
				Error: "author id is not a number",
			})
			log.Error(fmt.Sprintf("%s: author id is not a number: %s", errMsg, err))
			return
		}

		author, err := as.GetAuthor(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(GetResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: author with %d id doesn't exist: %s", errMsg, id, err))
			return
		}

		log.Debug("get author success", "author", author)

		w.WriteHeader(http.StatusOK)

		we.Encode(GetResponse{
			"",
			*author,
		})
	}
}

type PutRequest = storage.Author

type PutResponse struct {
	Error string `json:"error,omitempty"`
}

func Put(log *slog.Logger, as AuthorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't put author"

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

		err = as.PutAuthor(&req)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(PutResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("put author success")

		w.WriteHeader(http.StatusOK)

		we.Encode(PutResponse{})
	}
}

type DeleteResponse struct {
	Error string `json:"error,omitempty"`
}

func Delete(log *slog.Logger, as AuthorStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const errMsg = "can't delete author"

		we := json.NewEncoder(w)

		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			we.Encode(DeleteResponse{
				Error: "author id is not a number",
			})
			log.Error(fmt.Sprintf("%s: author id is not a number: %s", errMsg, err))
			return
		}

		err = as.DeleteAuthor(id)
		// TODO: check type of error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			we.Encode(DeleteResponse{
				Error: "db error",
			})
			log.Error(fmt.Sprintf("%s: %s", errMsg, err))
			return
		}

		log.Debug("delete author success")

		w.WriteHeader(http.StatusOK)

		we.Encode(DeleteResponse{})
	}
}
