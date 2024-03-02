package handlers

import (
	"log"
	"net/http"

	"github.com/sikozonpc/fullstackgo/views"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return
	}

	views.Home(user).Render(r.Context(), w)
}
