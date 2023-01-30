package container_service

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNextVersion(t *testing.T) {
	v := "0.0.1"
	nv, e := nextVersion(v, "Patch")
	assert.Equal(t, e, nil)
	assert.Equal(t, nv, "0.0.2")

	nv, e = nextVersion(v, "Minor")
	assert.Equal(t, e, nil)
	assert.Equal(t, nv, "0.1.0")

	nv, e = nextVersion(v, "Major")
	assert.Equal(t, e, nil)
	assert.Equal(t, nv, "1.0.0")

	_, e = nextVersion(v, "Random")
	assert.NotEqual(t, e, nil)
}
