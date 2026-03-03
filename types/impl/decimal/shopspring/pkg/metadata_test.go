package shopspring

import (
	"testing"

	"github.com/smart-libs/go-crosscutting/types/lib/pkg/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_Format(t *testing.T) {
	var metadata Metadata

	dec := metadata.FromFloat64(2334.8765)
	assert.Equal(t, "2334.8765", metadata.String(dec))
	assert.Equal(t, "2,334.8765", metadata.Format(dec, decimal.FormatOptions{Precision: -1}))
	assert.Equal(t, "2,334.88", metadata.Format(dec, decimal.FormatOptions{Precision: 2}))
}
