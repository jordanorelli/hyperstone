package ent

import (
	"fmt"
	"io"
	"sort"

	"github.com/jordanorelli/hyperstone/bit"
)

type node interface {
	weight() int
	maxRank() int
}

// intermediate node
type iNode struct{ left, right node }

func (n iNode) String() string {
	return fmt.Sprintf("{%d %d *}", n.left.weight()+n.right.weight(), n.maxRank())
}

func (n iNode) weight() int { return n.left.weight() + n.right.weight() }
func (n iNode) maxRank() int {
	r := 0
	switch v := n.left.(type) {
	case iNode:
		r2 := v.maxRank()
		if r2 > r {
			r = r2
		}
	case lNode:
		if v.rank > r {
			r = v.rank
		}
	}
	switch v := n.right.(type) {
	case iNode:
		r2 := v.maxRank()
		if r2 > r {
			r = r2
		}
	case lNode:
		if v.rank > r {
			r = v.rank
		}
	}
	return r
}

// leaf node
type lNode struct {
	name string
	rank int
	freq int
	fn   func()
}

func (n lNode) String() string {
	return fmt.Sprintf("{%d %d %s}", n.freq, n.rank, n.name)
}

func (n lNode) weight() int  { return n.freq }
func (n lNode) maxRank() int { return n.rank }

// three-way comparison for nodes, used in sorting
func n_compare(n1, n2 node) int {
	switch {
	case n1.weight() < n2.weight():
		return -1
	case n1.weight() > n2.weight():
		return 1
	case n1.maxRank() < n2.maxRank():
		return -1
	case n1.maxRank() > n2.maxRank():
		return 1
	default:
		return 0
	}
}

// joins two nodes, creating and returning their parent node
func n_join(n1, n2 node) node {
	switch n_compare(n1, n2) {
	case -1:
		return iNode{n1, n2}
	case 0, 1:
		return iNode{n2, n1}
	default:
		panic("not reached")
	}
}

// a list of huffman nodes, for assembling a huffman tree
type nodeList []node

func (l nodeList) Len() int           { return len(l) }
func (l nodeList) Less(i, j int) bool { return n_compare(l[i], l[j]) == -1 }
func (l nodeList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func makeTree(l nodeList) node {
	// at each step:
	//   sort the list of nodes in the tree.
	//   remove the last two nodes, joining them to create their parent node.
	//   append this parent node to the list of nodes.
	// repeat until only one node remains, which is the root node.
	for len(l) > 1 {
		sort.Sort(sort.Reverse(l))
		l = append(l[:len(l)-2], n_join(l[len(l)-2], l[len(l)-1]))
	}
	return l[0]
}

func walk(n node, br bit.Reader) func() {
	switch v := n.(type) {
	case lNode:
		return v.fn
	case iNode:
		if bit.ReadBool(br) {
			return walk(v.right, br)
		} else {
			return walk(v.left, br)
		}
	default:
		panic("not reached")
	}
}

func dump(n node, prefix string, w io.Writer) {
	switch v := n.(type) {
	case lNode:
		fmt.Fprintf(w, "%s\t%v\n", prefix, v)
	case iNode:
		dump(v.left, prefix+"0", w)
		dump(v.right, prefix+"1", w)
	}
}

var hlist = nodeList{
	lNode{"PlusOne", 0, 36271, func() {}},
	lNode{"FieldPathEncodeFinish", 39, 25474, func() {}},
	lNode{"PushOneLeftDeltaNRightNonZeroPack6Bits", 11, 10530, func() {}},
	lNode{"PlusTwo", 1, 10334, func() {}},
	lNode{"PlusN", 4, 4128, func() {}},
	lNode{"PushOneLeftDeltaOneRightNonZero", 8, 2942, func() {}},
	lNode{"PopAllButOnePlusOne", 29, 1837, func() {}},
	lNode{"PlusThree", 2, 1375, func() {}},
	lNode{"PlusFour", 3, 646, func() {}},
	lNode{"PopAllButOnePlusNPack6Bits", 32, 634, func() {}},
	lNode{"PushOneLeftDeltaNRightZero", 9, 560, func() {}},
	lNode{"PushOneLeftDeltaOneRightZero", 7, 521, func() {}},
	lNode{"PushOneLeftDeltaNRightNonZero", 10, 471, func() {}},
	lNode{"PopAllButOnePlusNPack3Bits", 31, 300, func() {}},
	lNode{"NonTopoPenultimatePlusOne", 37, 271, func() {}},
	lNode{"PushOneLeftDeltaNRightNonZeroPack8Bits", 12, 251, func() {}},
	lNode{"PopAllButOnePlusN", 30, 149, func() {}},
	lNode{"NonTopoComplexPack4Bits", 38, 99, func() {}},
	lNode{"NonTopoComplex", 36, 76, func() {}},
	lNode{"PushOneLeftDeltaZeroRightZero", 5, 35, func() {}},
	lNode{"PushOneLeftDeltaZeroRightNonZero", 6, 3, func() {}},
	lNode{"PopNAndNonTopographical", 35, 1, func() {}},
	lNode{"PopNPlusN", 34, 0, func() {}},
	lNode{"PopNPlusOne", 33, 0, func() {}},
	lNode{"PopOnePlusN", 28, 0, func() {}},
	lNode{"PopOnePlusOne", 27, 2, func() {}},
	lNode{"PushNAndNonTopological", 26, 310, func() {}},
	lNode{"PushN", 25, 0, func() {}},
	lNode{"PushThreePack5LeftDeltaN", 24, 0, func() {}},
	lNode{"PushThreeLeftDeltaN", 23, 0, func() {}},
	lNode{"PushTwoPack5LeftDeltaN", 22, 0, func() {}},
	lNode{"PushTwoLeftDeltaN", 21, 0, func() {}},
	lNode{"PushThreePack5LeftDeltaOne", 20, 0, func() {}},
	lNode{"PushThreeLeftDeltaOne", 19, 0, func() {}},
	lNode{"PushTwoPack5LeftDeltaOne", 18, 0, func() {}},
	lNode{"PushTwoLeftDeltaOne", 17, 0, func() {}},
	lNode{"PushThreePack5LeftDeltaZero", 16, 0, func() {}},
	lNode{"PushThreeLeftDeltaZero", 15, 0, func() {}},
	lNode{"PushTwoPack5LeftDeltaZero", 14, 0, func() {}},
	lNode{"PushTwoLeftDeltaZero", 13, 0, func() {}},
}

var htree = makeTree(hlist)
