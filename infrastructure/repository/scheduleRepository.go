package repository

import (
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/redis"
	"time"
)

type scheduleRepository struct {
	*redis.Client
}

func (receiver *scheduleRepository) NewLock(name string) core.ILock {
	return receiver.Lock.GetLocker("FSS_Scheduler:"+name, 5*time.Second)
}
