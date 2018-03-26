package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/engbreaker/pkg/config"
)

func TestEnvImplDefault(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	impl := config.EnvImpl{}
	a.Equal("default", impl.GetString("OHYEAHHHHWHOWOULDuseThisKindofNAME", "default"))
}

func TestEnvImplGetString(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	impl := config.EnvImpl{}
	os.Setenv("asdf", "asdf")
	a.Equal("asdf", impl.GetString("asdf", "hohoho"))
}

func TestEnvImplInvalid(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	impl := config.EnvImpl{}
	a.Equal("", impl.GetString("12313", "default"))
	a.Equal("", impl.GetString("!@##@", "default"))
	a.Equal("", impl.GetString("한글", "default"))
}
