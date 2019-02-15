package defaults

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.NotEmpty(t,Version,"empty version")
}
