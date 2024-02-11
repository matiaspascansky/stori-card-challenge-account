package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdGenerator(t *testing.T) {
	gen := NewUserIDGenerator()
	id := gen.GenerateID()
	idStr := strconv.FormatInt(id, 10)

	assert.NotEmpty(t, id)
	assert.Equal(t, 6, len(idStr))
}
