// schemas/callOpeningRequest.go

package schemas

import (
	"github.com/go-playground/validator/v10"
)

// CallOpeningRequest representa a estrutura de uma solicitação para chamar uma abertura
type CallOpeningRequest struct {
	ID uint `json:"id" validate:"required"`
}

// Validate verifica se a solicitação de chamada de abertura é válida
func (r *CallOpeningRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
