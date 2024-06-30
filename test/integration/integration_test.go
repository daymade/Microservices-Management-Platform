package integration

import (
	"encoding/json"
	"fmt"
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/route"
	"catalog-service-management-api/internal/adapter/api/viewmodel"
	"catalog-service-management-api/internal/app/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	// 设置测试环境
	setupTestEnv()

	// 初始化路由
	router = setupTestRouter()

	// 运行测试
	code := m.Run()

	// 清理环境
	teardownTestEnv()

	os.Exit(code)
}

func setupTestEnv() {
	os.Setenv("USE_DB", "true")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USER", "test")
	os.Setenv("DB_PASSWORD", "test")
	os.Setenv("DB_NAME", "testdb")
}

func teardownTestEnv() {
	os.Unsetenv("USE_DB")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
}

func setupTestRouter() *gin.Engine {
	r := gin.Default()
	sm := service.NewManager()
	sh := handler.NewServiceHandler(sm)
	uh := handler.NewUserHandler()
	route.SetupRoutes(r, sh, uh)
	return r
}

func TestListServices(t *testing.T) {
	validToken := "valid-token"

	testCases := []struct {
		name           string
		query          string
		expectedStatus int
		expectedCount  int
	}{
		{"All Services", "", 200, 15},
		{"Filtered Services", "API", 200, 4},
		{"No Results", "NonexistentService", 200, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/api/v1/services?query=%s", tc.query)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Authorization", validToken)
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)

			var response viewmodel.PaginatedResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCount, response.TotalCount)

			services, ok := response.Data.([]interface{})
			assert.True(t, ok, "Data should be a slice of interfaces")
			assert.Equal(t, tc.expectedCount, len(services))
		})
	}

	t.Run("Pagination", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/services?page=2&page_size=5", nil)
		req.Header.Set("Authorization", validToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var response viewmodel.PaginatedResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		services, ok := response.Data.([]interface{})
		assert.True(t, ok, "Data should be a slice of interfaces")
		assert.Equal(t, 5, len(services))
		assert.Equal(t, 15, response.TotalCount)
		assert.Equal(t, 2, response.Page)
		assert.Equal(t, 5, response.PageSize)
		assert.Equal(t, 3, response.TotalPages)
	})

	t.Run("Sorting", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/services?sort_by=name&sort_direction=asc", nil)
		req.Header.Set("Authorization", validToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var response viewmodel.PaginatedResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		services, ok := response.Data.([]interface{})
		assert.True(t, ok, "Data should be a slice of interfaces")
		assert.Greater(t, len(services), 0)

		var lastServiceName string
		for _, service := range services {
			serviceMap, ok := service.(map[string]interface{})
			assert.True(t, ok, "Each service should be a map")
			name, ok := serviceMap["name"].(string)
			assert.True(t, ok, "Service name should be a string")
			if lastServiceName != "" {
				assert.LessOrEqual(t, lastServiceName, name)
			}
			lastServiceName = name
		}
	})
}

func TestGetService(t *testing.T) {
	validToken := "valid-token"

	t.Run("Existing Service", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/services/1", nil)
		req.Header.Set("Authorization", validToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var service viewmodel.ServiceDetailViewModel
		err := json.Unmarshal(w.Body.Bytes(), &service)
		assert.NoError(t, err)
		assert.NotEmpty(t, service.Name)
		assert.NotEmpty(t, service.Description)
		assert.NotEmpty(t, service.OwnerID)
		assert.NotEmpty(t, service.Versions)
	})

	t.Run("Non-existent Service", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/services/999", nil)
		req.Header.Set("Authorization", validToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("Invalid Service ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/services/invalid", nil)
		req.Header.Set("Authorization", validToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

func TestUnauthorized(t *testing.T) {
	endpoints := []string{
		"/api/v1/services",
		"/api/v1/services/1",
	}

	for _, endpoint := range endpoints {
		t.Run(fmt.Sprintf("Unauthorized access to %s", endpoint), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", endpoint, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, 401, w.Code)
			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, "No authorization header provided", response["error"])
		})
	}
}
