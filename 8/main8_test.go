package main

type wgRequest struct {
	op   string
	resp chan struct{}
}

type CustomWaitGroup struct {
	opCh chan wgRequest
}

func NewCustomWaitGroup() *CustomWaitGroup {
	wg := &CustomWaitGroup{opCh: make(chan wgRequest)}
	go wg.controlLoop()
	return wg
}

func (wg *CustomWaitGroup) controlLoop() {
	count := 0
	var waiters []chan struct{}

	for req := range wg.opCh {
		switch req.op {
		case "add":
			count++
		case "done":
			count--
		case "wait":
			if count == 0 {
				close(req.resp)
				return
			}
			waiters = append(waiters, req.resp)
		}

		if count == 0 && len(waiters) > 0 {
			for _, w := range waiters {
				close(w)
			}
			return
		}
	}
}

func (wg *CustomWaitGroup) Add() {
	wg.opCh <- wgRequest{op: "add"}
}

func (wg *CustomWaitGroup) Done() {
	wg.opCh <- wgRequest{op: "done"}
}

func (wg *CustomWaitGroup) Wait() {
	resp := make(chan struct{})
	wg.opCh <- wgRequest{op: "wait", resp: resp}
	<-resp
}
