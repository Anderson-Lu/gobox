package concurrent

import (
	"sync"
	"time"
)

type ConcurrentTaskDetail struct {
	fn     func(...interface{})
	params []interface{}
}

//并发任务调度器
type ConcurrentTask struct {
	wg    sync.WaitGroup
	tasks []ConcurrentTaskDetail
	max   int
	cost  int64
}

func NewConcurrentTask(maxConcurentNum int) *ConcurrentTask {
	return &ConcurrentTask{
		wg:    sync.WaitGroup{},
		tasks: make([]ConcurrentTaskDetail, 0),
		max:   maxConcurentNum,
	}
}

func (self *ConcurrentTask) Cost() int64 {
	return self.cost
}

func (self *ConcurrentTask) AddTask(task func(...interface{}), params ...interface{}) {
	self.tasks = append(self.tasks, ConcurrentTaskDetail{
		fn:     task,
		params: params,
	})
}

func (self *ConcurrentTask) Start() {
	curTaskNum := 0
	begin := time.Now().UnixNano() / 1000000
	for _, v := range self.tasks {
		self.wg.Add(1)
		curTaskNum++
		go func(v ConcurrentTaskDetail) {
			defer func() {
				self.wg.Done()
				curTaskNum--
			}()
			v.fn(v.params)
		}(v)
		for {
			if curTaskNum < self.max {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}
	}
	self.wg.Wait()
	self.cost = time.Now().UnixNano()/1000000 - begin
}
