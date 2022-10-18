package unmarshaltomap

import (
	"github.com/toddnguyen47/util-go/pkg/jsonwrapper"
)

func UnmarshalToMap(input interface{}, jsonWrapper jsonwrapper.Interface) (map[string]interface{}, error) {

	resultsMap := make(map[string]interface{})
	bytes1, err := jsonWrapper.Marshal(input)
	if err != nil {
		return resultsMap, err
	}

	err = jsonWrapper.Unmarshal(bytes1, &resultsMap)
	if err != nil {
		return resultsMap, err
	}

	return resultsMap, nil
}
