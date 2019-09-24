/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-11-09 11:06:20
 */

package goutils

import (
	"time"
)

// Future is an interface to get results from other goroutines.
type Future interface {
	// Get the data until the task is done.
	Get() interface{}
	// Get the data until the task is done or the timeout is reached.
	GetTimeout(timeout time.Duration) interface{}
	// Wheather the task is done or not.
	IsDone() bool
	// Ouput data used by producer
	Output(interface{})
}

type FutureTask struct {
	ch   chan interface{}
	done bool
}

// NewFutureTask will create a task which implements Future interface.
func NewFutureTask() *FutureTask {
	return &FutureTask{make(chan interface{}, 1), false}
}

func (f *FutureTask) Get() interface{} {
	v := <-f.ch
	f.done = true

	return v
}

func (f *FutureTask) GetTimeout(timeout time.Duration) interface{} {
	select {
	case v := <-f.ch:
		f.done = true
		return v
	case <-time.After(timeout):
		return nil
	}
}

func (f *FutureTask) IsDone() bool {
	return f.done
}

// Output will be called when the task is done.
func (f *FutureTask) Output(v interface{}) {
	f.ch <- v
}
