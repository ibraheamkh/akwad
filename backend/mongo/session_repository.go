package mongo

import (
	"github.com/ibraheamkh/akwad"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SessionRepository struct {
	Session *mgo.Session
}

func (s *SessionRepository) CreateSession(newSession *akwad.Session) error {
	//copy the session
	session := s.Session.Copy()

	defer session.Close()

	//collection users
	c := session.DB("akwad").C("sessions")
	err := c.Insert(newSession)
	return err
}

func (s *SessionRepository) GetSession(sessionID string) (*akwad.Session, error) {

	session := s.Session.Copy()

	defer session.Close()

	result := &akwad.Session{}

	c := session.DB("akwad").C("sessions")

	err := c.Find(bson.M{"session_id": sessionID}).One(result)
	return result, err
}

func (s *SessionRepository) UpdateSession(updatedSession *akwad.Session) error {
	//copy the session
	session := s.Session.Copy()

	defer session.Close()

	//collection users
	c := session.DB("akwad").C("sessions")
	err := c.Update(bson.M{"session_id": updatedSession.SessionID}, updatedSession)
	return err
}
func (s *SessionRepository) DestroySession(sessionID string) error {
	session := s.Session.Copy()

	defer session.Close()

	c := session.DB("akwad").C("sessions")

	return c.Remove(bson.M{"session_id": sessionID})
}
