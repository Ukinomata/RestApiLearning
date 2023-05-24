package handlers

import "github.com/julienschmidt/httprouter"

type Handler interface {
	Register(router *httprouter.Router)
} // регает все наши хендлеры в роутере
