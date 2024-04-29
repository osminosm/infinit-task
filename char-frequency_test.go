package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharFrequency(t *testing.T) {

	t.Run("should count occurences of chars in string (example 1)", func(t *testing.T) {

		freq := NewCharFrequency("AaBbCC;")
		values := freq.Values()

		assert.Equal(t, len(values), 6)
		assert.Contains(t, values, Frequency{Char: 'A', Count: 1})
		assert.Contains(t, values, Frequency{Char: 'a', Count: 1})
		assert.Contains(t, values, Frequency{Char: 'B', Count: 1})
		assert.Contains(t, values, Frequency{Char: 'b', Count: 1})
		assert.Contains(t, values, Frequency{Char: 'C', Count: 2})
		assert.Contains(t, values, Frequency{Char: ';', Count: 1})

	})

	t.Run("should count occurences of chars in string (example 2)", func(t *testing.T) {

		freq := NewCharFrequency("\n$**")
		values := freq.Values()

		assert.Equal(t, len(values), 3)
		assert.Contains(t, values, Frequency{Char: '\n', Count: 1})
		assert.Contains(t, values, Frequency{Char: '$', Count: 1})
		assert.Contains(t, values, Frequency{Char: '*', Count: 2})

	})

}
