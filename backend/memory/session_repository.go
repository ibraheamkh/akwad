package memory

import "github.com/ibraheamkh/clinicy"

func NewSessionRepository(presistantRepository clinicy.SessionService) *SessionRepository {
	return &SessionRepository{make(map[string]*clinicy.Session, 0), presistantRepository}
}

type SessionRepository struct {
	sessionsMap       map[string]*clinicy.Session
	presistantService clinicy.SessionService
}

func (s *SessionRepository) CreateSession(newSession *clinicy.Session) error {
	s.sessionsMap[newSession.SessionID] = newSession
	return s.presistantService.CreateSession(newSession)
}

func (s *SessionRepository) GetSession(sessionID string) (session *clinicy.Session, err error) {
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

func (s *SessionRepository) UpdateSession(updatedSession *clinicy.Session) error {
	s.sessionsMap[updatedSession.SessionID] = updatedSession
	return s.presistantService.UpdateSession(updatedSession)
}

func (s *SessionRepository) DestroySession(sessionID string) error {
	delete(s.sessionsMap, sessionID)
	return s.presistantService.DestroySession(sessionID)
}
