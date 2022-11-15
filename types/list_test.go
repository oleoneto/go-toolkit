package types

import (
	"reflect"
	"testing"
)

func TestList_Insert(t *testing.T) {
	tests := []struct {
		name string
		args []*Node
	}{
		{
			name: "nodes - 0",
			args: []*Node{},
		},
		{
			name: "nodes - 1",
			args: []*Node{{Data: 1}},
		},
		{
			name: "nodes - 2",
			args: []*Node{{Data: 1}, {Data: 2}, {Data: 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := &List{}

			for _, node := range tt.args {
				L.Insert(node)
			}

			if L.Size() != len(tt.args) {
				t.Errorf(`expected size of list to be %v, but got %v`, len(tt.args), L.Size())
			}

			if L.IsEmpty() != (L.size == 0) {
				t.Errorf(`expected list to be have length %v, but got %v`, len(tt.args), L.Size())
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	type args struct {
		matchingFunc func(*Node) bool
	}

	tests := []struct {
		name         string
		data         []*Node
		args         args
		expectedSize int
	}{
		{
			name: "remove - empty list",
			data: []*Node{},
			args: args{
				matchingFunc: func(n *Node) bool { return false },
			},
			expectedSize: 0,
		},
		{
			name: "remove - single element list",
			data: []*Node{{Data: 1}},
			args: args{
				matchingFunc: func(n *Node) bool { return false },
			},
			expectedSize: 1,
		},
		{
			name: "remove - first and only element - 1",
			data: []*Node{{Data: 1}},
			args: args{
				matchingFunc: func(n *Node) bool { return n.next == nil },
			},
			expectedSize: 0,
		},
		{
			name: "remove - first and only element - 2",
			data: []*Node{{Data: 1}},
			args: args{
				matchingFunc: func(n *Node) bool { return n.previous == nil },
			},
			expectedSize: 0,
		},
		{
			name: "remove - missing element",
			data: []*Node{{Data: 1}, {Data: 5}},
			args: args{
				matchingFunc: func(n *Node) bool {
					if v, ok := n.Data.(int); ok {
						return v == 4
					}

					return false
				},
			},
			expectedSize: 2,
		},
		{
			name: "remove - middle element",
			data: []*Node{{Data: 4}, {Data: 10}, {Data: 20}, {Data: 30}, {Data: 40}},
			args: args{
				matchingFunc: func(n *Node) bool {
					if v, ok := n.Data.(int); ok {
						return v == 20
					}

					return false
				},
			},
			expectedSize: 4,
		},
		{
			name: "remove - first element of many - 1",
			data: []*Node{{Data: 4}, {Data: 10}, {Data: 20}, {Data: 30}, {Data: 40}},
			args: args{
				matchingFunc: func(n *Node) bool {
					if v, ok := n.Data.(int); ok {
						return v == 4
					}

					return false
				},
			},
			expectedSize: 4,
		},
		{
			name: "remove - first element of many - 2",
			data: []*Node{{Data: 4}, {Data: 10}, {Data: 20}, {Data: 30}, {Data: 40}},
			args: args{
				matchingFunc: func(n *Node) bool {
					return n.previous == nil
				},
			},
			expectedSize: 4,
		},
		{
			name: "remove - last element of many - 1",
			data: []*Node{{Data: 4}, {Data: 10}, {Data: 20}, {Data: 30}, {Data: 40}},
			args: args{
				matchingFunc: func(n *Node) bool {
					if v, ok := n.Data.(int); ok {
						return v == 40
					}

					return false
				},
			},
			expectedSize: 4,
		},
		{
			name: "remove - last element of many - 2",
			data: []*Node{{Data: 4}, {Data: 10}, {Data: 20}, {Data: 30}, {Data: 40}},
			args: args{
				matchingFunc: func(n *Node) bool {
					// Remove the last element
					return n.next == nil
				},
			},
			expectedSize: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &List{}

			for _, node := range tt.data {
				list.Insert(node)
			}

			list.Remove(tt.args.matchingFunc)

			if list.Size() != tt.expectedSize {
				t.Errorf(`expected size to be %v, but got %v`, tt.expectedSize, list.Size())
			}
		})
	}
}

func TestList_Find(t *testing.T) {
	type args struct {
		matchingFunc func(*Node) bool
	}

	tests := []struct {
		name      string
		data      []*Node
		list      *List
		args      args
		wantFound bool
	}{
		{
			name: "nodes - 0",
			list: NewList(),
			args: args{
				matchingFunc: func(n *Node) bool {
					if v, ok := n.Data.(int); ok {
						return v == 20
					}

					return false
				},
			},
			wantFound: false,
		},
		{
			name: "nodes - 1",
			list: NewList(),
			args: args{
				matchingFunc: func(n *Node) bool {
					return true
				},
			},
			wantFound: false,
		},
		{
			name: "nodes - 2",
			list: NewList(4, 10, 20, 30, 40),
			args: args{
				matchingFunc: func(n *Node) bool {
					if v, ok := n.Data.(int); ok {
						return v == 20
					}

					return false
				},
			},
			wantFound: true,
		},
		{
			name: "nodes - 3",
			list: NewList(4, 10, 20, 30, 40),
			args: args{
				matchingFunc: func(n *Node) bool {
					if v, ok := n.Data.(int); ok {
						return v == 9
					}

					return false
				},
			},
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotFound := tt.list.Find(tt.args.matchingFunc)
			// if !reflect.DeepEqual(gotResult, tt.wantResult) {
			// 	t.Errorf("List.Find() gotResult = %v, want %v", gotResult, tt.wantResult)
			// }

			if gotFound != tt.wantFound {
				t.Errorf("List.Find() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestList_GetHead(t *testing.T) {
	x := &Node{Data: 42}

	tests := []struct {
		name string
		list *List
		want *Node
	}{
		{
			name: "test - 0",
			list: NewList(),
			want: nil,
		},
		{
			name: "test - 1",
			list: NewList(x),
			want: x,
		},
		{
			name: "test - 2",
			list: NewList(x, 2, 3, 4),
			want: x,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.GetHead(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.GetHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_GetTail(t *testing.T) {
	x := &Node{Data: 42}

	tests := []struct {
		name string
		data []*Node
		list *List
		want *Node
	}{
		{
			name: "find - empty list",
			list: NewList(),
			want: nil,
		},
		{
			name: "find - single element list",
			list: NewList(x),
			want: x,
		},
		{
			name: "find - last element of many",
			list: NewList(1, 2, 3, x),
			want: x,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.GetTail(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.GetTail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Reverse(t *testing.T) {
	tests := []struct {
		name string
		list *List
	}{
		{
			name: "reverse - empty list",
			list: NewList(),
		},
		{
			name: "reverse - single element list",
			list: NewList(1),
		},
		{
			name: "reverse - list with two elements",
			list: NewList(1, 2),
		},
		{
			name: "reverse - list with multiple elements - 1",
			list: NewList(1, 2, 3),
		},
		{
			name: "reverse - list with multiple elements - 2",
			list: NewList(1, 2, 3, 4),
		},
		{
			name: "reverse - list with multiple elements - 3",
			list: NewList(1, 2, 3, 4, 5, 6, 7, 8, 9),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := tt.list.GetHead()
			tail := tt.list.GetTail()
			size := tt.list.Size()

			tt.list.Reverse()

			// Reversed list should have the same number of elements
			if tt.list.Size() != size {
				t.Errorf(`expected size to be %v, but got %v`, size, tt.list.Size())
			}

			if !tt.list.IsEmpty() {
				// New tail should be the previous head
				if !reflect.DeepEqual(head, tt.list.GetTail()) {
					t.Errorf(`expected head to be %v, but got %v`, head, tt.list.GetTail())
				}

				// New head should be the previous tail
				if !reflect.DeepEqual(tail, tt.list.GetHead()) {
					t.Errorf(`expected tail to be %v, but got %v`, tail, tt.list.GetHead())
				}
			}
		})
	}
}

func TestList_ToArray(t *testing.T) {
	tests := []struct {
		name       string
		list       *List
		wantResult []any
	}{
		{
			name:       "toarray - empty list",
			list:       NewList(),
			wantResult: []any{},
		},
		{
			name:       "toarray - single element list",
			list:       NewList(49),
			wantResult: []any{49},
		},
		{
			name:       "toarray - list of two elements",
			list:       NewList(49, 3),
			wantResult: []any{49, 3},
		},
		{
			name:       "toarray - list of multiple elements",
			list:       NewList(49, 3, 7, 10, 90, 21, 14, 7, 3),
			wantResult: []any{49, 3, 7, 10, 90, 21, 14, 7, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.list.ToArray(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("List.ToArray() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
