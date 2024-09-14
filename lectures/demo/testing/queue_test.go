package queue

import "testing"

func TestAppend(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Queue length is %d, want %d", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("Append(%d) failed", i)
		}
	}
	if q.Append(4) {
		t.Errorf("Append(4) succeeded, want failure as queue should be full")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should be able to retrieve item from queue")
		}
		if item != i {
			t.Errorf("Next() = %d, want %d", item, i)
		}
	}
	item, ok := q.Next()
	if ok {
		t.Errorf("Next() = %d, want failure as queue should be empty", item)
	}
}
