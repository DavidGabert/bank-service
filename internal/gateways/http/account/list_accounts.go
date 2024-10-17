package account

import (
	"bank-service/extension/http/rest"
	"bank-service/internal/domain/entities"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ListAccountResponse struct {
	Accounts []AccResponse `json:"accounts"`
}

type AccResponse struct {
	Id        uuid.UUID `json:"id"`
	Balance   float64   `json:"balance"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	CreatedAt time.Time `json:"createdAt"`
}

func (h Handler) List(r *http.Request) rest.Response {
	ctx := r.Context()
	output, err := h.AccountUc.ListAccounts(ctx)
	if err != nil {
		//TODO: REFACTOR ERROR, PUT ERRORS ON PORTS?
		return rest.InternalServerError(err)
	}

	return rest.Ok(parseAccountListResponse(output))
}
func parseAccountListResponse(accounts []entities.Account) ListAccountResponse {
	var accResponse []AccResponse

	for _, acc := range accounts {
		accResponse = append(accResponse, AccResponse{
			Id:        acc.Id(),
			Balance:   acc.Balance(),
			Name:      acc.Name(),
			Cpf:       acc.Cpf(),
			CreatedAt: acc.CreatedAt(),
		})
	}

	response := ListAccountResponse{
		Accounts: accResponse,
	}
	return response
}
