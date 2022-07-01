package delivery

import (
	"groupie-tracker/config"
	"log"
	"net/http"
)

func RegisterHTTPEndPoints(mux *http.ServeMux, c config.Config) {
	h, err := NewHandler(c)
	if err != nil {
		log.Println(err)
		return
	}
	mux.HandleFunc("/", h.MainPage)
	mux.HandleFunc("/artist/", h.ArtistPage)
	mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	mux.Handle("/template/assets/", http.StripPrefix("/template/assets/", http.FileServer(http.Dir("./template/assets/"))))

}
