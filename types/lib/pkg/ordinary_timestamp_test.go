package types

import (
	"encoding/json"
	"fmt"
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TestValidFromMetadata struct {
	metadata.Date
}

type TestValidFrom struct {
	DateTemplate[TestValidFromMetadata]
}

func newTestValidFrom(t time.Time) TestValidFrom {
	return TestValidFrom{DateTemplate: NewDateTemplateFromTime[TestValidFromMetadata](t)}
}

func (v TestValidFromMetadata) DataTypeName() string { return "valid_from" }

func TestAssigningValidator(t *testing.T) {
	ts := NewOrdinaryTimestamp(time.Time{})
	assert.False(t, ts.IsSet())

	err := ts.ErrorIfNotSet()
	assert.Error(t, err)
	assert.Equal(t, "time.Time is not set", err.Error())

	err = NewTimestamp[TestValidFromMetadata](time.Time{}).ErrorIfNotSet()
	assert.Error(t, err)
	assert.Equal(t, "valid_from is not set", err.Error())

	type AnyStruct struct {
		ValidFrom TestValidFrom `json:"valid_from"`
	}
	ex := AnyStruct{ValidFrom: newTestValidFrom(time.Now())}

	b, err := json.Marshal(ex)
	if assert.NoError(t, err) {
		fmt.Print(string(b))
	}
}
