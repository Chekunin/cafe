package collector

type ICollector interface {
	AddJob(jobName string, arg interface{}) error
}
