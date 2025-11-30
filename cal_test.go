package cal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	assert := assert.New(t)
	result, err := add(1,2)
	assert.Equal(result, 3)
	assert.NoError(err)
}
