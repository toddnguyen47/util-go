package sarama_msk_wrapper

import (
	"sync"
	"sync/atomic"

	"github.com/IBM/sarama"
)

// ------------------------------------------------------------
// #region mockAsyncProducer

type mockAsyncProducer struct {
	sarama.AsyncProducer

	mutex sync.Mutex

	mockStopChan chan struct{}
	inputChan    chan *sarama.ProducerMessage
	inputCount   atomic.Uint32
	// inputErrorCode - "F" for fail, "P" for pass
	inputErrorCode string

	errorsChan chan *sarama.ProducerError
	errorCount atomic.Uint32

	closeCount atomic.Uint32
	closeCode  string

	successChan  chan *sarama.ProducerMessage
	successCount atomic.Uint32
}

func newMockAsyncProducer() *mockAsyncProducer {

	m1 := &mockAsyncProducer{
		errorsChan:   make(chan *sarama.ProducerError),
		inputChan:    make(chan *sarama.ProducerMessage),
		mockStopChan: make(chan struct{}, 1),
		successChan:  make(chan *sarama.ProducerMessage),
	}
	m1.start()
	return m1
}

func (m *mockAsyncProducer) Errors() <-chan *sarama.ProducerError {
	return m.errorsChan
}

func (m *mockAsyncProducer) Input() chan<- *sarama.ProducerMessage {
	success := true
	m.mutex.Lock()
	if m.inputErrorCode != "" {
		firstChar := m.inputErrorCode[0]
		m.inputErrorCode = m.inputErrorCode[1:]
		if firstChar == 'F' {
			m.errorsChan <- &sarama.ProducerError{
				Msg: &sarama.ProducerMessage{},
				Err: errForTests,
			}
			m.errorCount.Add(1)
			success = false
		}
	}
	m.mutex.Unlock()

	if success {
		m.successChan <- &sarama.ProducerMessage{}
	}

	return m.inputChan
}

func (m *mockAsyncProducer) Close() error {
	m.closeCount.Add(1)
	if m.closeCode != "" {
		firstChar := m.closeCode[0]
		m.closeCode = m.closeCode[1:]
		if firstChar == 'F' {
			return errForTests
		}
	}
	return nil
}

func (m *mockAsyncProducer) Successes() <-chan *sarama.ProducerMessage {
	m.successCount.Add(1)
	return m.successChan
}

func (m *mockAsyncProducer) start() {
	go func() {
	MockLoop:
		for {
			select {
			case <-m.inputChan:
				m.inputCount.Add(1)
			case <-m.mockStopChan:
				break MockLoop
			}
		}
	}()
}

func (m *mockAsyncProducer) stop() {
	m.mockStopChan <- struct{}{}
}

// #endregion mockAsyncProducer
// o----------------------------------------------------------o

func getIntFromAtomic(num *atomic.Uint32) int {
	if num == nil {
		return 0
	}
	return int(num.Load())
}
