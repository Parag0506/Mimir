package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/your-org/mimir-ai/services/api-gateway-go/internal/handlers"
	"github.com/your-org/mimir-ai/services/api-gateway-go/internal/middleware"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.AuthMiddleware)
	router.GET("/health", handlers.HealthCheck)
	router.POST("/api/route", handlers.RouteRequest)
	return router
}

func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expected := `{"status":"API Gateway is up and running"}`
	if w.Body.String() != expected+"\n" {
		t.Errorf("Expected body %s, got %s", expected, w.Body.String())
	}
}

func TestRouteRequest_Unauthorized(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/api/route", bytes.NewBuffer([]byte(`{}`)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestRouteRequest_Success(t *testing.T) {
	router := setupRouter()

	// Mock Authorization header
	req, _ := http.NewRequest("POST", "/api/route", bytes.NewBuffer([]byte(`{"key":"value"}`)))
	req.Header.Set("Authorization", "Bearer valid-token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expected := `{"message":"Request has been routed successfully"}`
	if w.Body.String() != expected+"\n" {
		t.Errorf("Expected body %s, got %s", expected, w.Body.String())
	}
}
