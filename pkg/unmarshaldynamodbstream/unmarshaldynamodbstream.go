package unmarshaldynamodbstream

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Monkey patching for tests
var (
	_jsonUnmarshal = json.Unmarshal
)

// Unmarshal converts events.DynamoDBAttributeValue to struct
// Reference: https://stackoverflow.com/a/50017398/6323360
func Unmarshal(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {

	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range attribute {

		var dbAttr dynamodb.AttributeValue

		bytes, err := v.MarshalJSON()
		if err != nil {
			return err
		}

		err = _jsonUnmarshal(bytes, &dbAttr)
		if err != nil {
			return err
		}
		dbAttrMap[k] = &dbAttr
	}

	return dynamodbattribute.UnmarshalMap(dbAttrMap, out)
}
