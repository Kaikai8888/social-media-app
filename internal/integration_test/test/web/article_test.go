package article_test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"webook/internal/integration_test/startup"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticleTestSuite struct {
	suite.Suite

	server *gin.Engine
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}

func (s *ArticleTestSuite) SetupSuite() {
	s.server = startup.InitApiServer()
}

func (s *ArticleTestSuite) TestCreateArticle() {
	testCases := []struct {
		name    string
		method  string
		path    string
		token   string
		reqBody string

		wantStatus   int
		wantResponse string
	}{
		{
			name:         "Success Case: create article",
			method:       "POST",
			path:         "/articles",
			token:        "mock-token",
			reqBody:      `{"title": "test", "content": "test"}`,
			wantStatus:   200,
			wantResponse: `{"status": "success"}`,
		}, {},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.reqBody))
			req.Header.Set("Authorization", "Bearer "+tc.token)
			w := httptest.NewRecorder()
			s.server.ServeHTTP(w, req)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			assert.Equal(t, tc.wantStatus, resp.StatusCode)
			assert.JSONEq(t, tc.wantResponse, string(body))
		})
	}

}
