package sarama_msk_wrapper

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/toddnguyen47/util-go/pkg/testhelpers"
)

// /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
// #region mockConsumerGroup
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

type mockConsumerGroup struct {
	sarama.ConsumerGroup

	mpfConsume testhelpers.MockPassFail
	stopChan   chan struct{}
	errorChan  chan error
}

func newMockConsumerGroup() *mockConsumerGroup {
	return &mockConsumerGroup{
		stopChan:   make(chan struct{}, 1),
		errorChan:  make(chan error, 1),
		mpfConsume: testhelpers.NewMockPassFail(),
	}
}

func (m *mockConsumerGroup) stop() {
	m.stopChan <- struct{}{}
}

func (m *mockConsumerGroup) Close() error {
	return nil
}

func (m *mockConsumerGroup) Errors() <-chan error {
	return m.errorChan
}

func (m *mockConsumerGroup) Consume(_ context.Context, _ []string,
	_ sarama.ConsumerGroupHandler) error {

	err := m.mpfConsume.WillPassIncrementCount()
	if err != nil {
		return err
	}
	// Wait before returning
	<-m.stopChan
	return nil
}

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// #endregion mockConsumerGroup
// \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/

// /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
// #region mockConsumerGroupSession
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

type mockConsumerGroupSession struct {
	sarama.ConsumerGroupSession
	ctx context.Context
}

func newMockConsumerGroupSession() *mockConsumerGroupSession {
	m := mockConsumerGroupSession{
		ctx: context.Background(),
	}
	return &m
}

func (m *mockConsumerGroupSession) MemberID() string { return "memberId" }

func (m *mockConsumerGroupSession) GenerationID() int32 { return 0 }

func (m *mockConsumerGroupSession) MarkMessage(_ *sarama.ConsumerMessage, _ string) {}

func (m *mockConsumerGroupSession) Context() context.Context {
	return m.ctx
}

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// #endregion mockConsumerGroupSession
// \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/

// /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\
// #region mockProcessor
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

type mockProcessor struct {
	ConsumedMessageProcessor

	mpfProcess testhelpers.MockPassFail
}

func newMockProcessor() *mockProcessor {
	m := mockProcessor{
		mpfProcess: testhelpers.NewMockPassFail(),
	}
	return &m
}

func (m *mockProcessor) ProcessConsumedMessage(_ *sarama.ConsumerMessage) error {
	err := m.mpfProcess.WillPassIncrementCount()
	if err != nil {
		return err
	}
	return nil
}

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// #endregion mockProcessor
// \@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/
