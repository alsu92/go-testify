package homework

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	city := "moscow"
	totalCount := len(cafeList[city])
	queryCount := totalCount + 1
	query := fmt.Sprintf("/cafe?count=%d&city=%s", queryCount, city)
	req := httptest.NewRequest(http.MethodGet, query, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	respBody := strings.Split(responseRecorder.Body.String(), ",")
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Len(t, respBody, totalCount)
}

func TestMainHandlerWhenOk(t *testing.T) {
	city := "moscow"
	totalCount := len(cafeList[city])
	query := fmt.Sprintf("/cafe?count=%d&city=%s", totalCount, city)
	req := httptest.NewRequest(http.MethodGet, query, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	respBody := strings.Split(responseRecorder.Body.String(), ",")
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	//require.NotEmpty(t, responseRecorder.Body)
	assert.Len(t, respBody, totalCount)

}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	expectedRespBody := "wrong city value"
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=1&city=some_string", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, expectedRespBody, responseRecorder.Body.String())

}
