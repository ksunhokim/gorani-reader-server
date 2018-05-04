package config_test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/pkg/config"
	"github.com/sunho/gorani-reader/pkg/util"
)

func init() {
	impl := config.EnvImpl{}
	config.Impl = impl
	logrus.SetOutput(util.DummyWriter{})
	os.Setenv("string", "string")
	os.Setenv("int", "1234")
	os.Setenv("iint", "-1234")
}

func TestConfigGetString(t *testing.T) {
	a := assert.New(t)

	a.Equal("string", config.GetString("string", "default"))
	a.Equal("default", config.GetString("WHOOOHOOADIDDOCERTJDSF", "default"))
}

func TestConfigGetInt(t *testing.T) {
	a := assert.New(t)

	a.Equal(1234, config.GetInt("int", 5124))
	a.Equal(-1234, config.GetInt("iint", 5124))
	a.Equal(1234, config.GetInt("JASGJSGKLDLKGL1232", 1234))
	a.Equal(0, config.GetInt("string", 1234))
}
