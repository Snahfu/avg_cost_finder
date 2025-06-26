package logic

import "sync"

// SafeWaitGroup hanya alias agar mudah digunakan (optional saja)
type SafeWaitGroup struct {
	WaitGroup sync.WaitGroup
}

func (swg *SafeWaitGroup) Add(delta int) {
	swg.WaitGroup.Add(delta)
}

func (swg *SafeWaitGroup) Done() {
	swg.WaitGroup.Done()
}

func (swg *SafeWaitGroup) Wait() {
	swg.WaitGroup.Wait()
}
