package handlers

import (
	"github.com/sikozonpc/fullstackgo/services/auth"
	"github.com/sikozonpc/fullstackgo/store"
)

type Handler struct {
	store *store.Storage
	auth  *auth.AuthService
}

func New(store *store.Storage, auth *auth.AuthService) *Handler {
	return &Handler{
		store: store,
		auth:  auth,
	}
}
