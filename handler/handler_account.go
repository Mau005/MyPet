package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mau005/MyPet/constants"
	"github.com/Mau005/MyPet/controller"
)

type HandlerAccount struct{}

func (hl *HandlerAccount) Login(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Println(name, password)

}

func (h1 *HandlerAccount) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var exceptionCtl controller.ControllerException
	name := r.FormValue("name")
	password := r.FormValue("password")
	passwordTwo := r.FormValue("passwordTwo")

	if !(len(password) > constants.LEN_PASSWORD && password == passwordTwo && len(name) >= constants.LEN_ACCOUNT_NAME) {
		except := exceptionCtl.NewException("Error", "Mensaje", http.StatusConflict, nil)
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
			Message: "Se ha creado la cuenta con exito!",
		},
	)
}
