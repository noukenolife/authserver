package helper

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestProjectRootPath(t *testing.T) {
	assert.MatchRegex(t, ProjectRootPath(), "/authserver$")
}
