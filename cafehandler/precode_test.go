package cafehandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	res := responseRecorder.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
	data, err := io.ReadAll(res.Body)
	require.Nil(t, err)
	assert.NotEmpty(t, data)
}

func TestMainHandlerWhenCityNotSupport(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=orenburg&count=1", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	res := responseRecorder.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
	data, err := io.ReadAll(res.Body)
	require.Nil(t, err)
	assert.Equal(t, []byte("wrong city value"), data)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=5", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	res := responseRecorder.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
	data, err := io.ReadAll(res.Body)
	require.Nil(t, err)
	assert.Equal(t, totalCount, len(strings.Split(string(data), ",")))
}
