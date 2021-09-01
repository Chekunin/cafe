package collector

import (
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/fatih/structs"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type Collector struct {
	enqueuer *work.Enqueuer
}

type NewCollectorParams struct {
	NameSpace string
	RedisPool *redis.Pool
}

func NewCollector(params NewCollectorParams) *Collector {
	enqueuer := work.NewEnqueuer(params.NameSpace, params.RedisPool)
	return &Collector{enqueuer: enqueuer}
}

func (c Collector) AddJob(jobName string, arg interface{}) error {
	_, err := c.enqueuer.Enqueue(jobName, structs.Map(arg))
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("enqueuer Enqueue jobName=%s arg=%+v", jobName, arg), err)
		return err
	}
	return nil
}
