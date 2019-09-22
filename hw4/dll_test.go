package dll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrevNextErrors(t *testing.T) {
	var item *DLItem

	item, err := item.Prev()
	assert.Equal(t, item, (*DLItem)(nil))
	assert.EqualError(t, err, "Can't get previous item, because item is nil")

	item = nil
	item, err = item.Next()
	assert.Equal(t, item, (*DLItem)(nil))
	assert.EqualError(t, err, "Can't get next item, because item is nil")
}

func TestLen(t *testing.T) {
	var list DLList

	assert.Equal(t, 0, list.Len())

	var v interface{}

	list.PushFront(v)
	list.PushBack(v)
	list.PushBack(v)
	list.PushFront(v)
	list.PushFront(v)

	assert.Equal(t, 5, list.Len())

	list.Remove(list.First())

	assert.Equal(t, 4, list.Len())

	list.Remove(list.Last())
	assert.Equal(t, 3, list.Len())

	item, _ := list.First().Next()
	list.Remove(item)
	assert.Equal(t, 2, list.Len())
}
