package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	data := ArrayInt{1, 3, 5, 7, 8}
	iterator := data.createIterator()
	// i 用于测试
	i := 0
	for iterator.HasNext() {
		assert.Equal(t, data[i], iterator.CurrentItem())
		iterator.Next()
		i++
	}
}
