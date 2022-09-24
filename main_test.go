package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDefaultRoute(test *testing.T) {
	test.Parallel()

	responseWritter := httptest.NewRecorder()
	testContext, router := gin.CreateTestContext(responseWritter)
	setupRouter(router)

	req, err := http.NewRequestWithContext(testContext, "GET", "/", nil)
	if err != nil {
		test.Errorf("Got an error: %s", err)
	}

	router.ServeHTTP(responseWritter, req)
	if http.StatusOK != responseWritter.Code {
		test.Fatalf("Expected response code %b, got %b", http.StatusOK, responseWritter.Code)
	}

	body := responseWritter.Body.String()

	expectedResponse := "Hello, gin!"

	if expectedResponse != strings.Trim(body, " \r\n") {
		test.Fatalf("Expected response body '%s', got '%s'", expectedResponse, body)
	}
}
