package linked_list

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("Create an empty LinkedList", func(t *testing.T) {
		actual := NewLinkedList()
		if actual.Length() != 0 {
			t.Errorf("Expected length of 0, got %d", actual.Length())
		}
	})

	t.Run("Append elements to LinkedList", func(t *testing.T) {
		l := NewLinkedList()
		l.Append(1)
		l.Append(2)
		l.Append(3)

		if l.Length() != 3 {
			t.Errorf("Expected length of 3, got %d", l.Length())
		}

		if l.Head().Value != 1 || l.Head().Next.Value != 2 || l.Tail().Value != 3 {
			t.Errorf("Nodes are not inserted correctly")
		}
	})

	t.Run("Prepend elements to LinkedList", func(t *testing.T) {
		l := NewLinkedList()
		l.Prepend(3)
		l.Prepend(2)
		l.Prepend(1)

		if l.Length() != 3 {
			t.Errorf("Expected length of 3, got %d", l.Length())
		}

		if l.Head().Value != 1 || l.Head().Next.Value != 2 || l.Tail().Value != 3 {
			t.Errorf("Nodes are not inserted correctly")
		}
	})

	t.Run("Create linked list from array", func(t *testing.T) {
		l := FromArray([]int{1, 2, 3})

		if l.Length() != 3 {
			t.Errorf("Expected length of 3, got %d", l.Length())
		}

		if l.Head().Value != 1 || l.Head().Next.Value != 2 || l.Tail().Value != 3 {
			t.Errorf("Nodes are not inserted correctly")
		}
	})

	t.Run("Create empty linkedlist from array", func(t *testing.T) {
		l := FromArray([]int{})
		assert.Equal(t, l.Length(), 0, "LinkedList from an empty array should have 0 length")
	})

	t.Run("Search LinkedList for existing value", func(t *testing.T) {
		l := buildTestLinkedList(t)
		searchVal := 7
		expected := 4
		actual := l.Search(searchVal)
		assert.Equal(t, expected, actual, "Index for Node with value of %d should be %d, got %d", searchVal, expected, actual)
	})

	t.Run("Search LinkedList for non-existing value", func(t *testing.T) {
		l := buildTestLinkedList(t)
		searchVal := 10
		expected := -1
		actual := l.Search(10)
		assert.Equal(t, expected, actual, "Index for Node with value of %d should be %d, got %d", searchVal, expected, actual)
	})

	t.Run("Get LinkedList Node based on valid index", func(t *testing.T) {
		l := buildTestLinkedList(t)
		searchIdx := 2
		expected := 5
		actual := l.Get(2)
		assert.Equal(t, expected, actual.Value, "Value for Node with index of %d should be %d, got %d", searchIdx, expected, actual.Value)
	})

	t.Run("Get LinkedList Node based on invalid index", func(t *testing.T) {
		l := buildTestLinkedList(t)
		actual := l.Get(10)
		assert.Nil(t, actual)
	})

	t.Run("Get from empty list returns nil", func(t *testing.T) {
		l := NewLinkedList()
		actual := l.Get(0)
		assert.Nil(t, actual)
	})

	t.Run("Insert at 0-th index empty LinkedList", func(t *testing.T) {
		expectedLength := 1
		l := NewLinkedList()
		l.Insert(0, 5)
		actualLength := l.Length()

		assert.Equal(t, expectedLength, actualLength, "Length after inserting one element at 0-th index into empty linked list should be %d, got %d", expectedLength, actualLength)
	})

	t.Run("Insert at non-zero index empty LinkedList", func(t *testing.T) {
		expectedLength := 0
		l := NewLinkedList()
		l.Insert(2, 5)
		actualLength := l.Length()

		assert.Equal(t, expectedLength, actualLength, "Length after inserting one element into empty linked list should be %d, got %d", expectedLength, actualLength)
	})

	t.Run("Handles insert at head", func(t *testing.T) {
		expectedLength := 6
		expectedValue := 5
		l := buildTestLinkedList(t)
		l.Insert(0, expectedValue)
		actualLength := l.Length()
		actualValue := l.Head().Value

		assert.Equal(t, expectedLength, actualLength, "Length after inserting at head should be %d, got %d", expectedLength, actualLength)
		assert.Equal(t, expectedValue, actualValue, "Value of inserted node should be %d, got %d", expectedValue, actualValue)
	})

	t.Run("Handles insert beyond tail", func(t *testing.T) {
		expectedLength := 6
		expectedValue := 5
		l := buildTestLinkedList(t)
		l.Insert(l.Length(), expectedValue)
		actualLength := l.Length()
		actualValue := l.Tail().Value

		assert.Equal(t, expectedLength, actualLength, "Length after inserting at head should be %d, got %d", expectedLength, actualLength)
		assert.Equal(t, expectedValue, actualValue, "Value of inserted node should be %d, got %d", expectedValue, actualValue)
	})

	t.Run("Handles insert", func(t *testing.T) {
		expectedLength := 6
		expectedValue := 5
		l := buildTestLinkedList(t)
		l.Insert(2, expectedValue)
		actualLength := l.Length()
		actualValue := l.Get(2).Value

		assert.Equal(t, expectedLength, actualLength, "Length after insertion should be %d, got %d", expectedLength, actualLength)
		assert.Equal(t, expectedValue, actualValue, "Value of inserted node should be %d, got %d", expectedValue, actualValue)
	})

	t.Run("Handles remove at head", func(t *testing.T) {
		expectedLength := 4
		expectedValue := 4
		l := buildTestLinkedList(t)
		l.Remove(0)
		actualLength := l.Length()
		actualValue := l.Head().Value

		assert.Equal(t, expectedLength, actualLength, "Length after removing at head should be %d, got %d", expectedLength, actualLength)
		assert.Equal(t, expectedValue, actualValue, "Value of new head node should be %d, got %d", expectedValue, actualValue)
	})

	t.Run("Handles remove at tail", func(t *testing.T) {
		expectedLength := 4
		expectedValue := 6
		l := buildTestLinkedList(t)
		l.Remove(l.Length() - 1)
		actualLength := l.Length()
		actualValue := l.Tail().Value

		assert.Equal(t, expectedLength, actualLength, "Length after removing at head should be %d, got %d", expectedLength, actualLength)
		assert.Equal(t, expectedValue, actualValue, "Value of new tail node should be %d, got %d", expectedValue, actualValue)
	})


}

func buildTestLinkedList(t testing.TB) *LinkedList {
	t.Helper()
	l := FromArray([]int{3, 4, 5, 6, 7})
	return l
}
