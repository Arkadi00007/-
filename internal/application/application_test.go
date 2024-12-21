package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// проверяет успешный запрос
func TestCalcHandler_Success(t *testing.T) {
	reqBody := Request{
		Expression: "2+2",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
	w := httptest.NewRecorder()

	CalcHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	expectedResult := 4.0 // для примера
	if response.Result != expectedResult {
		t.Errorf("expected result %f; got %f", expectedResult, response.Result)
	}
}

// проверяет случай, когда передан некорректный JSON
func TestCalcHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader([]byte(`invalid`)))
	w := httptest.NewRecorder()

	CalcHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d; got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

// проверяет случай, когда передано пустое выражение
func TestCalcHandler_EmptyExpression(t *testing.T) {
	reqBody := Request{}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
	w := httptest.NewRecorder()

	CalcHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("expected status %d; got %d", http.StatusUnprocessableEntity, resp.StatusCode)
	}
}

// проверяет случай некорректного выражения
func TestCalcHandler_IncorrectSymbols(t *testing.T) {
	reqBody := Request{
		Expression: "invalid_expr",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
	w := httptest.NewRecorder()

	CalcHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != 422 { // Unprocessable Entity
		t.Errorf("expected status %d; got %d", 422, resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Error != "Expression is not valid" {
		t.Errorf("expected error 'Expression is not valid'; got %s", response.Error)
	}
}

// TestCalcHandler_MethodNotAllowed проверяет случай, если метод не POST
func TestCalcHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	w := httptest.NewRecorder()

	CalcHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d; got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}
