package binarysearchtree

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				value: 0,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				value: 1,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				value: -1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.args.value)
			if got.value != tt.want {
				t.Errorf("NewBinarySearchTree() = %v, want %v", got, tt.want)
			}
			if got.left != nil {
				t.Errorf("NewBinarySearchTree() left = %v, want %v", got.left, nil)
			}
			if got.right != nil {
				t.Errorf("NewBinarySearchTree() right = %v, want %v", got.right, nil)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		rootValue   int
		insertValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test insert left",
			args: args{
				rootValue:   5,
				insertValue: 4,
			},
			want: 4,
		},
		{
			name: "test insert right",
			args: args{
				rootValue:   5,
				insertValue: 6,
			},
			want: 6,
		},
		{
			name: "test insert error",
			args: args{
				rootValue:   5,
				insertValue: 5,
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBinarySearchTree(tt.args.rootValue)
			err := tree.Insert(tt.args.insertValue)
			if tt.args.insertValue < tt.args.rootValue {
				if tree.left.value != tt.want {
					t.Errorf("Left value = %v, want %v", tree.left.value, tt.want)
				}
			}
			if tt.args.insertValue > tt.args.rootValue {
				if tree.right.value != tt.want {
					t.Errorf("Right value = %v, want %v", tree.left.value, tt.want)
				}
			}
			if tt.args.insertValue == tt.args.rootValue {
				if err == nil {
					t.Errorf("Insert() = %v, want %s", err, "error")
				}
			}
		})
	}
}
