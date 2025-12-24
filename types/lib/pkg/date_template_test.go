package types

import (
	"github.com/smart-libs/go-crosscutting/types/lib/pkg/metadata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateTemplate_UnmarshalJSON(t *testing.T) {
	type args struct {
		b string
	}
	type testCase[T metadata.Metadata[time.Time]] struct {
		name    string
		args    args
		wantErr func(t assert.TestingT, err error, d Date)
	}
	tests := []testCase[metadata.Metadata[time.Time]]{
		{
			name: "A date in the format YYMMDD must be parsed",
			args: args{b: `"220115"`},
			wantErr: func(t assert.TestingT, err error, d Date) {
				assert.Equal(t, "2022-01-15", d.String())
				assert.NoError(t, err)
			},
		},
		{
			name: "A date in the format YYYYMMDD must be parsed",
			args: args{b: `"20220115"`},
			wantErr: func(t assert.TestingT, err error, d Date) {
				assert.Equal(t, "2022-01-15", d.String())
				assert.NoError(t, err)
			},
		},
		{
			name: "A date in the format YYYY-MM-DD must be parsed",
			args: args{b: `"2022-01-15"`},
			wantErr: func(t assert.TestingT, err error, d Date) {
				assert.Equal(t, "2022-01-15", d.String())
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d Date
			tt.wantErr(t, d.UnmarshalJSON([]byte(tt.args.b)), d)
		})
	}
}
