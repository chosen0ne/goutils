/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2016-11-14 17:08:01
 */

package goutils

type Semaphore struct {
	semaphores chan interface{}
}

// NewSemaphore will create a Semaphore which support P, V operations.
// 'capacity' is the max number of goroutines that can enter critical section.
func NewSemaphore(capacity int) *Semaphore {
	sema := &Semaphore{}
	sema.semaphores = make(chan interface{}, capacity)

	return sema
}

func (s *Semaphore) P() {
	s.semaphores <- 1
}

func (s *Semaphore) V() {
	<-s.semaphores
}
