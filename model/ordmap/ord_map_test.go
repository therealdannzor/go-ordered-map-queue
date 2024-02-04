package ordmap

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdMap(t *testing.T) {
	o := New()

	// Test Update() and GetAll() on empty map
	o.Update("a", "1")
	expected := []string{"a,1\n"}
	actual := o.GetAll()
	assert.Equal(t, expected, actual)

	// Test Update() and GetAll() on non-empty map
	o.Update("b", "2")
	o.Update("c", "3")
	expected = []string{"a,1\n", "b,2\n", "c,3\n"}
	actual = o.GetAll()
	assert.Equal(t, expected, actual)

	// Test Update() on existing key with different value
	o.Update("a", "4")
	expected = []string{"b,2\n", "c,3\n", "a,4\n"}
	actual = o.GetAll()
	assert.Equal(t, expected, actual)

	// Test Update() on existing key with same value, no change
	o.Update("b", "2")
	assert.Equal(t, expected, actual)

	// Test Delete() on existing key
	o.Delete("b")
	expected = []string{"c,3\n", "a,4\n"}
	actual = o.GetAll()
	assert.Equal(t, expected, actual)

	// Test Delete() on non-existing key, no change
	o.Delete("d")
	assert.Equal(t, expected, actual)
}

func TestOrdMapConcurrency(t *testing.T) {
	o := New()
	var wg sync.WaitGroup
	const goroutines = 2000
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			k := fmt.Sprintf("key%s", strconv.Itoa(i))
			v := fmt.Sprintf("val%s", strconv.Itoa(i))
			o.Update(k, v)
		}(i)
	}
	wg.Wait()

	// Verify that all entries have been added successfully
	assert.Equal(t, goroutines, len(o.GetAll()))

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			k := fmt.Sprintf("key%s", strconv.Itoa(i))
			o.Delete(k)
		}(i)
	}
	wg.Wait()

	// Verify that all entries have been deleted successfully
	assert.True(t, o.list.Empty())
}
