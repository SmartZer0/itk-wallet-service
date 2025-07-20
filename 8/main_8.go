package main

import (
	"fmt"
)

type request struct {
	op   string
	resp chan struct{}
}

type WaitGroup struct {
	opCh chan request
}

func NewWaitGroup() *WaitGroup {
	wg := &WaitGroup{opCh: make(chan request)}
	go wg.control()
	return wg
}

func (wg *WaitGroup) control() {
	count := 0
	var waiters []chan struct{}

	for req := range wg.opCh {
		switch req.op {
		case "add":
			count++

		case "done":
			count--
			if count == 0 {

				for _, w := range waiters {
					close(w)
				}

				return
			}

		case "wait":
			if count == 0 {

				close(req.resp)

				return
			}

			waiters = append(waiters, req.resp)
		}
	}
}

func (wg *WaitGroup) Add() {
	wg.opCh <- request{op: "add"}
}

func (wg *WaitGroup) Done() {
	wg.opCh <- request{op: "done"}
}

func (wg *WaitGroup) Wait() {
	resp := make(chan struct{})
	wg.opCh <- request{op: "wait", resp: resp}
	<-resp
}

func main() {
	wg := NewWaitGroup()

	for i := 1; i <= 3; i++ {
		wg.Add()
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d: работает...\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все горутины завершены")
}
