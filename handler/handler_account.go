package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mau005/MyPet/constants"
	"github.com/Mau005/MyPet/controller"
	"github.com/Mau005/MyPet/models"
)

type HandlerAccount struct{}

func (hl *HandlerAccount) Login(w http.ResponseWriter, r *http.Request) {
	var exceptCtl controller.ControllerException
	name := r.FormValue("name")
	password := r.FormValue("password")
	var accountCtl controller.ControllerAccount

	account, token, err := accountCtl.LoginAccount(name, password)
	if err != nil {
		exceptCtl.NewException(w, "Module Login", err.Error(), http.StatusNotAcceptable, nil)
		return
	}

	accountCtl.SaveSession(&token, w, r)
	json.NewEncoder(w).Encode(
		struct {
			Status  int    `json:"Status"`
			Message string `json:"Message"`
			Account models.Account
		}{
			Status:  http.StatusOK,
			Message: fmt.Sprintf("has been entered without problems %s", account.Name),
			Account: account,
		},
	)

}

func (h1 *HandlerAccount) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var exceptCtl controller.ControllerException
	var accountCtl controller.ControllerAccount
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordTwo := r.FormValue("passwordTwo")

	account, err := accountCtl.CreateAccount(models.Account{
		Name:     name,
		Password: password,
		Email:    email,
	}, &passwordTwo)

	if err != nil {
		exceptCtl.NewException(w, "Module Create Account", err.Error(), http.StatusNotAcceptable, nil)
		return
	}

	json.NewEncoder(w).Encode(
		struct {
			Status  int    `json:"Status"`
			Message string `json:"Message"`
			Account models.Account
		}{
			Status:  http.StatusOK,
			Message: "the account has been created successfully",
			Account: account,
		},
	)
}

func (hl *HandlerAccount) Logout(w http.ResponseWriter, r *http.Request) {
	var accountCtl controller.ControllerAccount
	accountCtl.SaveSession(nil, w, r)
	json.NewEncoder(w).Encode(
		struct {
			Status  int    `json:"Status"`
			Message string `json:"Message"`
		}{
			Status:  http.StatusOK,
			Message: "account has logged out",
		},
	)
}

func (hl *HandlerAccount) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var accountCtl controller.ControllerAccount
	var exceptCtl controller.ControllerException

	claim, err := accountCtl.GetSessionClaims(r)
	if err != nil {
		return
	}

	responseClient := struct {
		Name        string `json:"name"`
		PasswordOld string `json:"passwordOld"`
		PasswordNew string `json:"passwordNew"`
		PasswordTwo string `json:"passwordTwo"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&responseClient)
	if err != nil {
		exceptCtl.NewException(w, "Module Change Password", err.Error(), http.StatusConflict, nil)
		return
	}

	if len(responseClient.Name) >= 1 {
		if !(claim.Access >= constants.PRIVILEGES_ADMINISTRATOR) {
			exceptCtl.NewException(w, "Module Change Password", constants.ERROR_UNHAUTORIZED, http.StatusUnauthorized, nil)
			return
		}
	} else {
		responseClient.Name = claim.AccountName
	}

	account, err := accountCtl.ChangePassword(responseClient.Name, responseClient.PasswordOld, responseClient.PasswordNew, responseClient.PasswordTwo)
	if err != nil {
		exceptCtl.NewException(w, "Module Change Password", err.Error(), http.StatusConflict, nil)
		return
	}

	json.NewEncoder(w).Encode(
		struct {
			Status  int    `json:"Status"`
			Message string `json:"Message"`
		}{
			Status:  http.StatusOK,
			Message: fmt.Sprintf("the user's password has been changed: %s", account.Name),
		},
	)
}
