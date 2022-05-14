package unmarshaldynamodbstream

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ############################################################################
// #region SETUP
// ############################################################################

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type UnmarshalDynamoDbStreamSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (suite *UnmarshalDynamoDbStreamSuite) SetupTest() {
	suite.resetMonkeyPatching()
	suite.ctxBg = context.Background()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUnmarshalDynamoDbStreamSuite(t *testing.T) {
	suite.Run(t, new(UnmarshalDynamoDbStreamSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (suite *UnmarshalDynamoDbStreamSuite) Test_GivenJsonUnmarshalErr_When_ThenReturnError() {
	testContent := suite.readTestFile("testdata/dynamodb1.json")
	var record map[string]events.DynamoDBAttributeValue
	_ = json.Unmarshal(testContent, &record)
	_jsonUnmarshal = func(data []byte, v interface{}) error {
		return errors.New("some error")
	}
	var out interface{}

	err := Unmarshal(record, &out)

	assert.NotNil(suite.T(), err)
}

func (suite *UnmarshalDynamoDbStreamSuite) Test_GivenValidEverything_When_ThenErrIsNil() {
	testContent := suite.readTestFile("testdata/dynamodb1.json")
	var record map[string]events.DynamoDBAttributeValue
	_ = json.Unmarshal(testContent, &record)
	var out interface{}

	err := Unmarshal(record, &out)

	assert.Nil(suite.T(), err)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (suite *UnmarshalDynamoDbStreamSuite) resetMonkeyPatching() {
	_jsonUnmarshal = json.Unmarshal
}

func (suite *UnmarshalDynamoDbStreamSuite) readTestFile(fileName string) []byte {
	content, err := os.ReadFile(fileName)
	if err != nil {
		assert.Fail(suite.T(), "Cannot open file: "+fileName)
	}
	return content
}

// #endregion
