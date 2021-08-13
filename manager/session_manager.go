package manager

import "fastchat/ws"

//SessionManager
//The management stores all logged-in user
//information and serves as the core brain of the I.M

type SessionManager struct {
	Sessions map[*ws.Session]bool //All Session

	Users map[string]*ws.Session // Login Users

	Registrar chan *ws.Session //Session 注册

	Logout chan *ws.Session // Session 注销

}

func NewSessionManager() *SessionManager {

	return &SessionManager{
		Sessions:  make(map[*ws.Session]bool),
		Users:     make(map[string]*ws.Session),
		Registrar: make(chan *ws.Session, 1024),
	}
}

func (p SessionManager) eventRegister(id string, session *ws.Session) {
	p.Users[id] = session
	p.Sessions[session] = true
}

func (p SessionManager) eventLogout(id string, session *ws.Session) {
	p.Users[id] = session
	p.Sessions[session] = true
}

//StartListen
//Start listening for all events of the session manager.
func (p SessionManager) StartListen() {

	for {
		select {

		case session := <-p.Registrar:
			p.eventRegister(session.UserId, session)
		case session := <-p.Logout:
			p.eventLogout(session.UserId, session)

		}

	}

}
