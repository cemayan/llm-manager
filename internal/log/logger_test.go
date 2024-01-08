package log

import (
	"git-observer/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"os"
	"strconv"
	"testing"
)

type LoggerTestSuite struct {
	suite.Suite
}

func (suite *LoggerTestSuite) SetupSuite() {
	os.Setenv("BACKEND_ENV", "development")
	config.New()
	New()
}

func TestLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}

func (suite *LoggerTestSuite) Test_Level() {
	os.Setenv("BACKEND_LOG_LEVEL", strconv.Itoa(3))
	config.New()
	New()
	suite.Equal(logrus.WarnLevel, LoggerInstance.Logger.GetLevel())
}

func (suite *LoggerTestSuite) Test_Env() {
	os.Setenv("BACKEND_ENV", "production")
	config.New()
	New()
	suite.Equal(&logrus.JSONFormatter{}, LoggerInstance.Logger.Formatter)
}

func (suite *LoggerTestSuite) Test_Env2() {
	os.Setenv("BACKEND_ENV", "development")
	config.New()
	New()
	suite.Equal(&logrus.TextFormatter{}, LoggerInstance.Logger.Formatter)
}
