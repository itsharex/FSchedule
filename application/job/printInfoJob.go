package job

import (
	"FSchedule/domain"
	"FSchedule/domain/serverNode"
	"fmt"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/tasks"
	"strings"
)

// PrintInfoJob 打印客户端、任务组信息
func PrintInfoJob(context *tasks.TaskContext) {
	flog.Printf("%s个客户端（%s个正常），%s个任务组（%s个运行中）\n", flog.Red(domain.ClientCount()), flog.Green(domain.ClientNormalCount()), flog.Red(domain.TaskGroupCount()), flog.Green(domain.TaskGroupEnableCount()))
	lst := container.Resolve[serverNode.Repository]().ToList()

	// 主节点
	leaderNode := lst.Where(func(item serverNode.DomainObject) bool {
		return item.IsLeader
	}).OrderByDescending(func(item serverNode.DomainObject) any {
		return item.ActivateAt
	}).First()
	if leaderNode.Id > 0 {
		flog.Printf("%s个Master节点：%s %s:%d\n", flog.Red(1), flog.Green(leaderNode.Id), flog.Yellow(leaderNode.Ip), leaderNode.Port)
	}

	// 从节点
	var serverNodes []string
	lst.Where(func(item serverNode.DomainObject) bool {
		return !item.IsLeader
	}).OrderByDescending(func(item serverNode.DomainObject) any {
		return item.Id
	}).Select(&serverNodes, func(item serverNode.DomainObject) any {
		return fmt.Sprintf("%s %s:%d", flog.Blue(item.Id), item.Ip, item.Port)
	})
	if len(serverNodes) > 0 {
		flog.Printf("%s个Slave节点：%s\n", flog.Red(len(serverNodes)), strings.Join(serverNodes, ","))
	}
	flog.Println("---------------------------------------")
}
