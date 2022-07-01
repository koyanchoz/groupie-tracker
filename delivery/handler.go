package delivery

import (
	"encoding/json"
	"groupie-tracker/config"
	"groupie-tracker/structs"
	"net/http"
	"strings"
	"text/template"
	"time"
)

var AllArtists []structs.Artist

type Handler struct {
	// templates
	MainT   *template.Template
	ArtistT *template.Template
	ErrT    *template.Template
	client  *http.Client
}

func NewHandler(cfg config.Config) (*Handler, error) {
	mainTemplate, err := template.ParseFiles(cfg["index"])
	if err != nil {
		return nil, err
	}
	errTemplate, err := template.ParseFiles(cfg["error"])
	if err != nil {
		return nil, err
	}
	artistTemplate, err := template.ParseFiles(cfg["artist"])
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	return &Handler{
		mainTemplate,
		artistTemplate,
		errTemplate,
		client,
	}, nil
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.ErrorHandler(w, http.StatusNotFound, ErrNotFound)
		return
	}
	if r.Method != http.MethodGet {
		h.ErrorHandler(w, http.StatusBadRequest, ErrMethodNotAllowed)
		return
	}
	resp, err := h.client.Get("http://groupietrackers.herokuapp.com/api")
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}

	var linksObject structs.Links

	if err = json.NewDecoder(resp.Body).Decode(&linksObject); err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	resp, err = h.client.Get(linksObject.ArtistsURL)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&AllArtists); err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}
	err = h.MainT.Execute(w, AllArtists)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}
}

func (h *Handler) ArtistPage(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	var (
		artist    structs.Artist
		relations structs.Relatations
	)
	response, err := h.client.Get("http://groupietrackers.herokuapp.com/api")
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}
	var linksObject structs.Links
	if err := json.NewDecoder(response.Body).Decode(&linksObject); err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}
	artistInfo, err := h.client.Get(linksObject.ArtistsURL + "/" + id)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}

	if err = json.NewDecoder(artistInfo.Body).Decode(&artist); err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	RelationsInfo, err := h.client.Get(artist.Relations)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}
	if err = json.NewDecoder(RelationsInfo.Body).Decode(&relations); err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, err)
	}

	info := structs.About{
		Image:        artist.Image,
		Name:         artist.Name,
		Members:      artist.Members,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Relations:    relations.DatesLocations,
	}
	if err = h.ArtistT.Execute(w, info); err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, ErrServer)
		return
	}
}
