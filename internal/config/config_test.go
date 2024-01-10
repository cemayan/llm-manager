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
	os.Setenv("LLM_BACKEND", "ollama")
	os.Setenv("BACKEND_LOG_LEVEL", "5")
	New()
	suite.Equal(log.DebugLevel, AppConfig.Config.Log.Level)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_LogLvl_Wrong() {
	os.Clearenv()
	os.Setenv("BACKEND_LOG_LEVEL", "aa")
	os.Setenv("LLM_BACKEND", "ollama")

	New()
	suite.Equal(log.InfoLevel, AppConfig.Config.Log.Level)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_Origins() {
	os.Clearenv()
	os.Setenv("BACKEND_ALLOWED_ORIGINS", "*")
	os.Setenv("LLM_BACKEND", "ollama")

	New()
	suite.Equal([]string{"*"}, AppConfig.Config.Security.AllowedOrigins)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_Env() {
	os.Clearenv()
	os.Setenv("BACKEND_ENV", "production")
	os.Setenv("LLM_BACKEND", "ollama")

	New()
	suite.Equal("production", AppConfig.Config.Environment)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_Port() {
	os.Clearenv()
	os.Setenv("BACKEND_PORT", "8888")
	os.Setenv("LLM_BACKEND", "ollama")

	New()
	suite.Equal(8888, AppConfig.Config.Serve.Port)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_PrivateKey() {
	os.Setenv("BACKEND_PRIVATE_KEY", "key")
	os.Setenv("LLM_BACKEND", "ollama")

	New()
	suite.Equal("key", AppConfig.Config.Serve.PrivateKey)
}

func (suite *ConfigTestSuite) Test_ConfigureParameters_Cert() {
	os.Setenv("BACKEND_CERT", "cert")
	os.Setenv("LLM_BACKEND", "ollama")

	New()
	suite.Equal("cert", AppConfig.Config.Serve.Certificate)
}
