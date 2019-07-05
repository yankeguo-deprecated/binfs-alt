package binfs

import "testing"

func TestNodeBasic(t *testing.T) {
	n := Node{}
	c := n.Child("a")
	if c == nil {
		t.Fatal("child not created")
	}
	if c.Name != "a" {
		t.Fatal("child name not set")
	}
	if len(c.Path) != 1 || c.Path[0] != "a" {
		t.Fatal("child path not set")
	}
	n.Ensure("a", "b", "c")
	n.Ensure("a", "c", "d")
	n.Ensure("a", "c", "d", "e")
	c = n.Find("a", "c", "d")
	if c == nil {
		t.Fatal("child not found")
	}
	if c.Name != "d" {
		t.Fatal("child name not set")
	}
	if len(c.Path) != 3 || c.Path[0] != "a" || c.Path[1] != "c" || c.Path[2] != "d" {
		t.Fatal("child path not set")
	}
	if len(c.Children) != 1 || c.Children["e"] == nil || c.Children["e"].Name != "e" {
		t.Fatal("child nodes not set")
	}
}
