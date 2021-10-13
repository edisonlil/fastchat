package ws

import "fastchat/domain"

//SessionManager
//The management stores all logged-in user
//information and serves as the core brain of the I.M
type SessionManager struct {
	Sessions map[*Session]bool //All Session

	NameSpaces map[string]map[string]*Session //User Session in the specified namespace

	Registrar chan *Session //Session 注册

	Logout chan *Session // Session 注销

}

var Manager *SessionManager

func InitSessionManager() *SessionManager {

	if Manager != nil {
		return Manager
	}

	Manager = &SessionManager{
		Sessions:   make(map[*Session]bool),
		NameSpaces: make(map[string]map[string]*Session),
		Registrar:  make(chan *Session, 1024),
		Logout:     make(chan *Session, 1024),
	}

	return Manager
}

func (p *SessionManager) eventRegister(session *Session) {

	//获取当前用户
	user := session.Ctx.GetCurrentUser()

	//注册到指定的命名空间
	sessions := p.NameSpaces[user.Namespace]
	if sessions == nil {
		p.NameSpaces[user.Namespace] = make(map[string]*Session, 1024)
	}
	sessions[user.Id] = session

	p.Sessions[session] = true
}

func (p *SessionManager) eventLogout(session *Session) {

	delete(p.Sessions, session)

	//获取当前用户
	user := session.Ctx.GetCurrentUser()

	//注册到指定的命名空间
	sessions := p.NameSpaces[user.Namespace]
	delete(sessions, user.Id)

	session.Close()

}

//StartListen
//Start listening for all events of the session manager.
func (p *SessionManager) StartListen() {

	for {
		select {

		case session := <-p.Registrar:
			p.eventRegister(session)
		case session := <-p.Logout:
			p.eventLogout(session)

		}

	}

}

func (p *SessionManager) getUserSession(namespace string, userId string) *Session {
	return p.NameSpaces[namespace][userId]
}

//SendMsg 发送信息
func (p *SessionManager) SendMsg(msg *domain.Message) error {
	err := p.getUserSession(msg.Namespace, msg.AcceptOpenId).WriteMsg(msg)
	if err != nil {
		return err
	}
	return err
}

//SendMsgToNamespace 发送信息
func (p *SessionManager) SendMsgToNamespace(msg *domain.Message) error {
	users := p.NameSpaces[msg.Namespace]
	for _, session := range users {
		session.WriteMsg(msg)
	}
	return nil
}
