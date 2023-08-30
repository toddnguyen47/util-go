package chantrial

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/toddnguyen47/util-go/pkg/jsonutils"
)

const _maxNumMessages = 100

func Test_Chan(t *testing.T) {
	// Reference: https://go.dev/blog/pipelines
	c1 := make(chan string)
	publisherDone := make(chan struct{})
	subscriberDone := make(chan struct{})
	totalCount := uint64(0)
	go subscriber(c1, publisherDone, subscriberDone, &totalCount)
	go publisher(c1, publisherDone)

	<-subscriberDone
	fmt.Printf("Total count: %d\n", totalCount)
	assert.Equal(t, _maxNumMessages, int(totalCount))
}

func publisher(chanString chan<- string, publisherDone chan<- struct{}) {

	reader := rand.Reader
	min := int64(10)
	max := big.NewInt(100 - min + 1)

	for i := 0; i < _maxNumMessages; i++ {
		if i > 0 {
			bigInt, _ := rand.Int(reader, max)
			bigInt2 := bigInt.Int64() + min
			sleepTime := time.Duration(bigInt2) * time.Millisecond
			fmt.Printf("Sleeping for: %s\n", sleepTime.String())
			time.Sleep(sleepTime)
		}
		msg := fmt.Sprintf("Number %d", i)
		//fmt.Println(msg)
		chanString <- msg
	}
	// We are done now; send signal
	publisherDone <- struct{}{}
}

func subscriber(chanString <-chan string, publisherDone <-chan struct{}, subscriberDone chan<- struct{},
	totalCount *uint64) {

	elems := make([]string, 0)

	for {
		select {
		case elem := <-chanString:
			// We received an element
			elems = append(elems, elem)
			lenElems := len(elems)
			if lenElems >= 5 {
				doWork(elems, totalCount)
				// Reset now
				elems = make([]string, 0)
			}
		case <-time.After(75 * time.Millisecond):
			// Timed out
			fmt.Println("timed out")
			doWork(elems, totalCount)
			// Reset now
			elems = make([]string, 0)
		case <-publisherDone:
			// Do any work on the remaining elements
			fmt.Println("last stretch!")
			doWork(elems, totalCount)
			// Reset now
			elems = make([]string, 0)
			// Notify that subscriber is finished
			subscriberDone <- struct{}{}
		}
	}
}

func doWork(elems []string, totalCount *uint64) {
	lenElems := len(elems)
	fmt.Printf("There are currently %d items in the list. \n", lenElems)
	msg := strings.Join(elems, " @@@ ")
	fmt.Println("msg: " + msg)
	atomic.AddUint64(totalCount, uint64(lenElems))
}

type struct1 struct {
	Id string `json:"id"`
}

type struct2 struct {
	Name string `json:"name"`
}

type combinedStruct struct {
	Struct1 struct1 `json:"struct1"`
	Struct2 struct2 `json:"struct2"`
}

func Test_Wg(t *testing.T) {

	const numberOfCalls = 2
	var wg sync.WaitGroup
	wg.Add(numberOfCalls)
	var mutex sync.Mutex
	combined := combinedStruct{}

	go func() {
		defer wg.Done()
		sleepTime := 500 * time.Millisecond
		fmt.Printf("sleeping for %s\n", sleepTime.String())
		time.Sleep(sleepTime)
		s1 := struct1{Id: "id1"}
		mutex.Lock()
		defer mutex.Unlock()
		combined.Struct1 = s1
	}()

	go func() {
		defer wg.Done()
		sleepTime := 200 * time.Millisecond
		fmt.Printf("sleeping for %s\n", sleepTime.String())
		time.Sleep(sleepTime)
		s2 := struct2{Name: "name"}
		mutex.Lock()
		defer mutex.Unlock()
		combined.Struct2 = s2
	}()

	wg.Wait()
	b1, err := jsonutils.MarshalNoEscapeHtml(combined)
	assert.Nil(t, err)
	fmt.Println(string(b1))
	// outputs: {"struct1":{"id":"id1"},"struct2":{"name":"name"}}
}
