package taskGroup

import (
	"github.com/farseer-go/collections"
)

type Repository interface {
	// ToEntity 获取任务组信息
	ToEntity(id int64) DomainObject
	// ToList 获取所有任务组中的任务
	ToList() collections.List[DomainObject]
	// ToListByName 获取所有任务组中的任务
	ToListByName(name string) collections.List[DomainObject]
	// Save 保存任务组信息
	Save(do DomainObject)
	// SaveAndTask 保存任务组、任务信息
	SaveAndTask(do DomainObject)
	// SaveTask 保存任务信息
	SaveTask(taskEO TaskEO)
	// GetTask 获取任务信息
	GetTask(taskGroupId int64, taskId int64) TaskEO
	// ToTaskSpeedList 当前任务组下所有任务的执行速度
	ToTaskSpeedList(taskGroupId int64) []int64
	// ToFinishList 获取指定任务组执行成功的任务列表
	ToFinishList(id int64, top int) collections.List[TaskEO]
	// ClearFinish 清除成功的任务记录（1天前）
	ClearFinish(id int64, taskId int)
	// Sync 同步任务组数据
	Sync()
}
