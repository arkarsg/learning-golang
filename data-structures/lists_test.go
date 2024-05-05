package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"datastructures/lists"
)

func TestLinkedList(t *testing.T) {
	t.Run("Create an empty LinkedList", func(t *testing.T) {
		actual := lists.NewLinkedList()
		if actual.Length() != 0 {
			t.Errorf("Expected length of 0, got %d", actual.Length())
		}
	})

	t.Run("Append elements to LinkedList", func(t *testing.T) {
		l := lists.NewLinkedList()
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
		l := lists.NewLinkedList()
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
		l := lists.FromArray([]int{1, 2, 3})

		if l.Length() != 3 {
			t.Errorf("Expected length of 3, got %d", l.Length())
		}

		if l.Head().Value != 1 || l.Head().Next.Value != 2 || l.Tail().Value != 3 {
			t.Errorf("Nodes are not inserted correctly")
		}
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
}

func buildTestLinkedList(t testing.TB) *lists.LinkedList {
	t.Helper()
	l := lists.FromArray([]int{3, 4, 5, 6, 7})
	return l
}
