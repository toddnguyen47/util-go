package buildurl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	championStr  = "champion"
	startDateStr = "startDate"
	endDateStr   = "endDate"
)

// ------------------------------------------------------------
// #region SETUP
// ------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type BuildUrlTestSuite struct {
	suite.Suite
}

func (suite *BuildUrlTestSuite) SetupTest() {
	suite.resetMonkeyPatching()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestBuildUrlTestSuite(t *testing.T) {
	suite.Run(t, new(BuildUrlTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (suite *BuildUrlTestSuite) Test_GivenUrlWithSlash_When_ThenUrlIsBuiltCorrectly() {

	baseUrl := "https://some-url.dev-int/"
	pathUrl := "currentpatch/coldstart"
	queryParams := make(map[string]interface{})
	queryParams[championStr] = "LEBLANC"
	queryParams[startDateStr] = "2020-02-06T13:00:00.000Z"
	queryParams[endDateStr] = "2020-02-06T15:30:00.000Z"

	url := Build(baseUrl, pathUrl, queryParams)

	assert.True(suite.T(),
		strings.Contains(url, "https://some-url.dev-int/currentpatch/coldstart?"),
	)
	assert.True(suite.T(), strings.Contains(url, "champion=LEBLANC"))
	assert.True(suite.T(), strings.Contains(url, startDateStr))
	assert.True(suite.T(), strings.Contains(url, endDateStr))
	assert.Equal(suite.T(), 2, strings.Count(url, "&"))
}

func (suite *BuildUrlTestSuite) Test_GivenUrlWithoutSlash_When_ThenUrlIsBuiltCorrectly() {

	baseUrl := "https://some-url.dev-int/currentpatch"
	pathUrl := "coldstart"
	queryParams := make(map[string]interface{})
	queryParams[championStr] = "LEBLANC"
	queryParams[startDateStr] = "2020-02-06T13:00:00.000Z"
	queryParams[endDateStr] = "2020-02-06T15:30:00.000Z"
	queryParams["retryCount"] = 2

	url := Build(baseUrl, pathUrl, queryParams)

	assert.True(suite.T(),
		strings.Contains(url, "https://some-url.dev-int/currentpatch/coldstart?"),
	)
	assert.True(suite.T(), strings.Contains(url, "champion=LEBLANC"))
	assert.True(suite.T(), strings.Contains(url, startDateStr))
	assert.True(suite.T(), strings.Contains(url, endDateStr))
	assert.True(suite.T(), strings.Contains(url, "retryCount=2"))
	assert.Equal(suite.T(), len(queryParams)-1, strings.Count(url, "&"))
}

func (suite *BuildUrlTestSuite) Test_GivenUrlNoQueryParams_When_ThenUrlIsBuiltCorrectly() {

	baseUrl := "https://some-url.dev-int/currentpatch"
	pathUrl := "coldstart"

	url := Build(baseUrl, pathUrl, nil)

	assert.True(suite.T(),
		strings.Contains(url, "https://some-url.dev-int/currentpatch/coldstart"),
	)
	assert.Equal(suite.T(), 0, strings.Count(url, "&"))
}

func (suite *BuildUrlTestSuite) Test_GivenUrlParseErr_When_ThenReturnNil() {
	url := Build("postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require", "", nil)

	assert.Equal(suite.T(), "", url)
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (suite *BuildUrlTestSuite) resetMonkeyPatching() {
}

// #endregion
