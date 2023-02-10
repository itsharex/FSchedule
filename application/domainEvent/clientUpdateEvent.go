package domainEvent

import (
	"FSchedule/domain"
	"FSchedule/domain/client"
	"encoding/json"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/flog"
)

// ClientUpdateEvent 客户端有更新（Redis订阅）
func ClientUpdateEvent(message any, _ core.EventArgs) {
	var clientDO client.DomainObject
	err := json.Unmarshal([]byte(message.(string)), &clientDO)
	if err != nil {
		return
	}
	flog.Debugf("收到客户端（%d）更新通知：%s:%d", clientDO.Id, clientDO.Ip, clientDO.Port)
	domain.MonitorClientPush(&clientDO)
}
