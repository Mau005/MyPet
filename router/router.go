package router

import (
	"github.com/Mau005/MyPet/handler"
	"github.com/Mau005/MyPet/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	var handlerAccount handler.HandlerAccount
	// Public Router:
	r.Use(middleware.CommonMiddleware)
	r.HandleFunc("/account/login", handlerAccount.Login).Methods("POST")
	r.HandleFunc("/account/create_account", handlerAccount.CreateAccount).Methods("POST")
	r.HandleFunc("/account/logout", handlerAccount.Logout)

	// Private Router:
	s := r.PathPrefix("/api/v1/auth").Subrouter()
	s.Use(middleware.CommonMiddleware)  //content api json
	s.Use(middleware.SessionMiddleware) // security
	s.HandleFunc("/account/change_password", handlerAccount.ChangePassword).Methods("PUT")
	return r
}
