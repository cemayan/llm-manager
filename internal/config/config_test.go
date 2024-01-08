package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (suite *ConfigTestSuite) SetupSuite() {
	New()
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (suite *ConfigTestSuite) Test_ConfigureParameters() {
	configureParameters()
	suite.Equal("development", AppConfig.Config.Environment)
	suite.Equal("", AppConfig.Config.Serve.Certificate)
	suite.Equal(0, len(AppConfig.Config.Security.AllowedOrigins))
	suite.Equal(log.InfoLevel, AppConfig.Config.Log.Level)
	suite.Equal("", AppConfig.Config.Serve.PrivateKey)
	suite.Equal(8996, AppConfig.Config.Serve.Port)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_LogLvl() {
	os.Clearenv()
	os.Setenv("BACKEND_LOG_LEVEL", "5")
	New()
	suite.Equal(log.DebugLevel, AppConfig.Config.Log.Level)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_LogLvl_Wrong() {
	os.Clearenv()
	os.Setenv("BACKEND_LOG_LEVEL", "aa")
	New()
	suite.Equal(log.InfoLevel, AppConfig.Config.Log.Level)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_Origins() {
	os.Clearenv()
	os.Setenv("BACKEND_ALLOWED_ORIGINS", "*")
	New()
	suite.Equal([]string{"*"}, AppConfig.Config.Security.AllowedOrigins)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_Env() {
	os.Clearenv()
	os.Setenv("BACKEND_ENV", "production")
	New()
	suite.Equal("production", AppConfig.Config.Environment)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_Port() {
	os.Clearenv()
	os.Setenv("BACKEND_PORT", "8888")
	New()
	suite.Equal(8888, AppConfig.Config.Serve.Port)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_PrivateKey() {
	os.Clearenv()
	os.Setenv("BACKEND_PRIVATE_KEY", "key")
	New()
	suite.Equal("key", AppConfig.Config.Serve.PrivateKey)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_Cert() {
	os.Clearenv()
	os.Setenv("BACKEND_CERT", "cert")
	New()
	suite.Equal("cert", AppConfig.Config.Serve.Certificate)
}
