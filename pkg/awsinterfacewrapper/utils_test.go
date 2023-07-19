package awsinterfacewrapper

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/toddnguyen47/util-go/pkg/pointerutils"
)

var errForTests = errors.New("errorForTests")

// ############################################################################
// #region SETUP
// ############################################################################

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type UtilsTestSuite struct {
	suite.Suite
	ctxBg context.Context
}

func (s *UtilsTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
}

func (s *UtilsTestSuite) TearDownTest() {
	s.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *UtilsTestSuite) Test_GivenInput_ReturnProperString() {
	// -- ARRANGE --
	// You must provide at least the partitionKey and partitionKeyValue
	keyConditionBuilder := expression.KeyAnd(
		expression.Key("pk").Equal(expression.Value("pkValue")),
		expression.Key("sk").GreaterThanEqual(expression.Value("skValue")),
	)
	conditionBuilder := expression.GreaterThan(
		expression.Name("pvod"),
		expression.Value("pvodValue"),
	)
	builder := expression.NewBuilder().WithKeyCondition(keyConditionBuilder).WithCondition(conditionBuilder)
	expr, err := builder.Build()
	assert.Nil(s.T(), err)
	queryInput := dynamodb.QueryInput{
		TableName:                 pointerutils.PtrString("tableName"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
	}
	// -- ACT --
	str1 := GetStringFromMapStringAttributeValue(queryInput.ExpressionAttributeValues)
	// -- ASSERT --
	assert.Greater(s.T(), len(str1), 0)
}

func (s *UtilsTestSuite) Test_GivenError_ThenReturnEmptyString() {
	// -- ARRANGE --
	// You must provide at least the partitionKey and partitionKeyValue
	keyConditionBuilder := expression.KeyAnd(
		expression.Key("pk").Equal(expression.Value("pkValue")),
		expression.Key("sk").GreaterThanEqual(expression.Value("skValue")),
	)
	conditionBuilder := expression.GreaterThan(
		expression.Name("pvod"),
		expression.Value("pvodValue"),
	)
	builder := expression.NewBuilder().WithKeyCondition(keyConditionBuilder).WithCondition(conditionBuilder)
	expr, err := builder.Build()
	assert.Nil(s.T(), err)
	queryInput := dynamodb.QueryInput{
		TableName:                 pointerutils.PtrString("tableName"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
	}
	_attributevalueUnmarshalMap = func(m map[string]types.AttributeValue, out interface{}) error {
		return errForTests
	}
	// -- ACT --
	str1 := GetStringFromMapStringAttributeValue(queryInput.ExpressionAttributeValues)
	// -- ASSERT --
	assert.Equal(s.T(), len(str1), 0)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *UtilsTestSuite) resetMonkeyPatching() {
	_attributevalueUnmarshalMap = attributevalue.UnmarshalMap
}

// #endregion
