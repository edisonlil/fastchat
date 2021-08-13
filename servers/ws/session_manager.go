package ws

//SessionManager
//The management stores all logged-in user
//information and serves as the core brain of the I.M
type SessionManager struct {
	Sessions map[*Session]bool //All Session

	Users map[string]*Session // Login Users

	Registrar chan *Session //Session 注册

	Logout chan *Session // Session 注销

}

func NewSessionManager() *SessionManager {

	return &SessionManager{
		Sessions:  make(map[*Session]bool),
		Users:     make(map[string]*Session),
		Registrar: make(chan *Session, 1024),
		Logout:    make(chan *Session),
	}
}

func (p SessionManager) eventRegister(id string, session *Session) {
	p.Users[id] = session
	p.Sessions[session] = true
}

func (p SessionManager) eventLogout(id string, session *Session) {
	delete(p.Users, id)
	delete(p.Sessions, session)
	session.Close()

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

func (p SessionManager) getUserSession(userId string) *Session {
	return p.Users[userId]
}

//SendMsg
//发送信息
func (p SessionManager) SendMsg(msg Message) error {

	err := p.getUserSession(msg.TargetId).WriteMsg(msg)
	if err != nil {
		return err
	}
	return err
}
