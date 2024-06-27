package route

import (
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/app/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	type args struct {
		r  *gin.Engine
		sh *handler.ServiceHandler
		uh *handler.UserHandler
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"TestSetupRoutes",
			args{
				r:  gin.Default(),
				sh: handler.NewServiceHandler(service.NewManager()),
				uh: handler.NewUserHandler(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupRoutes(tt.args.r, tt.args.sh, tt.args.uh)

			// Create a test server
			ts := httptest.NewServer(tt.args.r)
			defer ts.Close()

			// Define test cases for each route
			testCases := []struct {
				method     string
				path       string
				statusCode int
			}{
				{"GET", "/api/v1/services", http.StatusOK},
				{"GET", "/api/v1/services/1", http.StatusOK},
				// Assuming POST is not allowed in read-only API
				// gin 在不匹配方法时就是 404 而不是 405
				{"POST", "/api/v1/services", http.StatusNotFound},
			}

			for _, tc := range testCases {
				t.Run(tc.method+" "+tc.path, func(t *testing.T) {
					req, _ := http.NewRequest(tc.method, ts.URL+tc.path, nil)

					req.Header.Add("Authorization", "Bearer token")

					resp, err := http.DefaultClient.Do(req)
					assert.NoError(t, err)
					assert.Equal(t, tc.statusCode, resp.StatusCode)
				})
			}
		})
	}
}
