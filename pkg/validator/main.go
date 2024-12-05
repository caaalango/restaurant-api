package validator

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
)

// Define uma chave para o contexto (usando um tipo anônimo para evitar colisões)
type contextKey string

const bodyContextKey contextKey = "body"

// Função para armazenar o corpo no contexto
func WithBody(ctx context.Context, body io.Reader) context.Context {
	return context.WithValue(ctx, bodyContextKey, body)
}

func ValidateBody[T any](ctx context.Context) error {
	body, ok := ctx.Value(bodyContextKey).(io.Reader)
	if !ok {
		return fmt.Errorf("body not found in context")
	}

	var req T
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&req)
	if err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}

	// Criar o validador
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return fmt.Errorf("validation failed: %v", err)
	}

	return nil
}
