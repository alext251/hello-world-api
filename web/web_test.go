package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, header http.Header, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	req.Header = header
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// TestGetWithAcceptJSONHeader tests the GET request with the Accept
// Header set to "application/json" and expects the response
// {"message": "Hello, World"}.
func TestGetWithAcceptJSONHeader(t *testing.T) {
	router := setupRouter(false)
	header := http.Header{"Accept": []string{"application/json"}}
	expected := gin.H{"message": "Hello, World"}

	w := performRequest(router, header, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)

	var actual gin.H
	err := json.Unmarshal([]byte(w.Body.String()), &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

// TestGetWithNoAcceptJSONHeader tests the GET request with
// Accept header unset, and with an Accept header image/webp
// and expectes reponse <p>Hello, World</p>.
func TestGetWithNoAcceptJSONHeader(t *testing.T) {
	router := setupRouter(false)
	headers := []http.Header{
		http.Header{},
		http.Header{"Accept": []string{"image/webp"}},
	}
	expected := []byte("<p>Hello, World</p>")

	for _, header := range headers {
		w := performRequest(router, header, "GET", "/")
		assert.Equal(t, http.StatusOK, w.Code)

		actual := []byte(w.Body.String())
		assert.Equal(t, expected, actual)
	}
}

// TestPost tests the POST request and returns response
// {"message": "Hello, World"}
func TestPost(t *testing.T) {
	router := setupRouter(false)
	header := http.Header{}
	expected := []byte("<p>Hello, World</p>")

	w := performRequest(router, header, "POST", "/")

	assert.Equal(t, http.StatusOK, w.Code)

	actual := []byte(w.Body.String())
	assert.Equal(t, expected, actual)
}
