package shorten

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	// go-sqlite3 loading ...
	_ "github.com/mattn/go-sqlite3"
)

// API ...
type API struct {
	r Repository
}

// NewShorten ...
func NewShorten() (*API, error) {

	r, err := NewRepository("sqlite3", "shorten_url.db", 10, 10)
	if err != nil {
		return nil, err
	}

	api := API{r: r}
	return &api, nil

}

func (sapi API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var (
			err           error
			clientRequest ShortURLModel
		)

		if err := json.NewDecoder(r.Body).Decode(&clientRequest); err != nil {
			log.Printf("failed to decode JSON %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println(clientRequest)
		shortenID := uuid.New()
		clientRequest.ID = shortenID.String()
		err = sapi.r.Create(&clientRequest)
		if err != nil {
			log.Printf("error creating entry: %v", err)
		}
		w.Header().Set("Location", "/"+string(shortenID.String()))
		w.WriteHeader(http.StatusCreated)

	case "GET":
		model, err := sapi.r.FindByID(r.URL.Path[1:])
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		w.Header().Add("Location", model.Location)
		w.WriteHeader(http.StatusSeeOther)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}
