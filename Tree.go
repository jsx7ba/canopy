package canopy

import "cmp"

// Tree The interface for all tree types.
type Tree[E cmp.Ordered] interface {
	// Insert Places a value into the tree.
	// Returns true if the value was inserted, false if the value exists already.
	Insert(value E) bool

	// Delete Removes a value into the tree.
	// Returns true if the value was removed.
	Delete(value E) bool

	// Find Returns true if a value exists in the tree.
	Find(value E) bool
}
