package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mau005/MyPet/controller"
	"github.com/Mau005/MyPet/models"
)

type HandlerAccount struct{}

func (hl *HandlerAccount) Login(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Println(name, password)

}

func (h1 *HandlerAccount) CreateAccount(w http.ResponseWriter, r *http.Request) {
	exceptionCtl := controller.ControllerException{
		ResponseWriter: w,
	}
	var accountCtl controller.ControllerAccount
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordTwo := r.FormValue("passwordTwo")

	// check password account
	if !(password == passwordTwo) {
		except := exceptionCtl.NewException("Module Create Account", "error in password compare", http.StatusConflict, nil)
		json.NewEncoder(w).Encode(except)
		return
	}

	account, err := accountCtl.CreateAccount(models.Account{
		Name:     name,
		Password: password,
		Email:    email,
	})

	if err != nil {
		except := exceptionCtl.NewException("Module Create Account", err.Error(), http.StatusNotAcceptable, nil)
		json.NewEncoder(w).Encode(except)
		return
	}

	/*
		Logica de seguridad pendiente
		TODO: gestionar
	*/
	json.NewEncoder(w).Encode(
		struct {
			Status  int    `json:"Status"`
			Message string `json:"Message"`
		}{
			Status:  http.StatusAccepted,
			Message: fmt.Sprintf("Se ha creado la cuenta con exito! %s", account.Name),
		},
	)
}
