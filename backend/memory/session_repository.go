package memory

import "github.com/ibraheamkh/akwad"

func NewSessionRepository(presistantRepository akwad.SessionService) *SessionRepository {
	return &SessionRepository{make(map[string]*akwad.Session, 0), presistantRepository}
}

type SessionRepository struct {
	sessionsMap       map[string]*akwad.Session
	presistantService akwad.SessionService
}

func (s *SessionRepository) CreateSession(newSession *akwad.Session) error {
	s.sessionsMap[newSession.SessionID] = newSession
	return s.presistantService.CreateSession(newSession)
}

func (s *SessionRepository) GetSession(sessionID string) (session *akwad.Session, err error) {
	// searches for a given session in memory, and if not found then search in the presistant storage

	session = s.sessionsMap[sessionID]
	if session == nil {
		session, err = s.presistantService.GetSession(sessionID)
		if err != nil {
			return session, err
		}
	}
	if session != nil {
		s.sessionsMap[session.SessionID] = session
	}

	return
}

func (s *SessionRepository) UpdateSession(updatedSession *akwad.Session) error {
	s.sessionsMap[updatedSession.SessionID] = updatedSession
	return s.presistantService.UpdateSession(updatedSession)
}

func (s *SessionRepository) DestroySession(sessionID string) error {
	delete(s.sessionsMap, sessionID)
	return s.presistantService.DestroySession(sessionID)
}
