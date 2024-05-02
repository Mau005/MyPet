package middleware

import (
	"net/http"

	"github.com/Mau005/MyPet/configuration"
	"github.com/Mau005/MyPet/constants"
	"github.com/Mau005/MyPet/controller"
)

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var exceptionCtl controller.ControllerException
		session, err := configuration.Store.Get(r, constants.ACCOUNT_SESSION)
		if err != nil {
			http.Error(w, "Error al obtener la sesión", http.StatusInternalServerError)
			return
		}
		var accountCtl controller.ControllerAccount
		if tokenStr, ok := session.Values["token"].(string); !ok {
			exceptionCtl.NewException(w, "Module Token", "error get token", http.StatusNetworkAuthenticationRequired, nil)
			return
		} else {
			err = accountCtl.AuthenticateJWT(tokenStr)
			if err != nil {
				exceptionCtl.NewException(w, "Module Token", "error validation token", http.StatusNetworkAuthenticationRequired, nil)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
