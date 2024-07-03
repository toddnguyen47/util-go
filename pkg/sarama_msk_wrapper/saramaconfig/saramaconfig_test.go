package saramaconfig

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/toddnguyen47/util-go/pkg/osutils"
	"github.com/toddnguyen47/util-go/pkg/sarama_msk_wrapper/saramainject"
)

var errForTests = errors.New("errForTests")

// ------------------------------------------------------------
// #region SETUP
// ------------------------------------------------------------

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type SaramaConfigTestSuite struct {
	suite.Suite
	ctxBg     context.Context
	certPool  *x509.CertPool
	principal string
}

func (s *SaramaConfigTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.certPool = x509.NewCertPool()
	_getCertsFrom = func(certsLocation string) (*x509.CertPool, error) {
		return s.certPool, nil
	}
	s.principal = "username@realm"
}

func (s *SaramaConfigTestSuite) TearDownTest() {
	s.resetMonkeyPatching()

	// Remove tmpCertFolder
	tmpCertFolder := saramainject.TmpCertFolder(s.principal)
	files, err := filepath.Glob(tmpCertFolder + "/*")
	assert.Nil(s.T(), err)
	for _, file := range files {
		_ = osutils.RemoveIfExists(file)
	}
	_ = osutils.RemoveIfExists(tmpCertFolder)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSaramaConfigTestSuite(t *testing.T) {
	suite.Run(t, new(SaramaConfigTestSuite))
}

// #endregion

// ------------------------------------------------------------
// #region TESTS ARE BELOW
// ------------------------------------------------------------

func (s *SaramaConfigTestSuite) Test_GivenProperSSLCerts_ThenGetConfigProperly() {
	// -- GIVEN --
	pubKey, privateKey := s.getCerts()
	// -- WHEN --
	config := GetSaramaConfigSsl(pubKey, privateKey)
	// -- THEN --
	assert.NotNil(s.T(), config)
}

func (s *SaramaConfigTestSuite) Test_GivenTlsError_ThenPanics() {
	// -- GIVEN --
	// -- WHEN --
	// -- THEN --
	assert.Panics(s.T(), func() {
		GetSaramaConfigSsl([]byte("hello"), []byte("world"))
	})
}

func (s *SaramaConfigTestSuite) Test_GivenProperSASLSSLCerts_ThenGetConfigProperly() {
	// -- GIVEN --
	// -- WHEN --
	config := GetSaramaConfigSasl(s.principal, []byte("kerbKeyTab"), []byte("kerbConf"), []byte("sslCert"))
	// -- THEN --
	assert.NotNil(s.T(), config)
}

func (s *SaramaConfigTestSuite) Test_GivenNonProperSASLSSLCerts_ThenPanic() {
	// -- GIVEN --
	_getCertsFrom = func(certsLocation string) (*x509.CertPool, error) {
		return nil, errForTests
	}
	// -- WHEN --
	// -- THEN --
	assert.Panics(s.T(), func() {
		GetSaramaConfigSasl(s.principal, []byte("kerbKeyTab"), []byte("kerbConf"), []byte("sslCert"))
	})
}

// ------------------------------------------------------------
// #region TEST HELPERS
// ------------------------------------------------------------

func (s *SaramaConfigTestSuite) resetMonkeyPatching() {
}

// #endregion

func (s *SaramaConfigTestSuite) getCerts() ([]byte, []byte) {
	publicBase64 := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURCRENDQWV3Q0NRRHFvVm9aeERwU3lqQU5CZ2txaGtpRzl3MEJBUXNGQURCRU1Rc3dDUVlEVlFRR0V3SjEKY3pFTE1Ba0dBMVVFQ0F3Q1kyRXhGREFTQmdOVkJBY01DMnh2Y3lCaGJtZGxiR1Z6TVJJd0VBWURWUVFLREFsdAplV052YlhCaGJua3dIaGNOTWpNd05UQXhNVGMwTURReldoY05Nall3TVRJMU1UYzBNRFF6V2pCRU1Rc3dDUVlEClZRUUdFd0oxY3pFTE1Ba0dBMVVFQ0F3Q1kyRXhGREFTQmdOVkJBY01DMnh2Y3lCaGJtZGxiR1Z6TVJJd0VBWUQKVlFRS0RBbHRlV052YlhCaGJua3dnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDaQpKVVd3a2xzR2RETXJkZUNlSFNjbUNtS0x4Y2JJek94bWc2RjEwQmVENlVNdDRRZHFhUHpZM09ucVR0TWdkUjBICkgra0k1Qllzbi8zcXJyblJ4VWxYbExyQ0tjZlhScGNUZlhsTFk4QS90OEQ3K1pvOWNXYjhnR0toelJWWVg1YzQKWW5VZmJiMUxVU3VuVFJMaDN6VzVuWGl1RzlUMi90UDI1UW9vWjV4QUtLejVDcDJoUGg4cklWUEFiU1FzamtEcwpqckVsWEYxL080Wkc4TTh6bmxQMEJLSklFL2lyL2FKRnNjM2NQaGlEK0F3QXM5TFp1WXVpa0xjcGJDa0hqV0FVCk1ZbDdMUkp0OE1ZK1FqRlpJZmJKalNNZS9HQktKd1IzM3V5ZTJLVmhSU0g0K2xZRUQ1R0RtSU5teEdlOTZ2YzAKWUdXcXRNcktHb1AzZzVGRForQVJBZ01CQUFFd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFHL1BtbVZJL3BNZApqL0lxQXlWZ2JBTk1OdTFjUldtd0xhQWVmWWp4OE9UY2RlSG5aOWdyaVgyQVk5b2Q2anZBdW5IdXZhalVHVGJ2Ci9yYUQzVWxSYUE2bVZiQkorZ3JhNW1nbVlHakpWbDdmSzBmMjFKajIyV01leUQrbFQxSkNnKzVxQjhsQXFibDgKV0ttVVBLZXRuZitmaU9zYjNZVW5ySE85OGp6ODNmdWdXNHBsOVkzV2RhQ2ZiVW81a0kvMC8rMzFWWE4vaERvOQpLVzFyWkxZdWZhOGtKZnVYYjhiYitrUFo3MEc1bmI5L3F3WlphL01zWktvYzF1Z2RJSDdpbXZZczMyUCtZeEFQCnVJTFNoNWdxMVE1Mm5TdW9qZGhIR1V3K3M0VFBpY253MEtnRHpRN2hpclhpeG5rZkUwNWlJbGhtbk03QlNGYTEKWGMwZldsRHA4bVE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
	privateBase64 := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBb2lWRnNKSmJCblF6SzNYZ25oMG5KZ3BpaThYR3lNenNab09oZGRBWGcrbERMZUVICmFtajgyTnpwNms3VElIVWRCeC9wQ09RV0xKLzk2cTY1MGNWSlY1UzZ3aW5IMTBhWEUzMTVTMlBBUDdmQSsvbWEKUFhGbS9JQmlvYzBWV0YrWE9HSjFIMjI5UzFFcnAwMFM0ZDgxdVoxNHJodlU5djdUOXVVS0tHZWNRQ2lzK1FxZApvVDRmS3lGVHdHMGtMSTVBN0k2eEpWeGRmenVHUnZEUE01NVQ5QVNpU0JQNHEvMmlSYkhOM0Q0WWcvZ01BTFBTCjJibUxvcEMzS1d3cEI0MWdGREdKZXkwU2JmREdQa0l4V1NIMnlZMGpIdnhnU2ljRWQ5N3NudGlsWVVVaCtQcFcKQkErUmc1aURac1JudmVyM05HQmxxclRLeWhxRDk0T1JRMmZnRVFJREFRQUJBb0lCQVFDTGhIUzFuUWt6d1hKeQpGK2loUkVaMlJnTkxiMjgvUW95N0hOSW1ORnEzaDFQbHV6WU5TcThkenVCN1d6M3hOTFE0ZUMybEY5VTRxcVhxCmRGT3hld2REazljcTBKYUMxdHVSeXFvK3cyTWRzSGdlbUVRdVVGQ2tQYmdncnYySjRCNlhScWl1MVZkRzRsNGsKZzM0VzJtQTVDWlZZZ3R3NWQwVmRzUENQbWE5cHgvZEM0MjBwUmtqTEplV0tCYlk2amJ0Vzc3azV5Zi9oZU1pUAozeUF6aHc5bHN1MEtnaEFFRzhrNkZhL2piRmozcVFyN2dRaXdLWVdMak9NaHdFUTNaMThnY04zNjZPOXBhWEYyCjI4OVRSeEFZcHQxbUdFWkRFeDF4b284T0JJVUpvZlA5Uk4zeHZZU2pHYmQzTEpnY3czZG9DUVBVeUhLTjluQXcKeERnS2xVbjlBb0dCQU5Qbk5Fd3A1aG1MTEticjc2R0g5SEhteXBVZEJ1UFlBZHJqRCt5Wjd5eitsL0RWb2R2NgpXTHhKMlFCK0U1T0w4SThuN2pvQ1Vma0NUMUlEVFl1aisyeEREK3hXb3JLdHljenM4VWlhTElpK2lSNVdLdjhWCmF4aFhDeFJoQk5tYzlRKzI4cHVWTWdJKzk1MkNhd3VHbHJ6dUdOUXFjOGMxLytkQmZjdTdwQmp2QW9HQkFNUGoKVU1BQlprdHBneDE2czh1WVluT2hVZFdyUHVROEluQ0ZQZ3R1cDhvQkZlOStmNDV2cFh1RG1DWEJnek9IMUdtcwpTWFkwK05kSTJadjFid2IzYUluZndpTUhqZ0pHcEhkY3BvZytsUmwvRWM5bTVKQWdoOUp5dHhZL1FUWmF3VWFaCm0xZFluQThGdThRZ1B2cUJnRVYyY3VERktVbXQ0RVJPMHRHTEdaYi9Bb0dBT1FobWZBVmU2QXNjWm9Ua0J6N1gKWFB5NEU3QXZWWTJpMmkzNDhENXlNRk1KeEFsTHVqQkVSOUU4ZGJSNVFtSU15Z3IrUkdDeGZXclF2SXNsQ091RwoxUm1ycEhtZzZxUjV4dzBTMSs2ZkErTDhkc3pNWDhGOUJKMFEzMWhKZk9TUTFMenh5VXc3bkgwa0dpR3ErL3dxCmdBazVaNGxSaGhHVG9jTnZ2ekR1dHNNQ2dZQi9mYzBxaWo2bnlrNVp1MmlWaytKUDI2akZaaVVTcXNqSGJ4RUkKbzhaMHhPd2Y2YmJmWDI3V3lya1ZxYkxZc1FqZ2xnOWg3ZXdmUWZ6UGNwZ0djclFKT0NiRVljQmRYdGpnRHQ4YwpRWThNL3hUNlpiOVF4cnRmanVYMmhzak10WmloZUl3UDkwM3F3UktKL1dxLzQ5VTJZSGM0TDFwRjUvTFV3bkNYCmpPN2t4UUtCZ1FDTzlubEUyY3czRjUyaHVwc1ZSY0tKWXIwOWhVQ1ZEcWV1SlJrSXpnZWN5TnNnaTVDc3E4clcKRmpCOHVNUWxjUHpsbHZ4NndlZVJJZW5xcW9hUThVYWk0cjZKMjdUNEI4QWoweG1zRWZVbWdGcy9id0FlRlYrdwo2b0tacEpnOHYvRGNET3pUREMrUUZrVTc1emt1WjJUVTJPMmR5T013dWcyR0lxeHljVThBcmc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
	pubKey, err := base64.StdEncoding.DecodeString(publicBase64)
	assert.Nil(s.T(), err)
	privateKey, err := base64.StdEncoding.DecodeString(privateBase64)
	assert.Nil(s.T(), err)
	return pubKey, privateKey
}
