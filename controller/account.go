package controller

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Mau005/MyPet/configuration"
	"github.com/Mau005/MyPet/constants"
	"github.com/Mau005/MyPet/db"
	"github.com/Mau005/MyPet/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type ControllerAccount struct{}

func (ca *ControllerAccount) DeletedAccount(idAccount uint) (err error) {
	account, err := ca.getAccountID(idAccount)
	if err != nil {
		return
	}
	if err = db.DB.Delete(&account).Error; err != nil {
		return
	}
	return
}

func (ca *ControllerAccount) SaveAccount(account models.Account) (models.Account, error) {
	if err := db.DB.Save(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (ca *ControllerAccount) getAccountName(name string) (account models.Account, err error) {
	if err = db.DB.Where("name = ?", name).First(&account).Error; err != nil {
		return
	}
	return
}

func (ca *ControllerAccount) getAccountID(id uint) (account models.Account, err error) {
	if err = db.DB.Where("id = ?", id).First(&account).Error; err != nil {
		return
	}
	return
}

func (ca *ControllerAccount) getAccountEmail(email string) (account models.Account, err error) {
	if err = db.DB.Where("email = ?", email).First(&account).Error; err != nil {
		return
	}
	return
}

func (ca *ControllerAccount) GetAccount(content string) (account models.Account, err error) {
	account, err = ca.getAccountName(content)
	if err != nil {
		account, err = ca.getAccountEmail(content)
		return
	}
	return
}

func (ca *ControllerAccount) ChangePassword(name, passwordOld, password, passwordTwo string) (account models.Account, err error) {
	account, err = ca.GetAccount(name)
	if err != nil {
		return
	}

	if !(len(passwordTwo) >= constants.LEN_PASSWORD) {
		return account, errors.New(constants.ERROR_LEN_PASSWORD)
	}

	err = ca.CompareCryptPassword(account.Password, passwordOld)
	if err != nil {
		return
	}

	account.Password = ca.GenerateCryptPassword(password)

	account, err = ca.SaveAccount(account)
	if err != nil {
		return
	}

	return
}

func (ca *ControllerAccount) CreateAccount(account models.Account, passwordTwo *string) (models.Account, error) {
	if passwordTwo != nil {
		if !(account.Password == *passwordTwo) {
			return account, errors.New("error compare password")
		}
	}
	if !(len(account.Name) >= constants.LEN_ACCOUNT_NAME) {
		return account, errors.New("error len account name")
	}
	if !(len(account.Password) >= constants.LEN_PASSWORD) {
		return account, errors.New("error len password name")
	}
	if !(strings.Contains(account.Email, "@")) {
		return account, errors.New("error email validation")
	}
	// encryp password
	account.Password = ca.GenerateCryptPassword(account.Password)
	if err := db.DB.Create(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (ca *ControllerAccount) GenerateCryptPassword(password string) string {
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hasedPassword)
}

func (ca *ControllerAccount) CompareCryptPassword(password, passwordTwo string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordTwo))
}

func (ca *ControllerAccount) NewClaim(account models.Account) *models.Claims {
	expirationTime := time.Now().Add(constants.EXPIRATION_TOKEN * time.Hour)
	return &models.Claims{
		AccountName:  account.Name,
		AccountEmail: account.Email,
		Access:       account.Access,
		AccountID:    account.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
}

func (ca *ControllerAccount) GenerateToken(account models.Account) (tokenString string, err error) {

	claims := ca.NewClaim(account)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(configuration.Security))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func (ca *ControllerAccount) SaveSession(tokenString *string, w http.ResponseWriter, r *http.Request) {
	session, _ := configuration.Store.Get(r, constants.ACCOUNT_SESSION)
	if tokenString == nil {
		session.Values["token"] = nil
	} else {
		session.Values["token"] = *tokenString
	}
	session.Save(r, w)
}

func (ca *ControllerAccount) AuthenticateJWT(tokenSession string) error {

	token, err := jwt.Parse(tokenSession, func(token *jwt.Token) (interface{}, error) {
		return []byte(configuration.Security), nil
	})

	if err != nil || !token.Valid {
		return err
	}

	return nil
}

func (ca *ControllerAccount) GetSessionClaims(r *http.Request) (*models.Claims, error) {
	claims := &models.Claims{}
	session, err := configuration.Store.Get(r, constants.ACCOUNT_SESSION)
	if err != nil {
		return claims, err
	}

	token, ok := session.Values["token"].(string)
	if !ok {
		return claims, errors.New(constants.ERROR_GET_TOKEN)
	}
	tokenKey := []byte(configuration.Security)
	tokenParser := jwt.Parser{}

	_, err = tokenParser.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		return claims, errors.New(constants.ERROR_TOKEN)
	}
	return claims, nil
}

func (ca *ControllerAccount) GetSessionUser(r *http.Request) (models.Account, error) {
	claims, err := ca.GetSessionClaims(r)
	if err != nil {
		return models.Account{}, err
	}

	user, err := ca.GetAccount(claims.AccountName)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ca *ControllerAccount) LoginAccount(name, password string) (account models.Account, token string, err error) {

	account, err = ca.GetAccount(name)
	if err != nil {
		return
	}

	err = ca.CompareCryptPassword(account.Password, password)
	if err != nil {
		err = errors.New(constants.ERROR_ACCESS_CREDENTIALS)
		return
	}

	token, err = ca.GenerateToken(account)
	if err != nil {
		return
	}

	return account, token, nil
}
