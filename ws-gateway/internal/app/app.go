package app

import (
	"log"
	"net/http"
	"ws-gateway/internal/config"
	"ws-gateway/internal/handler"
	"ws-gateway/internal/hub"

	chat_service "ws-gateway/internal/service/chat"
	jwt_service "ws-gateway/internal/service/jwt/impl"
)

func StartApp() {
	env := config.LoadEnv()
	h := hub.NewHub()

	jwtService := jwt_service.NewJWTService(env.GetSecret(), env.Exp)
	chatService := chat_service.NewChatService(h)

	wsHandler := handler.NewWSHandler(h, chatService, jwtService)

	go h.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsHandler.HandleWS)

	srv := &http.Server{
		Addr: ":8001",
	}

	go func() {
		log.Println("WS server running on :8001")

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	log.Println("WS server running on :8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
