package akwad

import "github.com/graphql-go/graphql"

//UserService provides the basic methods to retrieve a user;
//for example if we have different mongodb and postgres we will implement this interface for each
type AccountService interface {
	CreateUser(u *Account) error
	//Returns a user by ID
	User(id string) (*Account, error)
	//Returns a user by username
	GetUser(username string) (*Account, error)
	//Returns all users in the system
	Users() ([]*Account, error)
	//Creates new user
	//DeleteUser(id int) error
	//Returns a true if the username and password exists in the db
	//This method should hash the given password and compares the hashes
	IsUser(username, password string) bool
	//Get user by email
	GetUserByEmail(email string) (*Account, error)
	//Updates the user
	UpdateUser(u *Account) error
	//Return a list of admin users
	Admins() ([]*Account, error)

	DeleteUser(id string) error
}

type CustomerService interface {
	CreateCustomerAccount(*Customer) error
	GetCustomerByMobile(mobile string) (*Customer, error)
}

type OrderService interface {
	// createes a new order for a customer
	CreateOrder(*Order) error
	GetOrderByCustomerID(customerID string) error
}

// TokenService provides general operations on JWT Tokens
type TokenService interface {
	//create token for the givin user
	CreateToken(u *Account) (string, error)
	//get user out of token
	GetUser(token string) (*Account, error)
	//Returns a user role
	GetRole(tokenStr string) string
	//TODO add this and cancel the others
	//WithClaims(map[string]string) (string, error)
}

type SessionService interface {
	CreateSession(*Session) error
	GetSession(sessionID string) (*Session, error)
	UpdateSession(*Session) error
	DestroySession(sessionID string) error
}

//this will handle all requests to graphql services
//FIXME: this must be changed based on graphql concepts
//GraphQLService
type GraphQLService interface {
	Do(body []byte) *graphql.Result
	// Query(query string) error
	// CreateDoctor(d *Doctor) error
	// CreateClinic(c *Clinic) error
	// CreateAppointment(a *Appointment) error
	// CreateBooking(b *Booking) error
}
