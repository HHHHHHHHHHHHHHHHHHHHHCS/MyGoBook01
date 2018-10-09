package OtherLib

import (
	"log"
	"os"
	"testing"
	"time"
)

const timeout = 3 * time.Second

func TestOps(t *testing.T) {
	log.Println("Starting work")

	r := New(timeout)

	r.Add(CreateTask(), CreateTask(), CreateTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println("Terminating die to timeout.")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating die to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func CreateTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
