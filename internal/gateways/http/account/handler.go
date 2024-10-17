package account

import (
	"bank-service/extension/http/rest"
	"bank-service/internal/domain/usecase/account"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	AccountUc account.UseCase
}

func (h Handler) AccountRoutes(r chi.Router) {
	r.Post("/accounts", rest.Handle(h.Create))
	r.Get("/accounts", rest.Handle(h.GetBalance))
	r.Get("/accounts/{account-id}/balance", rest.Handle(h.GetBalance))
}
