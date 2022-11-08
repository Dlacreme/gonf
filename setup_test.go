package gonf

import (
	"testing"
)

func TestBuildEnvFilename(t *testing.T) {
	r := buildEnvFilename()
	e := "gonfig/dev.gonf"
	if r != e {
		setupError(t, "buildEnvFilename()", e, r)
	}
}

func TestBuildBaseFilename(t *testing.T) {
	r := buildBaseFilename()
	e := "gonfig/base.gonf"
	if r != e {
		setupError(t, "buildBaseFilename()", e, r)
	}
}

func TestSetFolder(t *testing.T) {
	SetFolder("hello")
	r := buildBaseFilename()
	e := "hello/base.gonf"
	if r != e {
		t.Error("SetFolder failed.")
		setupError(t, "SetFolder(hello)", e, r)
	}
}

func TestSetBase(t *testing.T) {
	SetFolder("gonfig")
	SetBase("esab")
	r := buildBaseFilename()
	e := "gonfig/esab"
	if r != e {
		setupError(t, "SetBase(esab)", e, r)
	}
}

func TestUse(t *testing.T) {
	Use("prod")
	r := buildEnvFilename()
	e := "gonfig/prod.gonf"
	if r != e {
		setupError(t, "Use(prod)", e, r)
	}
}

func setupError(t *testing.T, fn string, expect string, got string) {
	t.Errorf("%s failed.\nExpect: [%s]\nGot: [%s]", fn, expect, got)
}
