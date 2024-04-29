package http_server

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	// 创建路由器
	router := NewRouter()

	// 创建 HTTP 服务器实例
	httpServer := &http.Server{
		Addr:    ":6317",
		Handler: router,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start(address string) error {
	log.Printf("Server listening on %s\n", address)

	// 启动 HTTP 服务器
	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
