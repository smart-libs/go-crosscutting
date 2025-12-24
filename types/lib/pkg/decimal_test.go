package types

import (
	"encoding/json"
	"github.com/smart-libs/go-crosscutting/types/impl/decimal/shopspring/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	TestDecimal struct {
		shopspring.Metadata
	}
)

func (t TestDecimal) DataTypeName() string {
	return "UnitPrice"
}

func TestDecimalTemplate(t *testing.T) {
	var d Decimal
	assert.Equal(t, "0", d.String())
	assert.True(t, d.IsSet())
	assert.True(t, d.IsValid())
}

func TestDecimalTemplateFromString(t *testing.T) {
	var d2 DecimalTemplate[TestDecimal]
	d, err := DecimalFromString("3.4568")
	assert.NoError(t, err)
	assert.Equal(t, "3.4568", d.String())
	d2 = DecimalFromDecimal[TestDecimal](d)
	assert.Equal(t, "3.4568", d2.String())
}

func TestDecimalPow(t *testing.T) {
	d := DecimalFromInt(2)
	assert.Equal(t, "2", d.String())
	double := d.Pow(DecimalFromInt(2))
	assert.Equal(t, "4", double.String())
	root := double.Pow(DecimalFromFloat64(0.5))
	assert.Equal(t, "2", root.String())
}

func TestDecimalTemplateJSON(t *testing.T) {
	type MyStruct struct {
		Money Decimal `json:"money"`
	}

	var d MyStruct
	b, err := json.Marshal(d)
	assert.NoError(t, err)
	assert.Equal(t, `{"money":"0"}`, string(b))

	err = json.Unmarshal([]byte(`{"money":"3.678"}`), &d)
	assert.NoError(t, err)
	assert.Equal(t, "3.678", d.Money.String())
}
