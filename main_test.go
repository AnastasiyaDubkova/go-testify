package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, fmt.Sprintf("expected status code: %d, got %d", http.StatusOK, responseRecorder.Code))
	assert.NotEmpty(t, responseRecorder.Body.String(), "expected body: not empty, got empty")
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=penza", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, fmt.Sprintf("expected status code: %d, got %d", http.StatusBadRequest, responseRecorder.Code))

	expected := "wrong city value"
	assert.Equal(t, expected, responseRecorder.Body.String(), fmt.Sprintf("expected body: %s, got %s", expected, responseRecorder.Body.String()))

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, fmt.Sprintf("expected status code: %d, got %d", http.StatusOK, responseRecorder.Code))

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount, fmt.Sprintf("expected cafe count: %d, got %d", totalCount, len(list)))

}
