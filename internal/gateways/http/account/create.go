package account

import (
	"bank-service/extension/http/rest"
	"bank-service/internal/domain/entities"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type CreateAccountRequest struct {
	Name   string `json:"name" validate:"required"`
	Cpf    string `json:"cpf" validate:"required"`
	Secret string `json:"secret" validate:"required"`
}

type CreateAccountResponse struct {
	Id        uuid.UUID `json:"id"`
	Balance   float64   `json:"balance"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	CreatedAt time.Time `json:"createdAt"`
}

func (h Handler) Create(r *http.Request) rest.Response {
	ctx := r.Context()

	var body CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return rest.BadRequest(err)
	}

	output, err := h.AccountUc.Create(ctx, entities.CreateAccountInput{
		Name:   body.Name,
		CPF:    body.Cpf,
		Secret: body.Secret,
	})

	if err != nil {
		//TODO: REFACTOR ERROR, PUT ERRORS ON PORTS?
		//ERROR WHEN A ACC ALREADY EXISTS? 422!?
		return rest.InternalServerError(err)
	}

	return rest.Created(CreateAccountResponse{
		Id:        output.Id(),
		Balance:   output.Balance(),
		Name:      output.Name(),
		Cpf:       output.Cpf(),
		CreatedAt: output.CreatedAt(),
	})
}
