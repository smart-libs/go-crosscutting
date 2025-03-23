package samples

import (
	converter "github.com/smart-libs/go-crosscutting/converter/lib/pkg"
	converterdefault "github.com/smart-libs/go-crosscutting/converter/lib/pkg/default"
	funcs "github.com/smart-libs/go-crosscutting/converter/lib/pkg/funcs"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestSample01(t *testing.T) {
	d1 := "20240101"
	d2 := "20210101"
	d3 := "19981123"

	registry := converterdefault.NewRegistry()
	t.Run("Adding converter without the Generics fallback conversion functions", func(t *testing.T) {
		// This addition method does not add the Generics fallback conversion functions which means the reflection fallbacks
		// will be used.
		registry.Add(converter.CreateHandler(funcs.FromStringToDate))
		myConverters := converterdefault.NewConverters(registry)

		date1, err := myConverters.ConvertToType(d1, reflect.TypeFor[time.Time]())
		assert.NoError(t, err)
		assert.Equal(t, d1, date1.(time.Time).Format("20060102"))

		// Using a helper function to cast to you
		date2, err2 := converter.To[time.Time](myConverters, d1)
		assert.NoError(t, err2)
		assert.Equal(t, d1, date2.Format("20060102"))

		date4 := time.Time{}
		assert.NoError(t, myConverters.Convert(d2, &date4))
		assert.Equal(t, d2, date4.Format("20060102"))

		var d4 *string // nil
		var date5 *time.Time
		assert.Error(t, myConverters.Convert(d4, &date5))

		// Because we added the conversion function without callbacks, then these conversions will fail
		// Using a helper function to cast to you
		_, err3 := converter.To[*time.Time](myConverters, d1)
		assert.Error(t, err3)

		assert.Error(t, myConverters.Convert(d2, &date5))

		assert.Error(t, myConverters.Convert(&d3, &date5))
	})

	t.Run("Adding conversion function with fallbacks", func(t *testing.T) {
		converter.AddHandler(registry, funcs.FromStringToDate)
		myConverters := converterdefault.NewConverters(registry)

		// Using a helper function to cast to you
		pTDate3, err3 := converter.To[*time.Time](myConverters, d1)
		if assert.NoError(t, err3) {
			assert.NotNil(t, d1, pTDate3)
			assert.Equal(t, d1, pTDate3.Format("20060102"))
		}

		var date5 *time.Time
		if assert.NoError(t, myConverters.Convert(d2, &date5)) {
			assert.Equal(t, d2, (*date5).Format("20060102"))
		}

		if assert.NoError(t, myConverters.Convert(&d3, &date5)) {
			assert.Equal(t, d3, (*date5).Format("20060102"))
		}
	})
}
