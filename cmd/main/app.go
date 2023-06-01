package main

import (
	"Rotterdam/internal/user"
	"Rotterdam/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New() //web-server.txt = создание роутера

	logger.Info("register user handler")
	handler := user.NewHandler(logger) // создаем хэндлер
	handler.Register(router)           // регает хэндлеры в роутере
	start(router)                      // запускаем сервак
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", ":1234") //код для запуска сервера с помощью пакета net
	// 0.0.0.0 все ip и все интерфейсы
	//:1234 порт
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,           //передали созданный роутер
		WriteTimeout: 15 * time.Second, // указать таймауты на запись //зависит от приложения и контекста и тд
		ReadTimeout:  15 * time.Second, // указать таймауты на чтение
	}

	logger.Info("server is listeting port 1234")
	logger.Fatal(server.Serve(listener)) //почему log fatal????? в случае ошибок завершает работу программы с выводом сообщения об ошибке.
}
