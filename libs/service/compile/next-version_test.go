package compile_service

import (
	"context"
	"fmt"
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

func TestFindLastOrderByProjectId1(t *testing.T) {
	co, err := FindLastOrderByProjectId(context.TODO(), 1)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v", co)
}

func TestFindLastOrderByProjectId2(t *testing.T) {
	co, err := FindLastOrderByProjectId(context.TODO(), 2)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v", co)
}
