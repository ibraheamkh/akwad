package http

// http logic

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ibraheamkh/clinicy"
	"github.com/ibraheamkh/clinicy/graphql"
)

//Here we handle routing and mapping http requests to functions

//--- Curren Design is that we have one Handler struct that will hold all other services

//-- Another design is to have a Handler for each service, more on this desing when we meet

//Handler represents an HTTP API interface for our app. implements ServeHTTP
type Handler struct {
	Router          *chi.Mux
	CustomerService clinicy.CustomerService
	SessionService  clinicy.SessionService
	TokenService    clinicy.TokenService
	GraphQLService  clinicy.GraphQLService
}

//ServeHTTP currently just a wrapper for chi.Mux.ServeHTTP
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}

//the router
// func NewAPIHandler(h *Handler) *Handler {
// 	r := chi.NewRouter()
// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.RealIP)
// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)

// 	// Routes with no authentication
// 	r.Route("/customer", func(r chi.Router) {
// 		r.Route("/auth", func(r chi.Router) {
// 			r.Post("/", h.customerAuthHandler)
// 			r.Route("/verify-sms", func(r chi.Router) {
// 				r.Post("/", h.customerAuthVerifySMSHandler)
// 			})
// 		})

// 		r.Route("/orders", func(r chi.Router) {
// 			r.Use(h.authorizationHandler)
// 			r.Post("/", h.customerNewOrderHandler)
// 			r.Get("/", h.customerOrdersHandler)
// 		})
// 	})

// 	h.Router = r
// 	return h
// }

func NewAPIHandler(h *Handler) *Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes with no authentication
	r.Route("/customer", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			// r.Post("/", h.customerAuthHandler)
			// r.Route("/verify-sms", func(r chi.Router) {
			// 	r.Post("/", h.customerAuthVerifySMSHandler)
			// })
		})

		r.Route("/graphql", func(r chi.Router) {
			// r.Use(h.authorizationHandler)
			r.Post("/", GetDiscs)
			// r.Get("/", h.customerOrdersHandler)
		})
	})

	//public router

	h.Router = r
	return h
}

func RunServer() {
	//move this logic to the router logic and use chi instead
	router := mux.NewRouter().StrictSlash(true)
	var handler http.Handler
	handler = http.HandlerFunc(GetDiscs)
	router.Methods("POST").Path("/").Name("GetDiscs").Handler(handler)
	//router.Methods("POST").Path("/").Handler(handler)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Accept-Language", "X-CSRF-Token", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "HEAD", "OPTIONS"})

	fmt.Println("Now server is running on port 8090")

	// TODO use this middliw ware else where, maybe middleware cahin?
	// launch server
	log.Fatal(http.ListenAndServe(":8090",
		handlers.CORS(allowedOrigins, headersOk, allowedMethods)(router)))
}

func GetDiscs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error GetDiscs", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error GetDiscs", err)
	}

	//TODO remove the Do method call to seperate package
	//we use owr own graphql package and we pass it the body

	//Do should be called by graphql service
	result := graphql.Do(body)
	json.NewEncoder(w).Encode(result)
	log.Println(http.StatusOK)
	w.WriteHeader(http.StatusOK)
	return
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	log.Printf("got request")
}
