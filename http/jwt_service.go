package http

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/ibraheamkh/clinicy"
)

//TODO: generate a key
// token stored in the cookie might change this in future
// decide on which claims to store
// set the right values when creating the claims
// verify the alg in the claim
// maybe add role to the claims ?
// find a way to prevent two devices from signing in using the same user creds
// split the package to its own package

type tokenContextKey string

var (
	//TODO fix this key
	mySigningKey = []byte("mySecret")
	adminRole    = "admin"
	userRole     = "user"
)

//Claims wraps standard claims and my custom claims
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type tokenExtractor struct{}

func (tx *tokenExtractor) ExtractToken(r *http.Request) (string, error) {

	//the name of the jwt token in the Cookie
	c, err := r.Cookie("jwt")

	if err != nil {
		log.Println("Token not found in cookie looking for token in Http Header")
		token := r.Header.Get("Authorization")
		return token, nil
	}

	return c.Value, nil
}

//Authenticator middileware wrapper for jwtauth.Authenticator
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, err := request.ParseFromRequestWithClaims(r, &tokenExtractor{}, &Claims{}, func(token *jwt.Token) (interface{}, error) {

			return mySigningKey, nil
		})

		if err != nil {
			log.Printf("Error parsing token from request: %v", err)
			return
		}

		log.Printf("sucessfully parsed token from request: %v", token.Claims)

		if token == nil || !token.Valid {
			// 	log.Printf("jwt not valid: %v, ok: %v", jwtToken.Valid, ok)
			http.Error(w, http.StatusText(401), 401)
			return
		}
		log.Println("Token authenticated")
		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

//Verifier middleware a wrapper for jwtauth.TokenAuth
func Verifier(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO :Parse the token from the context here
		//then verify it

		token, err := parseTokenFromRequest(r)

		if err != nil {
			//TODO parse the token from request hesder
			// tokenStr := r.Header.Get("Authorization")
			//
			// if tokenStr != "" {
			// 	next.ServeHTTP(w, r)
			// }

			// if strings.HasPrefix(tokenStr, "Bearer") {
			// 	next.ServeHTTP(w, r)
			// }

			log.Printf("Token is empty or invalid : %v", err)
			http.Error(w, http.StatusText(401), 401)
			return
		}

		log.Println("Inside verify, token value from cookie: ", token)

		next.ServeHTTP(w, r)
	})
	//return tokenAuth.Verifier(next)
}

//UserRoleMiddleware middileware wrapper for jwtauth.Authenticator
func UserRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, err := request.ParseFromRequestWithClaims(r, &tokenExtractor{}, &Claims{}, func(token *jwt.Token) (interface{}, error) {

			return mySigningKey, nil
		})

		if err != nil {
			log.Println("Error parsing token: ", err)
			return
		}

		role := GetRole(token.Raw)

		if role != userRole && role != adminRole {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//TokenService implement token service
type TokenService struct {
	AccountService clinicy.AccountService
}

//CreateToken creates a token for the given user
func (s *TokenService) CreateToken(user *clinicy.Account) (string, error) {

	//TODO improve this implementaion && test it

	// token := jwtauth.New("HS256", []byte(mySigningKey), nil)
	//
	// claims := jwtauth.Claims{}
	// claims["user"] = u
	// claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	//
	// _, tokenString, _ := token.Encode(claims)

	exp := time.Now().Add(time.Hour * 24).Unix()

	// Create the Claims
	claims := Claims{
		user.Email,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	log.Printf("%v %v", ss, err)

	return ss, nil

}

//GetUser gets a user based on a given token
func (s *TokenService) GetUser(tokenStr string) (*clinicy.Account, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return mySigningKey, nil
	})

	if err != nil {
		return nil, nil
	}

	if !token.Valid {
		log.Println("Not valid token")
		return nil, nil
	}

	claims := token.Claims.(jwt.MapClaims)
	log.Println("Claims", claims)
	user, err := s.AccountService.GetUserByEmail(claims["email"].(string))
	if err != nil {
		log.Println("Current User Found")
	}
	return user, err

}

func (s *TokenService) GetRole(tokenStr string) string {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return mySigningKey, nil
	})

	if err != nil {
		log.Println("Error parsing the token")
		return ""
	}

	if !token.Valid {
		log.Println("Not valid token")
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	log.Println("Role in the claims", claims["role"].(string))
	role := claims["role"].(string)
	return role
}

func GetRole(tokenStr string) string {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return mySigningKey, nil
	})

	if err != nil {
		log.Println("Error parsing the token")
		return ""
	}

	if !token.Valid {
		log.Println("Not valid token")
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	log.Println("Role in the claims", claims["role"].(string))
	role := claims["role"].(string)
	return role
}

func parseTokenFromRequest(r *http.Request) (string, error) {
	token, err := request.ParseFromRequestWithClaims(r, &tokenExtractor{}, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		return mySigningKey, nil
	})

	if err != nil {
		log.Println("Failed to parse token from request: ", err)
		return "", err
	}

	return token.Raw, nil
}
