package delivery

import (
	"net/http"
)

func (h *Handler) PathPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.ErrorHandler(w, http.StatusNotFound, ErrNotFound)
		return
	}
	if r.Method != http.MethodGet {
		h.ErrorHandler(w, http.StatusBadRequest, ErrBadRequest)
		return
	}
	h.MainT.Execute(w, AllArtists)
}
