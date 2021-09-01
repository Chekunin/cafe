package dispatcher

import (
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type Dispatcher struct {
	workerPool *work.WorkerPool
}

// он будет доставать задачу из редиса, передавать воркеру, который вызовет нужную функцию

type NewDispatcherParams struct {
	NameSpace      string
	RedisPool      *redis.Pool
	MaxConcurrency uint
	JobName        string
	WorkerFunc     func(*work.Job) error
}

type Context struct {
	customerID int64
}

func NewDispatcher(params NewDispatcherParams) *Dispatcher {
	workerPool := work.NewWorkerPool(Context{}, params.MaxConcurrency, params.NameSpace, params.RedisPool)
	workerPool.Job(params.JobName, params.WorkerFunc)

	dispatcher := Dispatcher{workerPool: workerPool}

	return &dispatcher
}

func (d *Dispatcher) Start() {
	d.workerPool.Start()
}

func (d *Dispatcher) Stop() {
	d.workerPool.Stop()
}
