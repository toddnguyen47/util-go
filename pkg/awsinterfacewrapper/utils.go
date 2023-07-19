package awsinterfacewrapper

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Monkey patching for tests
var _attributevalueUnmarshalMap = attributevalue.UnmarshalMap

func GetStringFromMapStringAttributeValue(input map[string]types.AttributeValue) string {
	m1 := make(map[string]interface{})
	const emptyStr = ""
	err := _attributevalueUnmarshalMap(input, &m1)
	if err != nil {
		return emptyStr
	}
	b1, _ := json.Marshal(m1)
	s1 := string(b1)
	return s1
}
