package mongo

import (
	"log"

	"github.com/ibraheamkh/akwad"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type CustomerRepository struct {
	Session *mgo.Session
}

//CreateCustomerAccount creates a customer account in the db
func (s *CustomerRepository) CreateCustomerAccount(newCustomer *akwad.Customer) error {
	log.Println("Creating customer account")
	//copy the session
	newCustomer.ID = bson.NewObjectId().Hex()
	session := s.Session.Copy()

	defer session.Close()

	//collection users
	c := session.DB("akwad").C("customers")

	err := c.Insert(newCustomer)

	return err
}

func (s *CustomerRepository) GetCustomerByMobile(mobile string) (*akwad.Customer, error) {

	session := s.Session.Copy()

	defer session.Close()

	result := &akwad.Customer{}
	//collection users
	c := session.DB("akwad").C("customers")

	err := c.Find(bson.M{"account.mobile": mobile}).One(result)
	return result, err
}
