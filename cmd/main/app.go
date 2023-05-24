package main

import (
	"Rotterdam/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("create router ")
	router := httprouter.New() //web-server.txt = создание роутера

	log.Println("register user handler")
	handler := user.NewHandler() // создаем хэндлер
	handler.Register(router)     // регает хэндлеры в роутере
	start(router)                // запускаем сервак
}

func start(router *httprouter.Router) {
	log.Println("start application")

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

	log.Println("server is listeting port 1234")
	log.Fatal(server.Serve(listener)) //почему log fatal????? в случае ошибок завершает работу программы с выводом сообщения об ошибке.
}
