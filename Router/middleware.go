package Router

import (
	"PicturePerfect2/Logic"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

type Handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (handler *Handlers) AuthenticateUser(h http.HandlerFunc) http.HandlerFunc {
	authenticate:= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Missing cookie", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		tknStr := c.Value
		claims := &Logic.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return Logic.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Signature Invalid", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		h(w,r)
	})
	return handler.Logger(authenticate)
}

func (handler *Handlers) AuthenticateAdmin(h http.HandlerFunc) http.HandlerFunc {
	authenticate:= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Missing cookie", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		tknStr := c.Value
		claims := &Logic.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return Logic.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Signature Invalid", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims.Role!="admin" {
			http.Error(w, "Invalid admin token", http.StatusUnauthorized)
			return
		}
		h(w,r)
	})
	return handler.Logger(authenticate)
}

func (handler *Handlers) CheckURLParams(h http.HandlerFunc) http.HandlerFunc {
	checkparams := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query().Get("pageNo")) == 0 || len(r.URL.Query().Get("pageSize")) == 0{
			http.Error(w, "Missing params", http.StatusBadRequest)
			return // don't call original handler
		}
		h(w, r)
	})
	return handler.Logger(checkparams)

}

func (handler *Handlers) Logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer handler.logger.Printf("request processed")
		h(w, r)
	}

}
