package governor

import (
	"sync"
)

// Governor is a concurrency controlling tool.
// It is passed into goroutines in order to limit how many can run at the same
// time.
// It also pauses control flow until all of the goroutines have finished.
type Governor struct {
	wg        *sync.WaitGroup
	throttle  int
	semaphore chan bool
	errs      chan error
}

// NewGovernor creates and initializes a new Governor.
// It will control `size` total goroutines
// by allowing only `throttle` goroutines to run at the same time.
func NewGovernor(size int, throttle int) Governor {

	var wg sync.WaitGroup
	semaphore := make(chan bool, throttle)
	errs := make(chan error, size)

	g := Governor{
		wg:        &wg,
		throttle:  throttle,
		semaphore: semaphore,
		errs:      errs,
	}

	return g
}

// Accelerate tells the Governor to control another goroutine.
func (g *Governor) Accelerate() {

	g.wg.Add(1)
	g.semaphore <- true
}

// Decelerate tells the Governor that a goroutine has finished.
func (g *Governor) Decelerate(err error) {

	g.errs <- err
	<-g.semaphore
	g.wg.Done()
}

// Regulate tells the Governor to start watching goroutines and to stop further
// command execution until all of them have finished.
// It returns the first error encountered, if any.
func (g *Governor) Regulate() error {

	g.spin()
	g.coast()
	g.stop()

	return g.condition()
}

func (g *Governor) spin() {

	for govolution := 0; govolution < g.throttle; govolution++ {
		g.semaphore <- true
	}
}

func (g *Governor) coast() {

	g.wg.Wait()
}

func (g *Governor) stop() {

	close(g.errs)
}

func (g *Governor) condition() error {

	for err := range g.errs {
		if err != nil {
			return err
		}
	}
	return nil
}