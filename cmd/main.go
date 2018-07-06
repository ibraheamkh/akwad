package main

import (
	"log"

	"github.com/ibraheamkh/akwad/backend/memory"
	"github.com/ibraheamkh/akwad/backend/mongo"
	"github.com/ibraheamkh/akwad/graphql"
	"github.com/ibraheamkh/akwad/http"
)

//until this moment this is just a testing server
//this will be the main server for clinicy - or it should be

//Important TODO // refactor to make main method simple and just to glue things up
//  we want http package handle only http logic
// grpahql package will handle graphql problems, init, Do etc, may be acceble throug services

// for graphql we need schema and query
// schema to describe the data
// query to fetch the data
// each query has a resovle method ro get the needed information
// heavey weight on backend but fun and interesting -- can be mixed with exsisting tech
// front end developers will develop faster than before
// this is testing side project wich will  contain alot of ideas
// containers, cloud, AI, Bockchain, pentesting, anything litiralyy

func main() {
	//logs settings
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//run docker containers we need for this application
	//mongoURL := docker.GetMongoURL()

	//run mongo db server

	//init mongo db
	session, err := mongo.Dial("localhost")
	if err != nil {
		log.Printf("Error connecting to sqlite db: %v", err)
	}

	// defer session.Close()

	//populate db
	//mongo.PopulateDB(session.Copy())
	//TODO commented the services until really implemented

	customerRepo := &mongo.CustomerRepository{Session: session}
	sessionsRepo := memory.NewSessionRepository(&mongo.SessionRepository{Session: session})
	grpahqlRepo := &graphql.Repo{}
	//TODO: token service should be refactored in its own jwt package
	// Attach to HTTP handler.
	h := &http.Handler{
		CustomerService: customerRepo,
		SessionService:  sessionsRepo,
		GraphQLService:  grpahqlRepo,
	}

	//Todo fix the router to take graqhql handlers - no run different server first and keep this
	r := http.NewAPIHandler(h)

	log.Println(" Running Graphql Server")
	//TODO use chi instead of gorilla mux
	go http.RunGraphQLServer()

	log.Println("Starting the server")
	// start http server...
	log.Fatal(http.ListenAndServe(":8080", r))
}

// func main() {
// 	//TODO customize server arguments
// 	log.Println(" Running Graphql Server")
// 	//TODO use chi instead of gorilla mux
// 	go http.RunGraphQLServer()
// 	log.Println(" Running Chat Server")
// 	RunChat()

// }
