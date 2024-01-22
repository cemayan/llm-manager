package log

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"llm-manager/internal/config"
	"os"
	"strconv"
	"testing"
)

type LoggerTestSuite struct {
	suite.Suite
}

func (suite *LoggerTestSuite) SetupSuite() {
	os.Setenv("BACKEND_ENV", "development")
	os.Setenv("LLM_BACKEND", "ollama")
	config.New()
	New()
}

func TestLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}

func (suite *LoggerTestSuite) Test_Level() {
	os.Setenv("BACKEND_LOG_LEVEL", strconv.Itoa(3))
	os.Setenv("LLM_BACKEND", "ollama")
	config.New()
	New()
	suite.Equal(logrus.WarnLevel, LoggerInstance.Logger.GetLevel())
}

func (suite *LoggerTestSuite) Test_Env() {
	os.Setenv("BACKEND_ENV", "production")
	os.Setenv("LLM_BACKEND", "ollama")
	config.New()
	New()
	suite.Equal(&logrus.JSONFormatter{}, LoggerInstance.Logger.Formatter)
}

func (suite *LoggerTestSuite) Test_Env2() {
	os.Setenv("BACKEND_ENV", "development")
	os.Setenv("LLM_BACKEND", "ollama")
	config.New()
	New()
	suite.Equal(&logrus.TextFormatter{}, LoggerInstance.Logger.Formatter)
}
