package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// Создаёт обратный прокси для заданного target (целевого сервера)
// и удаляет указанный префикс пути перед проксированием запросов.
func NewProxy(target string, stripPrefix string) gin.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		panic("invalid target URL: " + target)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// переопределяем Director чтобы убрать префикс пути
	originalDirector := proxy.Director

	// убедимся, что stripPrefix начинается с "/"
	if stripPrefix != "" && !strings.HasPrefix(stripPrefix, "/") {
		stripPrefix = "/" + stripPrefix
	}

	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// TrimPrefix от пути и нормализация результата
		req.URL.Path = strings.TrimPrefix(req.URL.Path, stripPrefix)
		if req.URL.Path == "" {
			req.URL.Path = "/"
		} else {
			// удалить возможные двойные слэши и нормализовать
			req.URL.Path = path.Clean("/" + req.URL.Path)
		}
		req.Header.Set("X-Forwarded-Host", req.Host)

	}

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
