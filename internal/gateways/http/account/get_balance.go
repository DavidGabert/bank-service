package account

import (
	"bank-service/extension/http/rest"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
}

func (h Handler) GetBalance(r *http.Request) rest.Response {
	ctx := r.Context()
	accUuid := rest.URLParam(r, "account-id")
	parse, err := uuid.Parse(accUuid)

	if err != nil {
		return rest.BadRequest(fmt.Errorf("invalid UUID: %s", accUuid))
	}

	acc, err := h.AccountUc.GetAccountById(ctx, parse)
	if err != nil {
		//TODO: REFACTOR ERROR, PUT ERRORS ON PORTS?
		//404!?
		return rest.InternalServerError(err)
	}
	return rest.Ok(GetBalanceResponse{Balance: acc.Balance()})
}
