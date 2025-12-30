package serror

import (
	"errors"
	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsIllegalArgumentError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "returns true for IllegalArgument error",
			err:  errorx.IllegalArgument.New("test error"),
			want: true,
		},
		{
			name: "returns false for non-IllegalArgument error",
			err:  errors.New("regular error"),
			want: false,
		},
		{
			name: "returns false for nil error",
			err:  nil,
			want: false,
		},
		{
			name: "returns true for wrapped IllegalArgument error",
			err:  errorx.IllegalArgument.WrapWithNoMessage(errors.New("wrapped")),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIllegalArgumentError(tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsIllegalConfigError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "returns true for IllegalConfig error",
			err:  IllegalConfig.New("test error"),
			want: true,
		},
		{
			name: "returns false for non-IllegalConfig error",
			err:  errors.New("regular error"),
			want: false,
		},
		{
			name: "returns false for nil error",
			err:  nil,
			want: false,
		},
		{
			name: "returns false for IllegalArgument error",
			err:  errorx.IllegalArgument.New("test error"),
			want: false,
		},
		{
			name: "returns true for wrapped IllegalConfig error",
			err:  IllegalConfig.WrapWithNoMessage(errors.New("wrapped")),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIllegalConfigError(tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsTimeoutError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "returns true for timeout error",
			err:  errorx.TimeoutElapsed.New("timeout"),
			want: true,
		},
		{
			name: "returns false for non-timeout error",
			err:  errors.New("regular error"),
			want: false,
		},
		{
			name: "returns false for nil error",
			err:  nil,
			want: false,
		},
		{
			name: "returns true for wrapped timeout error",
			err:  errorx.TimeoutElapsed.WrapWithNoMessage(errors.New("wrapped")),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsTimeoutError(tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsNotFoundError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "returns true for NotFound error",
			err:  NotFoundError.New("not found"),
			want: true,
		},
		{
			name: "returns false for non-NotFound error",
			err:  errors.New("regular error"),
			want: false,
		},
		{
			name: "returns false for nil error",
			err:  nil,
			want: false,
		},
		{
			name: "returns true for wrapped NotFound error",
			err:  NotFoundError.WrapWithNoMessage(errors.New("wrapped")),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNotFoundError(tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsDuplicateError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "returns true for Duplicate error",
			err:  DuplicateError.New("duplicate"),
			want: true,
		},
		{
			name: "returns false for non-Duplicate error",
			err:  errors.New("regular error"),
			want: false,
		},
		{
			name: "returns false for nil error",
			err:  nil,
			want: false,
		},
		{
			name: "returns true for wrapped Duplicate error",
			err:  DuplicateError.WrapWithNoMessage(errors.New("wrapped")),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDuplicateError(tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIdentifyRootCause(t *testing.T) {
	testErr := errors.New("test error")
	illegalArgErr := errorx.IllegalArgument.New("illegal argument")
	timeoutErr := errorx.TimeoutElapsed.New("timeout")

	t.Run("returns false when error is nil", func(t *testing.T) {
		callbackInvoked := false
		fallbackInvoked := false

		got := IdentifyRootCause(nil,
			func(err error) { fallbackInvoked = true },
			CallbackCondition{
				Condition: func(err error) bool { return true },
				Callback:  func(err error) { callbackInvoked = true },
			},
		)

		assert.False(t, got)
		assert.False(t, callbackInvoked)
		assert.False(t, fallbackInvoked)
	})

	t.Run("returns false when no callbacks match and no fallback provided", func(t *testing.T) {
		callbackInvoked := false

		got := IdentifyRootCause(testErr,
			nil,
			CallbackCondition{
				Condition: func(err error) bool { return false },
				Callback:  func(err error) { callbackInvoked = true },
			},
		)

		assert.False(t, got)
		assert.False(t, callbackInvoked)
	})

	t.Run("returns true when fallback is invoked", func(t *testing.T) {
		callbackInvoked := false
		fallbackInvoked := false

		got := IdentifyRootCause(testErr,
			func(err error) { fallbackInvoked = true },
			CallbackCondition{
				Condition: func(err error) bool { return false },
				Callback:  func(err error) { callbackInvoked = true },
			},
		)

		assert.True(t, got)
		assert.False(t, callbackInvoked)
		assert.True(t, fallbackInvoked)
	})

	t.Run("returns true when first callback condition matches", func(t *testing.T) {
		firstCallbackInvoked := false
		secondCallbackInvoked := false
		fallbackInvoked := false

		got := IdentifyRootCause(illegalArgErr,
			func(err error) { fallbackInvoked = true },
			CallbackCondition{
				Condition: IsIllegalArgumentError,
				Callback:  func(err error) { firstCallbackInvoked = true },
			},
			CallbackCondition{
				Condition: IsTimeoutError,
				Callback:  func(err error) { secondCallbackInvoked = true },
			},
		)

		assert.True(t, got)
		assert.True(t, firstCallbackInvoked)
		assert.False(t, secondCallbackInvoked)
		assert.False(t, fallbackInvoked)
	})

	t.Run("returns true when second callback condition matches", func(t *testing.T) {
		firstCallbackInvoked := false
		secondCallbackInvoked := false
		fallbackInvoked := false

		got := IdentifyRootCause(timeoutErr,
			func(err error) { fallbackInvoked = true },
			CallbackCondition{
				Condition: IsIllegalArgumentError,
				Callback:  func(err error) { firstCallbackInvoked = true },
			},
			CallbackCondition{
				Condition: IsTimeoutError,
				Callback:  func(err error) { secondCallbackInvoked = true },
			},
		)

		assert.True(t, got)
		assert.False(t, firstCallbackInvoked)
		assert.True(t, secondCallbackInvoked)
		assert.False(t, fallbackInvoked)
	})

	t.Run("skips callback when condition is nil", func(t *testing.T) {
		callbackInvoked := false
		fallbackInvoked := false

		got := IdentifyRootCause(testErr,
			func(err error) { fallbackInvoked = true },
			CallbackCondition{
				Condition: nil,
				Callback:  func(err error) { callbackInvoked = true },
			},
		)

		assert.True(t, got)
		assert.False(t, callbackInvoked)
		assert.True(t, fallbackInvoked)
	})

	t.Run("skips callback when callback is nil even if condition matches", func(t *testing.T) {
		fallbackInvoked := false

		got := IdentifyRootCause(illegalArgErr,
			func(err error) { fallbackInvoked = true },
			CallbackCondition{
				Condition: IsIllegalArgumentError,
				Callback:  nil,
			},
		)

		assert.True(t, got)
		assert.True(t, fallbackInvoked)
	})

	t.Run("returns false when no callbacks provided and no fallback", func(t *testing.T) {
		got := IdentifyRootCause(testErr, nil)

		assert.False(t, got)
	})

	t.Run("returns true when no callbacks provided but fallback exists", func(t *testing.T) {
		fallbackInvoked := false

		got := IdentifyRootCause(testErr,
			func(err error) { fallbackInvoked = true },
		)

		assert.True(t, got)
		assert.True(t, fallbackInvoked)
	})

	t.Run("passes correct error to callback", func(t *testing.T) {
		var receivedErr error

		got := IdentifyRootCause(illegalArgErr,
			nil,
			CallbackCondition{
				Condition: IsIllegalArgumentError,
				Callback:  func(err error) { receivedErr = err },
			},
		)

		assert.True(t, got)
		assert.Equal(t, illegalArgErr, receivedErr)
	})

	t.Run("passes correct error to fallback", func(t *testing.T) {
		var receivedErr error

		got := IdentifyRootCause(testErr,
			func(err error) { receivedErr = err },
		)

		assert.True(t, got)
		assert.Equal(t, testErr, receivedErr)
	})
}
