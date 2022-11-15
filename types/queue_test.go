package types

import (
	"reflect"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name string
		data []any
	}{
		{
			name: "enqueue - empty queue",
			data: []any{},
		},
		{
			name: "enqueue - single element queue",
			data: []any{1},
		},
		{
			name: "enqueue - queue with two elements",
			data: []any{1, 7},
		},
		{
			name: "enqueue - queue with multiple elements - 1",
			data: []any{1, 2, 3, 4},
		},
		{
			name: "enqueue - queue with multiple elements - 2",
			data: []any{1, 3, 4, 7, 9, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Q := &Queue{}

			for _, d := range tt.data {
				Q.Enqueue(d)
			}

			if Q.Size() != len(tt.data) {
				t.Errorf(`expected size of list to be %v, but got %v`, len(tt.data), Q.Size())
			}

			if Q.IsEmpty() != (Q.list.size == 0) {
				t.Errorf(`expected list to be have length %v, but got %v`, len(tt.data), Q.Size())
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	x := 1

	tests := []struct {
		name  string
		queue Queue
		want  any
	}{
		{
			name:  "dequeue - empty queue",
			queue: NewQueue(),
			want:  nil,
		},
		{
			name:  "dequeue - single element queue",
			queue: NewQueue(x),
			want:  x,
		},
		{
			name:  "dequeue - queue with two elements - 1",
			queue: NewQueue(x, 3),
			want:  x,
		},
		{
			name:  "dequeue - queue with multiple elements - 2",
			queue: NewQueue(x, 3, 5, 7, 8),
			want:  x,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.queue.Dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Front(t *testing.T) {
	tests := []struct {
		name  string
		queue Queue
		want  any
	}{
		{
			name:  "front - empty queue",
			queue: NewQueue(),
			want:  nil,
		},
		{
			name:  "front - queue of two elements",
			queue: NewQueue(1, 3),
			want:  1,
		},
		{
			name:  "front - queue of multiple elements",
			queue: NewQueue(1, 3, 4, 7),
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.queue.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Back(t *testing.T) {
	tests := []struct {
		name  string
		queue Queue
		want  any
	}{
		{
			name:  "back - empty queue",
			queue: NewQueue(),
			want:  nil,
		},
		{
			name:  "back - queue of two elements",
			queue: NewQueue(1, 3),
			want:  3,
		},
		{
			name:  "back - queue of multiple elements",
			queue: NewQueue(1, 3, 4, 7),
			want:  7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.queue.Back(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}
