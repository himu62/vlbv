package web

import (
	"net/http"

	"github.com/himu62/vlbv/server/log"
)

// Web はhttpサーバのモデル
type Web struct {
	server *http.Server
	logger *log.Logger
}

// NewWebServer はWebインスタンスを生成する
func NewWebServer(logger *log.Logger) *Web {
	return &Web{
		server: &http.Server{},
		logger: logger,
	}
}

// Listen はWebサーバを起動してlistenする
func (s *Web) Listen(port string) {
	s.logger.Info("server listening on", port)
	s.server.ListenAndServe()
}
