package tickerutils

import (
	"time"

	"github.com/toddnguyen47/util-go/pkg/startstopper"
)

type Ticker2 interface {
	startstopper.Stopper
}

type impl struct {
	ticker    *time.Ticker
	idleFunc  func(t1 time.Time)
	stopChan  chan struct{}
	isStopped bool
}

func NewTicker2(duration time.Duration, idleFunc func(t1 time.Time)) Ticker2 {

	ticker := time.NewTicker(duration)

	i1 := impl{
		ticker:    ticker,
		stopChan:  make(chan struct{}, 1),
		idleFunc:  idleFunc,
		isStopped: false,
	}
	i1.start()
	return &i1
}

func (i1 *impl) start() {
	go i1.startInfLoop()
}

func (i1 *impl) startInfLoop() {
	keepRunning := true
	for keepRunning {
		select {
		case <-i1.stopChan:
			keepRunning = false
		case t1 := <-i1.ticker.C:
			i1.idleFunc(t1)
		}
	}
}

func (i1 *impl) Stop() {
	if i1.isStopped {
		return
	}
	i1.isStopped = true
	close(i1.stopChan)
	i1.ticker.Stop()
	time.Sleep(500 * time.Millisecond)
}
