package corn

import (
	"fastchat/domain"
	"fastchat/servers/ws"
	"fastchat/service"
	"time"
)

//healthy
//Session 健康检查
func healthy() {

	//session 健康检测
	//5分钟进行一次健康检测 （后续改为配置文件方式）
	time.AfterFunc(time.Duration(time.Minute*5), func() {
		//获取当前时间戳
		now := time.Now().Unix()
		for session := range ws.Manager.Sessions {
			//检测时间 + 5分钟
			healthy := session.Healthy + ((1000 * 60) * 5)
			if now > healthy {
				user := service.GetUserById(session.UserId)
				//If the heartbeat time exceeds 5 minutes, it is considered offline.
				user.Status = domain.Offline
			}
		}

	})

}
