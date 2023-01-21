package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	t.Run("Add domain name", func(t *testing.T) {
		// Create a new tree.
		tree := NewTree()

		// Add a new domain name to the tree.
		tree.Add("foo.bar.baz.com")
		ok, wildcard := tree.Has("bar.baz.com")
		assert.True(t, ok)
		assert.False(t, wildcard)
	})

	t.Run("Add wildcard domain name", func(t *testing.T) {
		// Create a new tree.
		tree := NewTree()

		// Add a new domain name to the tree.
		tree.Add("*.bar.baz.com")
		ok, wildcard := tree.Has("foo.bar.baz.com")
		assert.True(t, ok)
		assert.True(t, wildcard)

		ok, wildcard = tree.Has("smth.foo.bar.baz.com")
		assert.True(t, ok)
		assert.True(t, wildcard)
	})

	t.Run("Remove domain name", func(t *testing.T) {
		// Create a new tree.
		tree := NewTree()

		// Add a new domain name to the tree.
		tree.Add("foo.bar.baz.com")
		ok, wildcard := tree.Has("foo.bar.baz.com")
		assert.True(t, ok)
		assert.False(t, wildcard)

		// Remove the domain name from the tree.
		tree.Remove("foo.bar.baz.com")
		ok, wildcard = tree.Has("foo.bar.baz.com")
		assert.False(t, ok)
		assert.False(t, wildcard)

		ok, wildcard = tree.Has("bar.baz.com")
		assert.False(t, ok)
		assert.False(t, wildcard)

		ok, wildcard = tree.Has("baz.com")
		assert.False(t, ok)
		assert.False(t, wildcard)

		ok, wildcard = tree.Has("com")
		assert.False(t, ok)
		assert.False(t, wildcard)
	})

	t.Run("Remove wildcard domain name", func(t *testing.T) {
		// Create a new tree.
		tree := NewTree()

		// Add a new domain name to the tree.
		tree.Add("*.bar.baz.com")
		ok, wildcard := tree.Has("foo.bar.baz.com")
		assert.True(t, ok)
		assert.True(t, wildcard)

		ok, wildcard = tree.Has("smth.foo.bar.baz.com")
		assert.True(t, ok)
		assert.True(t, wildcard)

		// Remove the domain name from the tree.
		tree.Remove("*.bar.baz.com")
		ok, wildcard = tree.Has("foo.bar.baz.com")
		assert.False(t, ok)
		assert.False(t, wildcard)

		ok, wildcard = tree.Has("bar.baz.com")
		assert.False(t, ok)
		assert.False(t, wildcard)

		ok, wildcard = tree.Has("baz.com")
		assert.False(t, ok)
		assert.False(t, wildcard)

	})

}

func Test_reverse(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name     string
		args     args
		expected []string
	}{
		{
			name:     "base case",
			args:     args{[]string{"foo", "bar", "baz"}},
			expected: []string{"baz", "bar", "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.s)
			assert.Equal(t, tt.expected, tt.args.s)
		})
	}
}
