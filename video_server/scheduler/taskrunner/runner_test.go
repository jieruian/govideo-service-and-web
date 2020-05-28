package taskrunner

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestRuner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %v", i)
		}
		fmt.Println("走不走啊")
		return nil
	}

	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executor received: %v", d)
			default:
				break forloop
			}
		}

		return errors.New("Executor")
	}

	runner := NewRunner(60, false, d, e)
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}
