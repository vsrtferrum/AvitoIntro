package workers

import (
	"github.com/vsrtferrum/AvitoIntro/internal/model"
	"github.com/vsrtferrum/AvitoIntro/internal/module"
)

type WorkerPool struct {
	tasks chan func()
}

func NewWorkerPool(workerCount, queueSize int) *WorkerPool {
	wp := &WorkerPool{
		tasks: make(chan func(), queueSize),
	}

	for i := 0; i < workerCount; i++ {
		go wp.worker()
	}
	return wp
}

func (wp *WorkerPool) worker() {
	for task := range wp.tasks {
		task()
	}
}

func (wp *WorkerPool) Submit(task func()) {
	wp.tasks <- task
}

type ConcurrentModule struct {
	Module *module.Module
	wp     *WorkerPool
}

func NewConcurrentModule(module *module.Module, workerCount, queueSize int) *ConcurrentModule {
	return &ConcurrentModule{
		Module: module,
		wp:     NewWorkerPool(workerCount, queueSize),
	}
}

func (cm *ConcurrentModule) Auth(name, password string) (string, error) {
	type result struct {
		token string
		err   error
	}

	resChan := make(chan result, 1)

	cm.wp.Submit(func() {
		token, err := cm.Module.Auth(name, password)
		resChan <- result{token, err}
	})

	res := <-resChan
	return res.token, res.err
}
func (cm *ConcurrentModule) Identify(token string) bool {
	type result struct {
		ok bool
	}

	resChan := make(chan result, 1)

	cm.wp.Submit(func() {
		ok := cm.Module.Identify(token)
		resChan <- result{ok}
	})

	res := <-resChan
	return res.ok
}

func (cm *ConcurrentModule) Buy(token string, itemName string) error {
	errChan := make(chan error, 1)

	cm.wp.Submit(func() {
		errChan <- cm.Module.Buy(token, itemName)
	})

	return <-errChan
}

func (cm *ConcurrentModule) GetInfo(token string) (*model.InfoResponse, error) {
	type result struct {
		info *model.InfoResponse
		err  error
	}

	resChan := make(chan result, 1)

	cm.wp.Submit(func() {
		info, err := cm.Module.GetInfo(token)
		resChan <- result{info, err}
	})

	res := <-resChan
	return res.info, res.err
}

func (cm *ConcurrentModule) SendMoney(token string, receiverName string, amount uint64) error {
	errChan := make(chan error, 1)

	cm.wp.Submit(func() {
		errChan <- cm.Module.SendMoney(token, receiverName, amount)
	})

	return <-errChan
}
