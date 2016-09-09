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
	lNode{"PlusOne", 0, 36271, func() { panic("not implemented: PlusOne") }},
	lNode{"FieldPathEncodeFinish", 39, 25474, func() { panic("not implemented: FieldPathEncodeFinish") }},
	lNode{"PushOneLeftDeltaNRightNonZeroPack6Bits", 11, 10530, func() { panic("not implemented: PushOneLeftDeltaNRightNonZeroPack6Bits") }},
	lNode{"PlusTwo", 1, 10334, func() { panic("not implemented: PlusTwo") }},
	lNode{"PlusN", 4, 4128, func() { panic("not implemented: PlusN") }},
	lNode{"PushOneLeftDeltaOneRightNonZero", 8, 2942, func() { panic("not implemented: PushOneLeftDeltaOneRightNonZero") }},
	lNode{"PopAllButOnePlusOne", 29, 1837, func() { panic("not implemented: PopAllButOnePlusOne") }},
	lNode{"PlusThree", 2, 1375, func() { panic("not implemented: PlusThree") }},
	lNode{"PlusFour", 3, 646, func() { panic("not implemented: PlusFour") }},
	lNode{"PopAllButOnePlusNPack6Bits", 32, 634, func() { panic("not implemented: PopAllButOnePlusNPack6Bits") }},
	lNode{"PushOneLeftDeltaNRightZero", 9, 560, func() { panic("not implemented: PushOneLeftDeltaNRightZero") }},
	lNode{"PushOneLeftDeltaOneRightZero", 7, 521, func() { panic("not implemented: PushOneLeftDeltaOneRightZero") }},
	lNode{"PushOneLeftDeltaNRightNonZero", 10, 471, func() { panic("not implemented: PushOneLeftDeltaNRightNonZero") }},
	lNode{"PushNAndNonTopological", 26, 310, func() { panic("not implemented: PushNAndNonTopological") }},
	lNode{"PopAllButOnePlusNPack3Bits", 31, 300, func() { panic("not implemented: PopAllButOnePlusNPack3Bits") }},
	lNode{"NonTopoPenultimatePlusOne", 37, 271, func() { panic("not implemented: NonTopoPenultimatePlusOne") }},
	lNode{"PushOneLeftDeltaNRightNonZeroPack8Bits", 12, 251, func() { panic("not implemented: PushOneLeftDeltaNRightNonZeroPack8Bits") }},
	lNode{"PopAllButOnePlusN", 30, 149, func() { panic("not implemented: PopAllButOnePlusN") }},
	lNode{"NonTopoComplexPack4Bits", 38, 99, func() { panic("not implemented: NonTopoComplexPack4Bits") }},
	lNode{"NonTopoComplex", 36, 76, func() { panic("not implemented: NonTopoComplex") }},
	lNode{"PushOneLeftDeltaZeroRightZero", 5, 35, func() { panic("not implemented: PushOneLeftDeltaZeroRightZero") }},
	lNode{"PushOneLeftDeltaZeroRightNonZero", 6, 3, func() { panic("not implemented: PushOneLeftDeltaZeroRightNonZero") }},
	lNode{"PopOnePlusOne", 27, 2, func() { panic("not implemented: PopOnePlusOne") }},
	lNode{"PopNAndNonTopographical", 35, 1, func() { panic("not implemented: PopNAndNonTopographical") }},

	// all the other operations have weights of 0 in clarity, which makes no
	// sense.
	lNode{"PopNPlusN", 34, 1, func() { panic("not implemented: PopNPlusN") }},
	lNode{"PopNPlusOne", 33, 1, func() { panic("not implemented: PopNPlusOne") }},
	lNode{"PopOnePlusN", 28, 1, func() { panic("not implemented: PopOnePlusN") }},
	lNode{"PushN", 25, 1, func() { panic("not implemented: PushN") }},
	lNode{"PushThreePack5LeftDeltaN", 24, 1, func() { panic("not implemented: PushThreePack5LeftDeltaN") }},
	lNode{"PushThreeLeftDeltaN", 23, 1, func() { panic("not implemented: PushThreeLeftDeltaN") }},
	lNode{"PushTwoPack5LeftDeltaN", 22, 1, func() { panic("not implemented: PushTwoPack5LeftDeltaN") }},
	lNode{"PushTwoLeftDeltaN", 21, 1, func() { panic("not implemented: PushTwoLeftDeltaN") }},
	lNode{"PushThreePack5LeftDeltaOne", 20, 1, func() { panic("not implemented: PushThreePack5LeftDeltaOne") }},
	lNode{"PushThreeLeftDeltaOne", 19, 1, func() { panic("not implemented: PushThreeLeftDeltaOne") }},
	lNode{"PushTwoPack5LeftDeltaOne", 18, 1, func() { panic("not implemented: PushTwoPack5LeftDeltaOne") }},
	lNode{"PushTwoLeftDeltaOne", 17, 1, func() { panic("not implemented: PushTwoLeftDeltaOne") }},
	lNode{"PushThreePack5LeftDeltaZero", 16, 1, func() { panic("not implemented: PushThreePack5LeftDeltaZero") }},
	lNode{"PushThreeLeftDeltaZero", 15, 1, func() { panic("not implemented: PushThreeLeftDeltaZero") }},
	lNode{"PushTwoPack5LeftDeltaZero", 14, 1, func() { panic("not implemented: PushTwoPack5LeftDeltaZero") }},
	lNode{"PushTwoLeftDeltaZero", 13, 1, func() { panic("not implemented: PushTwoLeftDeltaZero") }},
}

var htree = makeTree(hlist)
