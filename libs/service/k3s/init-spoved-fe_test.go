package k3s_service

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInitSpovedFe(t *testing.T) {
	err := InitSpovedFe(context.TODO())

	if err != nil {
		t.Error(err)
	}
}

func TestIsNotFound(t *testing.T) {
	err := fmt.Errorf(`services "root--spoved-fe-service" not found`)

	assert.Equal(t, isNotFound(err), true)
}
