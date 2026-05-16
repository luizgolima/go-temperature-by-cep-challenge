package viacep

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestViaCEPResponse_Unmarshal(t *testing.T) {
	t.Run("should succeed to unmarshal when erro is string", func(t *testing.T) {
		jsonData := `{"erro": "true"}`
		var data ViaCEPResponse
		err := json.Unmarshal([]byte(jsonData), &data)
		assert.NoError(t, err)
		
		isError := false
		if data.Erro != nil {
			if b, ok := data.Erro.(bool); ok && b {
				isError = true
			} else if s, ok := data.Erro.(string); ok && s == "true" {
				isError = true
			}
		}
		assert.True(t, isError)
	})

	t.Run("should succeed to unmarshal when erro is bool", func(t *testing.T) {
		jsonData := `{"erro": true}`
		var data ViaCEPResponse
		err := json.Unmarshal([]byte(jsonData), &data)
		assert.NoError(t, err)

		isError := false
		if data.Erro != nil {
			if b, ok := data.Erro.(bool); ok && b {
				isError = true
			} else if s, ok := data.Erro.(string); ok && s == "true" {
				isError = true
			}
		}
		assert.True(t, isError)
	})

	t.Run("should succeed to unmarshal when erro is missing", func(t *testing.T) {
		jsonData := `{"localidade": "Sao Paulo"}`
		var data ViaCEPResponse
		err := json.Unmarshal([]byte(jsonData), &data)
		assert.NoError(t, err)
		assert.Nil(t, data.Erro)
		assert.Equal(t, "Sao Paulo", data.Localidade)
	})
}
