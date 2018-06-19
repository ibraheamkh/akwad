package main

import (
	"bufio"
	"log"
	"net"

	"github.com/ibraheamkh/clinicy/http"
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
	//TODO customize server arguments
	log.Println(" Running Graphql Server")
	//TODO use chi instead of gorilla mux
	go http.RunServer()
	log.Println(" Running Chat Server")
	RunChat()

}

type Client struct {
	incoming chan string
	outgoing chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.incoming <- line
	}
}

func (client *Client) Write() {
	for data := range client.outgoing {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		incoming: make(chan string),
		outgoing: make(chan string),
		reader:   reader,
		writer:   writer,
	}

	client.Listen()

	return client
}

type ChatRoom struct {
	clients  []*Client
	joins    chan net.Conn
	incoming chan string
	outgoing chan string
}

func (chatRoom *ChatRoom) Broadcast(data string) {
	for _, client := range chatRoom.clients {
		client.outgoing <- data
	}
}

func (chatRoom *ChatRoom) Join(connection net.Conn) {
	client := NewClient(connection)
	chatRoom.clients = append(chatRoom.clients, client)
	go func() {
		for {
			chatRoom.incoming <- <-client.incoming
		}
	}()
}

func (chatRoom *ChatRoom) Listen() {
	go func() {
		for {
			select {
			case data := <-chatRoom.incoming:
				chatRoom.Broadcast(data)
			case conn := <-chatRoom.joins:
				chatRoom.Join(conn)
			}
		}
	}()
}

func NewChatRoom() *ChatRoom {
	chatRoom := &ChatRoom{
		clients:  make([]*Client, 0),
		joins:    make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}

	chatRoom.Listen()

	return chatRoom
}

func RunChat() {
	chatRoom := NewChatRoom()

	listener, _ := net.Listen("tcp", ":5037")

	for {
		conn, _ := listener.Accept()
		chatRoom.joins <- conn
	}
}

// func main() {
// 	// Schema
// 	fields := graphql.Fields{
// 		"hello": &graphql.Field{
// 			Type: graphql.String,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				return "world", nil
// 			},
// 		},
// 	}
// 	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
// 	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
// 	schema, err := graphql.NewSchema(schemaConfig)
// 	if err != nil {
// 		log.Fatalf("failed to create new schema, error: %v", err)
// 	}

// 	// Query
// 	query := `
// 		{
// 			hello
// 		}
// 	`
// 	params := graphql.Params{Schema: schema, RequestString: query}
// 	r := graphql.Do(params)
// 	if len(r.Errors) > 0 {
// 		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
// 	}
// 	rJSON, _ := json.Marshal(r)
// 	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
// }

// func main() {
// 	//logs settings
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)

// 	//run mongo db server

// 	//init mongo db
// 	session, err := mongo.Dial("188.226.195.70")
// 	if err != nil {
// 		log.Printf("Error connecting to sqlite db: %v", err)
// 	}

// 	// defer session.Close()

// 	//populate db
// 	//mongo.PopulateDB(session.Copy())
// 	//TODO commented the services until really implemented

// 	CustomerRepo := &mongo.CustomerRepository{Session: session}
// 	sessionsRepo := memory.NewSessionRepository(&mongo.SessionRepository{Session: session})
// 	grpahqlRepo := &graphql.Repo{}
// 	//TODO: token service should be refactored in its own jwt package
// 	// Attach to HTTP handler.
// 	h := &http.Handler{
// 		CustomerService: customerRepo,
// 		SessionService:  sessionsRepo,
// 		GraphQLService: grpahqlRepo
// 	}

// 	//Todo fix the router to take graqhql handlers
// 	r := http.NewAPIHandler(h)

// 	log.Println("Starting the server")
// 	// start http server...
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }
