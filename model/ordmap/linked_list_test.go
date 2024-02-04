package ordmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()

	// Test Empty() on empty list
	assert.True(t, ll.Empty())

	// Test Append() and ToSlice() on empty list
	ll.Append("a")
	expected := []string{"a"}
	actual := ll.ToSlice()
	assert.Equal(t, expected, actual)

	// Test Append() and ToSlice() on non-empty list
	ll.Append("b")
	ll.Append("c")
	expected = []string{"a", "b", "c"}
	actual = ll.ToSlice()
	assert.Equal(t, expected, actual)

	// Test Delete() on non-empty list
	ll.Delete("b")
	expected = []string{"a", "c"}
	actual = ll.ToSlice()
	assert.Equal(t, expected, actual)

	// Test Delete() on non-existent data, no change from previous case
	ll.Delete("d")
	assert.Equal(t, expected, actual)

	// Test Swap() on non-empty list
	ll.Swap("a")
	expected = []string{"c", "a"}
	actual = ll.ToSlice()
	assert.Equal(t, expected, actual)

	// Test Swap() on non-existent data, no change from previous case
	ll.Swap("d")
	assert.Equal(t, expected, actual)
}
