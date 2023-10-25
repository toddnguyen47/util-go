package sarama_msk_wrapper

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/toddnguyen47/util-go/pkg/testhelpers"
)

// /----------------------------------------------------------\
// #region mockConsumerGroup
// ------------------------------------------------------------

type mockConsumerGroup struct {
	sarama.ConsumerGroup

	mpfConsume         testhelpers.MockPassFail
	consumeWaitForStop bool
	stopChan           chan struct{}
	errorChan          chan error
}

func newMockConsumerGroup() *mockConsumerGroup {
	return &mockConsumerGroup{
		stopChan:           make(chan struct{}, 1),
		errorChan:          make(chan error, 1),
		mpfConsume:         testhelpers.NewMockPassFail(),
		consumeWaitForStop: true,
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
	if m.consumeWaitForStop {
		<-m.stopChan
	}
	return nil
}

// ------------------------------------------------------------
// #endregion mockConsumerGroup
// \----------------------------------------------------------/

// /----------------------------------------------------------\
// #region mockConsumerGroupSession
// ------------------------------------------------------------

type mockConsumerGroupSession struct {
	sarama.ConsumerGroupSession
	ctx            context.Context
	cancel         context.CancelFunc
	mpfMarkMessage testhelpers.MockPassFail
}

func newMockConsumerGroupSession() *mockConsumerGroupSession {
	ctx, cancel := context.WithCancel(context.Background())
	m := mockConsumerGroupSession{
		ctx:            ctx,
		cancel:         cancel,
		mpfMarkMessage: testhelpers.NewMockPassFail(),
	}
	return &m
}

func (m *mockConsumerGroupSession) MemberID() string { return "memberId" }

func (m *mockConsumerGroupSession) GenerationID() int32 { return 0 }

func (m *mockConsumerGroupSession) MarkMessage(_ *sarama.ConsumerMessage, _ string) {
	_ = m.mpfMarkMessage.WillPassIncrementCount()
}

func (m *mockConsumerGroupSession) Context() context.Context {
	return m.ctx
}

// ------------------------------------------------------------
// #endregion mockConsumerGroupSession
// \----------------------------------------------------------/

// /----------------------------------------------------------\
// #region mockProcessor
// ------------------------------------------------------------

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

type mockBatchProcessor struct {
	ConsumedBatchOfMessagesProcessor

	setKeys    mapset.Set[string]
	mpfBatch   testhelpers.MockPassFail
	mpfProcess testhelpers.MockPassFail
}

func newMockBatchProcessor() *mockBatchProcessor {
	m := mockBatchProcessor{
		setKeys:    mapset.NewSet[string](),
		mpfBatch:   testhelpers.NewMockPassFail(),
		mpfProcess: testhelpers.NewMockPassFail(),
	}
	return &m
}

func (m *mockBatchProcessor) ProcessConsumedBatchOfMessages(
	consumedMessages []sarama.ConsumerMessage) ([]sarama.ConsumerMessage, error) {

	outerErr := m.mpfBatch.WillPassIncrementCount()
	if outerErr != nil {
		return nil, outerErr
	}

	successes := make([]sarama.ConsumerMessage, 0)
	totalErrorMessages := make([]string, 0)
	for _, msg := range consumedMessages {
		msgKey := string(msg.Key)
		if m.setKeys.Contains(msgKey) {
			panic("should not have duplicated message key!")
		}
		m.setKeys.Add(msgKey)
		err := m.mpfProcess.WillPassIncrementCount()
		if err != nil {
			totalErrorMessages = append(totalErrorMessages, err.Error())
			continue
		}
		successes = append(successes, msg)
	}
	if len(totalErrorMessages) > 0 {
		err := fmt.Errorf("total errors: %d; errors: %s", len(totalErrorMessages),
			strings.Join(totalErrorMessages, ","))
		return successes, err
	}
	return successes, nil
}

// ------------------------------------------------------------
// #endregion mockProcessor
// \----------------------------------------------------------/

// /----------------------------------------------------------\
// #region mockConsumerGroupClaimStruct
// ------------------------------------------------------------

type mockConsumerGroupClaimStruct struct {
	sarama.ConsumerGroupClaim

	chanConsumerMessage chan *sarama.ConsumerMessage
}

func newMockConsumerGroupClaimStruct() *mockConsumerGroupClaimStruct {
	m := mockConsumerGroupClaimStruct{
		chanConsumerMessage: make(chan *sarama.ConsumerMessage),
	}
	return &m
}

func (m *mockConsumerGroupClaimStruct) Topic() string {
	return "topic"
}

func (m *mockConsumerGroupClaimStruct) Partition() int32 {
	return 1
}

func (m *mockConsumerGroupClaimStruct) InitialOffset() int64 {
	return 200
}

func (m *mockConsumerGroupClaimStruct) Messages() <-chan *sarama.ConsumerMessage {
	return m.chanConsumerMessage
}

// ------------------------------------------------------------
// #endregion mockConsumerGroupClaimStruct
// \----------------------------------------------------------/
