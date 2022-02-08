package common

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_something(t *testing.T) {
	val := 1
	val++
	// Assert
	if val == 2 {
		t.Error("failed while testing the new account validation")
	}
}

// <- chan receive only notation
// async await like behavior simulation
func Test_AsyncAwaitRoutine(t *testing.T) {
	operation := func() <-chan string {
		c := make(chan string)

		go func() {
			time.Sleep(1 * time.Second)
			c <- "value"
		}()

		fmt.Println("task done and send")
		return c
	}

	fmt.Println("waiting for operation..")
	result := <-operation()
	fmt.Println("blocking and reading value..")

	if result != "value" {
		t.Log("fail value: " + result)
		t.Error("Test_AsyncAwaitRoutine test error")
	}
}

// <- chan receive only notation
// async await like behavior simulation with deferred value
func Test_AsyncAwaitRoutineDefer(t *testing.T) {
	operation := func() <-chan string {
		fmt.Println("task init..")
		c := make(chan string)

		go func() {
			fmt.Println("long operation started..")
			time.Sleep(3 * time.Second)
			c <- "value"
		}()

		fmt.Println("task done and send")
		return c
	}

	fmt.Println("waiting for operation..")
	start := operation()
	fmt.Println("still not blocking..")

	result := <-start
	fmt.Println("operation finished")

	if result != "value" {
		t.Log("fail value: " + result)
		t.Error("Test_AsyncAwaitRoutineDefer test error")
	}
}

// behavior like js Promise.all
func Test_MultipleGourutinesParalel(t *testing.T) {
	operation := func() <-chan string {
		fmt.Println("task init..")
		c := make(chan string)

		go func() {
			fmt.Println("long operation started..")
			time.Sleep(3 * time.Second)
			c <- "value"
		}()

		fmt.Println("task done and send")
		return c
	}

	fmt.Println("waiting for operation..")
	o1, o2, o3 := <-operation(), <-operation(), <-operation()
	fmt.Println("still not blocking..")

	fmt.Println("operation finished")

	result := o1 + o2 + o3

	if result != "valuevaluevalue" {
		t.Log("fail value: " + result)
		t.Error("Test_MultipleGourutinesParalel test error")
	}
}

// channel waiting via for loop to channel close
func Test_WaitingForChannelClose(t *testing.T) {
	c := make(chan string)

	count := func(thing string, c chan string) {
		for i := 1; i <= 5; i++ {
			c <- thing
			time.Sleep(time.Millisecond * 500)
		}
		close(c)
	}

	go count("value", c)

	var result string

	for {
		msg, open := <-c
		if !open {
			break
		}
		result += msg

		fmt.Println(msg)
	}

	if result != "valuevaluevaluevaluevalue" {
		t.Log("fail value: " + result)
		t.Error("Test_WaitingForChannelClose test error")
	}

}

// test buffered channel - important note - buffered channel are NOT blocking sice channel buffer is full, so, it can happen in same routine
func Test_BufferedChannel(t *testing.T) {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	result := <-c + <-c

	if result != 3 {
		t.Log("fail value: " + strconv.Itoa(result))
		t.Error("Test_WaitingForChannelClose test error")
	}
}

// test two routines where second routine slowdown first via block
func Test_SlowdownRoutines(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 1/2 second"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()
	for start := time.Now(); time.Since(start) < (time.Second * 10); {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
	// it takes 10 seconds and 2s routines loop count is same as 500ms count - because of blocking behavior
}

// test same as Test_SlowdownRoutines with select statement
// select lets a goroutine waits for multiple communication operations when is available and not blocking whole scope to wait
func Test_SelectRoutines(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Every 500ms"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Every two seconds"
		}
	}()

	for start := time.Now(); time.Since(start) < (time.Second * 10); {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// test worker pattern
func Test_WorkerPattern(t *testing.T) {
	worker := func(jobId int, jobs <-chan int, results chan<- int) {
		for j := range jobs {
			fmt.Println("worker", jobId, "started  job", j)
			time.Sleep(time.Second)
			fmt.Println("worker", jobId, "finished job", j)
			results <- j * 2
		}
	}

	const numJobs = 5
	//buffered channel
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	collect := make([]int, 0, 3)

	// number of worker (business) goroutines
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	// number of jobs todo
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//buffered channel receive
	for a := 1; a <= numJobs; a++ {
		collect = append(collect, <-results)
	}

	var sum int

	for _, res := range collect {
		sum += res
	}

	if sum != 30 {
		t.Log("fail value: " + strconv.Itoa(sum))
		t.Error("Test_WorkerPattern test error")
	}
}

// test waitgroup - dont care about results
func Test_Waitgroups(t *testing.T) {
	worker := func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i

		go func() {
			// information for wg - one of 5 is done
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
