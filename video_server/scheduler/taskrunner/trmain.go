package taskrunner

import "time"

type Work struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWork(interval time.Duration, r *Runner) *Work {
	return &Work{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Work) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecuor)
	w := NewWork(3, r)
	go w.startWorker()
}
