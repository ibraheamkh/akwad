package mongo

import (
	"log"

	"github.com/ibraheamkh/clinicy"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type CustomerRepository struct {
	Session *mgo.Session
}

//CreateCustomerAccount creates a customer account in the db
func (s *CustomerRepository) CreateCustomerAccount(newCustomer *clinicy.Customer) error {
	log.Println("Creating customer account")
	//copy the session
	newCustomer.ID = bson.NewObjectId().Hex()
	session := s.Session.Copy()

	defer session.Close()

	//collection users
	c := session.DB("clinicy").C("customers")

	err := c.Insert(newCustomer)

	return err
}

func (s *CustomerRepository) GetCustomerByMobile(mobile string) (*clinicy.Customer, error) {

	session := s.Session.Copy()

	defer session.Close()

	result := &clinicy.Customer{}
	//collection users
	c := session.DB("clinicy").C("customers")

	err := c.Find(bson.M{"account.mobile": mobile}).One(result)
	return result, err
}
