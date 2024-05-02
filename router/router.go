package router

import (
	"github.com/Mau005/MyPet/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	var handlerAccount handler.HandlerAccount
	r.HandleFunc("/login", handlerAccount.Login)
	r.HandleFunc("/create_account", handlerAccount.CreateAccount)
	return r
}
