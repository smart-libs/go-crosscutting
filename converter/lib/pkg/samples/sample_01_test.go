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

func TestWithoutFallback(t *testing.T) {
	// This Registry does not have the reflection fallback, so there must be a converter available to do the conversion.
	registry := converterdefault.NewRegistry()
	// This add method does not add the Generics fallback conversion functions
	registry.Add(converter.CreateHandler(funcs.FromStringToDate))
	myConverters := converterdefault.NewConverters(registry)

	d1 := "20240101"
	// This conversion from string to time.Time must be performed
	date1, err := myConverters.ConvertToType(d1, reflect.TypeFor[time.Time]())
	assert.NoError(t, err)
	assert.Equal(t, d1, date1.(time.Time).Format("20060102"))

	// The same conversion above using a different method, a helper function to cast to you
	date2, err2 := converter.To[time.Time](myConverters, d1)
	assert.NoError(t, err2)
	assert.Equal(t, d1, date2.Format("20060102"))

	// Because there is no fallback the conversion below should fail.
	pTDate3, err3 := converter.To[*time.Time](myConverters, d1)
	assert.Error(t, err3)
	assert.Nil(t, pTDate3)
}

func TestWithReflectionFallback(t *testing.T) {
	// This Registry does not have the reflection fallback, so there must be a converter available to do the conversion.
	registry := converterdefault.NewRegistry()
	// Decorates with the reflection fallback
	registry = converterdefault.NewFallbackRegistry(registry)
	// Add the conversion function
	registry.Add(converter.CreateHandler(funcs.FromStringToDate))
	myConverters := converterdefault.NewConverters(registry)

	d1 := "20240101"
	// This conversion from string to time.Time must be performed
	date1, err := myConverters.ConvertToType(d1, reflect.TypeFor[time.Time]())
	assert.NoError(t, err)
	assert.Equal(t, d1, date1.(time.Time).Format("20060102"))

	// The same conversion above using a different method, a helper function to cast to you
	date2, err2 := converter.To[time.Time](myConverters, d1)
	assert.NoError(t, err2)
	assert.Equal(t, d1, date2.Format("20060102"))

	// Because there is no fallback the conversion below should fail.
	pTDate3, err3 := converter.To[*time.Time](myConverters, d1)
	if assert.NoError(t, err3) {
		if assert.NotNil(t, pTDate3) {
			assert.Equal(t, d1, pTDate3.Format("20060102"))
		}
	}

	d2 := "20210101"
	date4 := time.Time{}
	assert.NoError(t, myConverters.Convert(d2, &date4))
	assert.Equal(t, d2, date4.Format("20060102"))

	var date5 *time.Time
	assert.NoError(t, myConverters.Convert(d2, &date5))
	assert.Equal(t, d2, (*date5).Format("20060102"))

	d3 := "19981123"
	assert.NoError(t, myConverters.Convert(&d3, &date5))
	assert.Equal(t, d3, (*date5).Format("20060102"))

	var d4 *string // nil
	assert.NoError(t, myConverters.Convert(d4, &date5))
	assert.Nil(t, date5)
}
