package mongo

import (
	"log"

	"github.com/ibraheamkh/clinicy"
	mgo "gopkg.in/mgo.v2"
)

//UserRepository implements UserService Methods
type UserRepository struct {
	Session *mgo.Session
}

//User is user UserRepository implementation
func (s *UserRepository) User(id string) (*clinicy.Account, error) {

	return nil, nil
}

//Users returns all the users in the system
func (s *UserRepository) Users() (users []*clinicy.Account, err error) {

	return users, nil
}

//Admins returns all the users in the system
func (s *UserRepository) Admins() (users []*clinicy.Account, err error) {
	return nil, nil
}

//IsUser returns true if the user is in the system
func (s *UserRepository) IsUser(email, password string) bool {
	return false
}

//GetUser returns a user from a db given a username
func (s *UserRepository) GetUser(id string) (*clinicy.Account, error) {
	return nil, nil
}

//CreateUser creates one user in the db
func (s *UserRepository) CreateUser(u *clinicy.Account) error {
	log.Println("Creating user")
	//copy the session
	session := s.Session.Copy()

	defer session.Close()

	//collection users
	c := session.DB("clinicy").C("users")

	err := c.Insert(u)

	return err
}

//GetUserByEmail returns user object given an email
func (s *UserRepository) GetUserByEmail(email string) (*clinicy.Account, error) {
	return nil, nil
}

//UpdateUser updates user in the db
func (s *UserRepository) UpdateUser(u *clinicy.Account) error {

	return nil
}

//DeleteUser deletes one user using id
func (s *UserRepository) DeleteUser(id string) error {

	return nil
}
