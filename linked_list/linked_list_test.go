package linkedlist

import "testing"

func TestInsert(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	if dll.Size() != 3 {
		t.Errorf("Insert method failed, expected size 3, got %d", dll.Size())
	}
}

func TestInsertAt(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.InsertAt(2,1)

	value, _ := dll.Get(1)
	if value != 2 {
		t.Errorf("Insert method failed, expected value 2 at position 1, got %d", value)
	}


	dll.InsertAt(0, 0)
	headVal, _ := dll.Head()
	if headVal != 0 {
		t.Errorf("Insert method failed, expected value 0 at position 0 (head), got %d", headVal)
	}

	dll.InsertAt(4, dll.Size())
	tailVal, _ := dll.Tail()
	if tailVal != 4 {
		t.Errorf("Insert method failed, expected value 4 at position %d (tail), got %d", dll.Size()-1, tailVal)
	}
}

func TestRemove(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)
	dll.Remove(2)

	if dll.Size() != 2 {
		t.Errorf("Remove method failed, expected size 2, got %d", dll.Size())
	}

	if dll.Contains(2) {
		t.Error("Remove method failed, value 2 should have been removed")
	}
}

func TestRemoveAt(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	dll.RemoveAt(1)

	if dll.Size() != 2 {
		t.Errorf("RemovedAt method failed, expected size 2, got %d", dll.Size())
	}

	if dll.Contains(2) {
		t.Error("RemoveAt method failed, value 2 should have beend removed")
	}
}

func TestGet(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	value, _ := dll.Get(1)
	if value != 2 {
		t.Errorf("Get method failed, expected value 2 at position 1, got %d", value)
	}
}

func TestContains(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	if !dll.Contains(2) {
		t.Error("Contains method failed, value 2 should be in the list")
	}

	if dll.Contains(4) {
		t.Error("Contains method failed, value 4 should not be in the list")
	}

}

func TestClear(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	dll.Clear()

	if !dll.IsEmpty() {
		t.Error("Clear method failed, list should be empty after clearing")
	}

	if dll.Size() != 0 {
		t.Errorf("Clear method failed, expected size 0 after clearing, got %d", dll.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	dll := &DoublyLinkedList{}

	if !dll.IsEmpty() {
		t.Error("IsEmpty method failed, list should be empty initially")
	}

	dll.Insert(1)

	if dll.IsEmpty() {
		t.Error("IsEmpty method failed, list should not be empty after Inserting an element")
	}
}

func TestReverse(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	dll.Reverse()

	headVal, _ := dll.Head()
	if headVal != 3 {
		t.Errorf("Reverse method failed, expected value 3 at head, got %d", headVal)
	}

	tailVal, _ := dll.Tail()
	if tailVal != 1 {
		t.Errorf("Reverse method failed, expected value 1 at tail, got %d", tailVal)
	}

	value, _ := dll.Get(1)
	if value != 2 {
		t.Errorf("Reverse method failed, expected value 2 at position 1, got %d", value)
	}
}

func TestHead(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	headVal, _ := dll.Head()
	if headVal != 1 {
		t.Errorf("Head method failed, expected value 1 at head, got %d", headVal)
	}
}

func TestTail(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	tailVal, _ := dll.Tail()
	if tailVal != 3 {
		t.Errorf("Tail method failed, expected value 3 at tail, got %d", tailVal)
	}
}
