package ent

import (
	"fmt"
	"io"
	"sort"

	"github.com/jordanorelli/hyperstone/bit"
)

type node interface {
	weight() int
	rank() int
}

// intermediate node
type iNode struct {
	left  node
	right node
	_rank int
}

func (n iNode) String() string {
	return fmt.Sprintf("{%d %d *}", n.left.weight()+n.right.weight(), n.rank())
}

func (n iNode) weight() int { return n.left.weight() + n.right.weight() }

func (n iNode) rank() int { return n._rank }

// leaf node
type lNode struct {
	name  string
	_rank int
	freq  int
	fn    selectionOp
}

func (n lNode) String() string {
	return fmt.Sprintf("{%d %d %s}", n.freq, n._rank, n.name)
}

func (n lNode) weight() int { return n.freq }
func (n lNode) rank() int   { return n._rank }

// three-way comparison for nodes, used in sorting
func n_compare(n1, n2 node) int {
	switch {
	case n1.weight() < n2.weight():
		return -1
	case n1.weight() > n2.weight():
		return 1
	case n1.rank() < n2.rank():
		return 1
	case n1.rank() > n2.rank():
		return -1
	default:
		return 0
	}
}

// joins two nodes, creating and returning their parent node
func n_join(n1, n2 node, rank int) node {
	switch n_compare(n1, n2) {
	case -1:
		return iNode{n1, n2, rank}
	case 0, 1:
		return iNode{n2, n1, rank}
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
	for rank := len(l); len(l) > 1; rank++ {
		sort.Sort(sort.Reverse(l))
		l = append(l[:len(l)-2], n_join(l[len(l)-2], l[len(l)-1], rank))
	}
	return l[0]
}

// walks the huffman tree rooted at the node n, returning the selection op
// encoded by the tree.
func walk(n node, br bit.Reader) selectionOp {
	switch v := n.(type) {
	case lNode:
		// Debug.Printf("fieldpath fn: %s", v.name)
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
	lNode{"PlusOne", 0, 36271, func(r *selectionReader, br bit.Reader) {
		r.add(1)
	}},
	lNode{"FieldPathEncodeFinish", 39, 25474, nil},
	lNode{"PushOneLeftDeltaNRightNonZeroPack6Bits", 11, 10530, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushOneLeftDeltaNRightNonZeroPack6Bits")
	}},
	lNode{"PlusTwo", 1, 10334, func(r *selectionReader, br bit.Reader) {
		r.add(2)
	}},
	lNode{"PlusN", 4, 4128, func(r *selectionReader, br bit.Reader) {
		r.add(int(bit.ReadUBitVarFP(br)) + 5)
	}},
	lNode{"PushOneLeftDeltaOneRightNonZero", 8, 2942, func(r *selectionReader, br bit.Reader) {
		Debug.Printf("PushOneLeftDeltaOneRightNonZero")
		r.add(1)
		r.push(int(bit.ReadUBitVarFP(br)))
	}},
	lNode{"PopAllButOnePlusOne", 29, 1837, func(r *selectionReader, br bit.Reader) {
		r.pop(-1)
		r.add(1)
	}},
	lNode{"PlusThree", 2, 1375, func(r *selectionReader, br bit.Reader) {
		r.add(3)
	}},
	lNode{"PlusFour", 3, 646, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PlusFour")
	}},
	lNode{"PopAllButOnePlusNPack6Bits", 32, 634, func(r *selectionReader, br bit.Reader) {
		r.pop(-1)
		r.add(int(br.ReadBits(6)) + 1)
	}},
	lNode{"PushOneLeftDeltaNRightZero", 9, 560, func(r *selectionReader, br bit.Reader) {
		r.add(int(bit.ReadUBitVarFP(br)))
		r.push(0)
	}},
	lNode{"PushOneLeftDeltaOneRightZero", 7, 521, func(r *selectionReader, br bit.Reader) {
		r.add(1)
		r.push(0)
	}},
	lNode{"PushOneLeftDeltaNRightNonZero", 10, 471, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushOneLeftDeltaNRightNonZero")
	}},
	lNode{"PushNAndNonTopological", 26, 310, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushNAndNonTopological")
	}},
	lNode{"PopAllButOnePlusNPack3Bits", 31, 300, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PopAllButOnePlusNPack3Bits")
	}},
	lNode{"NonTopoPenultimatePlusOne", 37, 271, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: NonTopoPenultimatePlusOne")
	}},
	lNode{"PushOneLeftDeltaNRightNonZeroPack8Bits", 12, 251, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushOneLeftDeltaNRightNonZeroPack8Bits")
	}},
	lNode{"PopAllButOnePlusN", 30, 149, func(r *selectionReader, br bit.Reader) {
		r.pop(-1)
		r.add(int(bit.ReadUBitVarFP(br)) + 1)
	}},
	lNode{"NonTopoComplexPack4Bits", 38, 99, func(r *selectionReader, br bit.Reader) {
		r.måp(func(i int) int {
			if bit.ReadBool(br) {
				return i + int(br.ReadBits(4)) - 7 // ?!
			}
			return i
		})
	}},
	lNode{"NonTopoComplex", 36, 76, func(r *selectionReader, br bit.Reader) {
		r.måp(func(i int) int {
			if bit.ReadBool(br) {
				return i + int(bit.ReadZigZag32(br))
			}
			return i
		})
	}},
	lNode{"PushOneLeftDeltaZeroRightZero", 5, 35, func(r *selectionReader, br bit.Reader) {
		r.push(0)
	}},
	lNode{"PushOneLeftDeltaZeroRightNonZero", 6, 3, func(r *selectionReader, br bit.Reader) {
		r.push(int(bit.ReadUBitVarFP(br)))
	}},
	lNode{"PopOnePlusOne", 27, 2, func(r *selectionReader, br bit.Reader) {
		r.pop(1)
		r.add(1)
	}},
	lNode{"PopNAndNonTopographical", 35, 1, func(r *selectionReader, br bit.Reader) {
		r.count -= int(bit.ReadUBitVarFP(br))
		r.måp(func(i int) int {
			if bit.ReadBool(br) {
				return i + int(bit.ReadZigZag32(br))
			}
			return i
		})
	}},

	// all the other operations have weights of 0 in clarity, which makes no
	// sense.
	lNode{"PopNPlusN", 34, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PopNPlusN")
	}},
	lNode{"PopNPlusOne", 33, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PopNPlusOne")
	}},
	lNode{"PopOnePlusN", 28, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PopOnePlusN")
	}},
	lNode{"PushN", 25, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushN")
	}},
	lNode{"PushThreePack5LeftDeltaN", 24, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushThreePack5LeftDeltaN")
	}},
	lNode{"PushThreeLeftDeltaN", 23, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushThreeLeftDeltaN")
	}},
	lNode{"PushTwoPack5LeftDeltaN", 22, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushTwoPack5LeftDeltaN")
	}},
	lNode{"PushTwoLeftDeltaN", 21, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushTwoLeftDeltaN")
	}},
	lNode{"PushThreePack5LeftDeltaOne", 20, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushThreePack5LeftDeltaOne")
	}},
	lNode{"PushThreeLeftDeltaOne", 19, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushThreeLeftDeltaOne")
	}},
	lNode{"PushTwoPack5LeftDeltaOne", 18, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushTwoPack5LeftDeltaOne")
	}},
	lNode{"PushTwoLeftDeltaOne", 17, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushTwoLeftDeltaOne")
	}},
	lNode{"PushThreePack5LeftDeltaZero", 16, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushThreePack5LeftDeltaZero")
	}},
	lNode{"PushThreeLeftDeltaZero", 15, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushThreeLeftDeltaZero")
	}},
	lNode{"PushTwoPack5LeftDeltaZero", 14, 1, func(r *selectionReader, br bit.Reader) {
		r.push(int(br.ReadBits(5)))
		r.push(int(br.ReadBits(5)))
	}},
	lNode{"PushTwoLeftDeltaZero", 13, 1, func(r *selectionReader, br bit.Reader) {
		panic("not implemented: PushTwoLeftDeltaZero")
	}},
}

var huffRoot = makeTree(hlist)
