package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"testing"
	"time"
)

type MyWork struct {
	Name      string "The Name of a person"
	BirthYear int    "The Year when a person was born"
	WP        *WorkPool
}

func (workPool *MyWork) DoWork() {
	fmt.Printf("%s : %d\n", workPool.Name, workPool.BirthYear)
	fmt.Printf("*******> QW: %d  AR: %d\n", workPool.WP.QueuedWork(), workPool.WP.ActiveRoutines())
	time.Sleep(100 * time.Millisecond)
	//panic("test")
}

func TestAPool(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	workPool := New(runtime.NumCPU()*3, 10)
	shutdown := false //for testing
	go func() {
		for i := 0; i < 1000; i++ {
			work := &MyWork{
				Name:      "A" + strconv.Itoa(i),
				BirthYear: i,
				WP:        workPool,
			}
			err := workPool.PostWork("name_routine", work)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err)
				time.Sleep(100 * time.Millisecond)
			}
			if shutdown == true {
				return
			}
		}
	}()
	t.Log("Hit any key to exit")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	shutdown = true
	t.Log("Shutting Down\n")
	workPool.Shutdown("name_routine")
}
