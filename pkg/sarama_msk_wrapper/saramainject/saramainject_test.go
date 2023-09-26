package saramainject

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/toddnguyen47/util-go/pkg/osutils"
	"github.com/toddnguyen47/util-go/pkg/testhelpers"
)

var errForTests = errors.New("errForTests")
var mpfWriteDir testhelpers.MockPassFail

// ############################################################################
// #region SETUP
// ############################################################################

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type SaramaInjectTestSuite struct {
	suite.Suite
	ctxBg     context.Context
	principal string
}

func (s *SaramaInjectTestSuite) SetupTest() {
	s.resetMonkeyPatching()
	s.ctxBg = context.Background()
	s.principal = "username@realm"
}

func (s *SaramaInjectTestSuite) TearDownTest() {
	s.resetMonkeyPatching()

	// Remove tmpCertFolder
	tmpCertFolder := TmpCertFolder(s.principal)
	files, err := filepath.Glob(tmpCertFolder + "/*")
	assert.Nil(s.T(), err)
	for _, file := range files {
		_ = osutils.RemoveIfExists(file)
	}
	_ = osutils.RemoveIfExists(tmpCertFolder)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSaramaInjectTestSuite(t *testing.T) {
	suite.Run(t, new(SaramaInjectTestSuite))
}

// #endregion

// ############################################################################
// #region TESTS ARE BELOW
// ############################################################################

func (s *SaramaInjectTestSuite) Test_GivenProperSASLSSLCerts_WhenInject_ThenGetConfigProperly() {
	// -- ARRANGE --
	// -- ACT --
	path := Inject(s.principal, []byte("kerbKeytab"), []byte("kerbConf"), []byte("sslCert"))
	certs, err := GetCertsFrom(path.SslCert)
	// -- ASSERT --
	assert.NotEqual(s.T(), "", path.KerbKeytab)
	assert.NotEqual(s.T(), "", path.KerbConf)
	assert.NotEqual(s.T(), "", path.SslCert)
	// Cert pool failure
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), certs)
	// just to test MissedLogger
	MissedLogger(5)
}

func (s *SaramaInjectTestSuite) Test_GivenNotProperSslCert_ThenReturnError() {
	// -- ARRANGE --
	// -- ACT --
	path := Inject(s.principal, []byte("kerbKeytab"), []byte("kerbConf"), []byte("sslCert"))
	certs, err := GetCertsFrom(path.SslCert)
	// -- ASSERT --
	assert.NotEqual(s.T(), "", path.KerbKeytab)
	assert.NotEqual(s.T(), "", path.KerbConf)
	assert.NotEqual(s.T(), "", path.SslCert)
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), certs)
}

func (s *SaramaInjectTestSuite) Test_GivenOsMkdirAllError_ThenPanic() {
	// -- ARRANGE --
	_osMkdirAll = func(path string, perm os.FileMode) error {
		return errForTests
	}
	// -- ACT --
	// -- ASSERT --
	assert.Panics(s.T(), func() {
		Inject(s.principal, []byte("kerbKeytab"), []byte("kerbConf"), []byte("sslCert"))
	})
}

func (s *SaramaInjectTestSuite) Test_GivenInjectKerbKeytabError_ThenPanic() {
	// -- ARRANGE --
	mockOsWriteFile(s.T())
	mpfWriteDir.SetCode("FPP")
	// -- ACT --
	// -- ASSERT --
	assert.Panics(s.T(), func() {
		Inject(s.principal, []byte("kerbKeytab"), []byte("kerbConf"), []byte("sslCert"))
	})
}

func (s *SaramaInjectTestSuite) Test_GivenInjectKerbConfError_ThenPanic() {
	// -- ARRANGE --
	mockOsWriteFile(s.T())
	mpfWriteDir.SetCode("PFP")
	// -- ACT --
	// -- ASSERT --
	assert.Panics(s.T(), func() {
		Inject(s.principal, []byte("kerbKeytab"), []byte("kerbConf"), []byte("sslCert"))
	})
}

func (s *SaramaInjectTestSuite) Test_GivenInjectSslCertError_ThenPanic() {
	// -- ARRANGE --
	mockOsWriteFile(s.T())
	mpfWriteDir.SetCode("PPF")
	// -- ACT --
	// -- ASSERT --
	assert.Panics(s.T(), func() {
		Inject(s.principal, []byte("kerbKeytab"), []byte("kerbConf"), []byte("sslCert"))
	})
}

func (s *SaramaInjectTestSuite) Test_GivenPathDoesNotExist_ThenReturnErr() {
	// -- ARRANGE --
	// -- ACT --
	certs, err := GetCertsFrom("somewhere over the rainbow")
	// -- ASSERT --
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), certs)
}

// ############################################################################
// #region TEST HELPERS
// ############################################################################

func (s *SaramaInjectTestSuite) resetMonkeyPatching() {
	_osMkdirAll = os.MkdirAll
	_osWriteFile = os.WriteFile
}

// #endregion

func mockOsWriteFile(_ *testing.T) {
	mpfWriteDir = testhelpers.NewMockPassFail()
	_osWriteFile = func(name string, data []byte, perm os.FileMode) error {
		err := mpfWriteDir.WillPassIncrementCount()
		return err
	}
}
