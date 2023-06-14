package chantrial

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const _maxNumMessages = 100

func Test_Chan(t *testing.T) {
	c1 := make(chan string)
	doneChannel := make(chan int)
	totalCount := uint64(0)
	go subscriber(c1, doneChannel, &totalCount)
	go publisher(c1)

	<-doneChannel
	fmt.Printf("Total count: %d\n", totalCount)
	assert.Equal(t, _maxNumMessages, int(totalCount))
}

func publisher(chanString chan string) {

	reader := rand.Reader
	min := int64(10)
	max := big.NewInt(100 - min + 1)

	for i := 0; i < _maxNumMessages; i++ {
		if i > 0 {
			bigInt, _ := rand.Int(reader, max)
			bigInt2 := bigInt.Int64() + min
			sleepTime := time.Duration(bigInt2)
			fmt.Printf("Sleeping for: %s\n", sleepTime.String())
			time.Sleep(sleepTime * time.Millisecond)
		}
		msg := fmt.Sprintf("Number %d", i)
		//fmt.Println(msg)
		chanString <- msg
	}

	// Close channel as we are the sender, otherwise channel will keep being opened
	close(chanString)
}

func subscriber(chanString <-chan string, doneChannel chan int, totalCount *uint64) {
	elems := make([]string, 0)
	start := time.Now().UTC()

	for elem := range chanString {
		elems = append(elems, elem)
		now := time.Now().UTC()
		dur := now.Sub(start)
		lenElems := len(elems)
		if lenElems >= 5 || dur.Milliseconds() > 200 {
			doWork(elems, dur, totalCount)
			// Reset now
			elems = make([]string, 0)
			start = time.Now().UTC()
		}
	}

	// Any remaining elements
	if len(elems) > 0 {
		fmt.Println("Finishing remaining elements")
		now := time.Now().UTC()
		dur := now.Sub(start)
		doWork(elems, dur, totalCount)
	}

	// Signals that the subscriber is done
	doneChannel <- 1
}

func doWork(elems []string, dur time.Duration, totalCount *uint64) {
	lenElems := len(elems)
	fmt.Printf("There are currently %d items in the list. Time duration: %s \n", lenElems, dur.String())
	msg := strings.Join(elems, " @@@ ")
	fmt.Println("msg: " + msg)
	atomic.AddUint64(totalCount, uint64(lenElems))
}
