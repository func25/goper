package goper

import (
	"sync"
	"time"

	"github.com/func25/gop"
	"github.com/func25/mafu"
)

// paraller is in development and can be broken anytime
type paraller struct {
	// batch state
	works    []func() error
	w        sync.WaitGroup
	batchNum int

	// execution report
	elapsed time.Duration

	// error state, report
	errors      []funcErr
	errChan     chan *funcErr
	stopWhenErr bool
}

func Paraller(batch int) *paraller {
	return &paraller{batchNum: batch, errChan: make(chan *funcErr)}
}

func (p *paraller) AddWorks(funcs ...func() error) *paraller {
	p.works = append(p.works, funcs...)
	return p
}

func (p *paraller) Clear() *paraller {
	p.works = make([]func() error, 0)

	return p
}

func (p *paraller) StopWhenError(stop bool) *paraller {
	p.stopWhenErr = stop

	return p
}

// Execute executes the functions parallelly but in batches
func (p *paraller) Execute() []funcErr {
	p.startReport()
	defer p.stopReport()

	for i := 0; i < len(p.works); {
		des := len(p.works)
		if p.batchNum > 0 {
			des = mafu.Min(des, i+p.batchNum)
		}

		for ; i < des; i++ {
			p.executeFunc(i)
		}

		p.w.Wait()

		if p.stopWhenErr && len(p.errors) > 0 {
			return p.errors
		}
	}

	return p.errors
}

// startReport starts a goroutine to report the execution time and errors
func (p *paraller) startReport() {
	go gop.SafeGo(func() {
		start := time.Now()
		p.errors = nil

		for err := range p.errChan {
			if err == nil {
				p.elapsed = time.Since(start)
				return
			}
			p.errors = append(p.errors, *err)
		}
	})
}

func (p *paraller) stopReport() {
	p.errChan <- nil
}

func (p *paraller) executeFunc(i int) {
	p.w.Add(1)
	go p.doJob(i)
}

func (p *paraller) doJob(i int) error {
	defer p.w.Done()

	if err := p.works[i](); err != nil {
		p.errChan <- &funcErr{
			Func: p.works[i],
			Err:  err,
		}
	}

	return nil
}

func (p *paraller) Report() time.Duration {
	return p.elapsed
}
