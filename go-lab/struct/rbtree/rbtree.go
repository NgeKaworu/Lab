package rbtree

import (
	"fmt"
	"math"
)

// 红黑树规则
// 1. 节点是红色或黑色。
// 2. 根是黑色。
// 3. 所有叶子都是黑色（叶子是NIL节点）。
// 4. 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
// 5. 从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
// PS: 本文中我们使用"nil叶子"或"空（null）叶子".
// https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91

// RBTree 红黑树
type RBTree struct {
	root *TreeNode
}

// TreeNode 树节点
type TreeNode struct {
	key                 int
	left, right, parent *TreeNode
	red                 bool
}

func (n *TreeNode) insert(k *TreeNode) {
	// 绑个父元素
	k.parent = n
	// DFS 深度优先
	if k.key <= n.key {
		// 小于等于走左边
		if n.left == nil {
			// 左树为空, 左树即是节点
			n.left = k
		} else {
			// 递归左树
			n.left.insert(k)
		}
	} else {
		// 大于走右边
		if n.right == nil {
			// 右树为空, 右树即节点
			n.right = k
		} else {
			// 递归右树
			n.right.insert(k)
		}
	}
}

// 返回叔父节点
func (n *TreeNode) uncle() *TreeNode {

	// 无祖父 或 无父 说明是根 或 根的左右子树
	if n.parent == nil || n.grandpa() == nil {
		return nil
	}

	// 返回叔父
	if n.parent == n.grandpa().left {
		return n.grandpa().right
	}

	return n.grandpa().left

}

// 返回兄弟节点
func (n *TreeNode) sibling() *TreeNode {
	p := n.parent
	if p != nil {
		if p.left == n {
			return p.right
		}
		return p.left
	}
	return nil
}

// 返回祖父节点
func (n *TreeNode) grandpa() *TreeNode {
	// 节点不是根
	if n.parent != nil {
		return n.parent.parent
	}

	return nil
}

// 用目标替换当前节点
func (t *RBTree) replace(old, new *TreeNode) {
	if new != nil {
		// 双向链 记得 改绑 父元素
		new.parent = old.parent
	}

	if old == t.root {
		// 如果要替换的是root
		t.root = new
	} else {
		p := old.parent
		if p.left == old {
			p.left = new
		}

		if p.right == old {
			p.right = new
		}
	}
}

// 左旋
func (t *RBTree) leftRotate(pivot *TreeNode) *RBTree {
	if pivot == t.root {
		return t
	}
	p := pivot.parent
	// 父节点的右节点 指向目标(pviot)节点的左节点
	p.right = pivot.left
	// 记得双向绑定
	if pivot.left != nil {
		pivot.left.parent = p
	}
	// 目标(pviot)节点的左节点 指向 父节点
	pivot.left = p
	// 处理双向绑定
	pivot.parent = nil
	// 用目标节点替换成父节点
	t.replace(p, pivot)
	// 最后还需要在处理一次双向绑定
	p.parent = pivot
	return t
}

// 右旋
func (t *RBTree) rightRotate(pivot *TreeNode) *RBTree {
	if pivot == t.root {
		return t
	}
	p := pivot.parent
	// 父节点的左节点 指向目标(pviot)节点的右节点
	p.left = pivot.right
	if pivot.right != nil {
		// 处理双向绑定
		pivot.right.parent = p
	}
	// 目标(pviot)节点的右节点指向父节点
	pivot.right = p
	// 处理双向绑定
	pivot.parent = nil
	// 用目标节点替换成父节点
	t.replace(p, pivot)
	// 最后还需要在处理一次双向绑定
	p.parent = pivot
	return t
}

// BFS 广度优先 打印
func (t *RBTree) String() (str string) {
	// 初始化队列
	queue := []*TreeNode{t.root}
	count, line := 0.0, 1.0
	for true {
		// 退出条件
		// 如果某一行的下一行都是空
		if float64(len(queue)) == math.Pow(2, line) {
			every := true
			for _, v := range queue {
				if v != nil {
					every = false
					break
				}
			}
			if every == true {
				break
			}
		}
		// 插入换行
		// 高度h = h(n) = log2(n) + 1, n是节点数
		h := math.Floor(math.Log2(count + 1))
		if h == line {
			str += "\n"
			line++
		}
		count++

		// pop
		leaf := queue[0]
		queue = queue[1:]
		if leaf != nil {
			str += fmt.Sprintf("(%v, %v) ", leaf.key, leaf.red)
			queue = append(queue, leaf.left, leaf.right)
		} else {
			str += " nil "
			// 保证格式一致
			queue = append(queue, nil, nil)
		}

	}
	return
}

// Find 查找
func (t *RBTree) Find(k int) *TreeNode {
	cur := t.root
	for cur != nil {
		if k > cur.key {
			cur = cur.right
		} else if k < cur.key {
			cur = cur.left
		} else {
			return cur
		}
	}
	return nil
}

// Insert 插入 DFS
func (t *RBTree) Insert(k int) *RBTree {
	n := &TreeNode{key: k, red: true}
	if t.root == nil {
		t.root = n
	} else {
		t.root.insert(n)
	}
	// 插入自平衡
	t.insertBlance(n)

	return t
}

// 插入自平衡
func (t *RBTree) insertBlance(n *TreeNode) {
	if n.parent == nil {
		// 如果是根节点
		// 根据: 2. 根是黑色。
		// 设为黑色
		n.red = false
		return

	}

	// 不是根节点的逻辑
	// 因为插入节点是红色
	// 父节点是红色
	// 则与规则4冲突 —— 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
	if n.parent.red == true {
		// 如果 父、叔父都是红色
		if n.uncle() != nil && n.uncle().red == true {
			// 把父、叔、爷节点变色
			n.parent.red = false
			n.uncle().red = false
			n.grandpa().red = true
			// 递归检查 爷节点
			t.insertBlance(n.grandpa())
		} else {
			// 这个场景是叔节点为空或是黑色
			// 假设我们的树不是“顺拐” 要先做一次旋转, 旋转规则如下:
			if n == n.parent.right && n.parent == n.grandpa().left {
				// 如果目标在父节点右边, 且父节点在爷节点左边, 左旋
				//         黑             红
				//        / \     染色    / \
				//       红  黑    ->    黑  黑
				//      /  \           /  \
				//     Nil (红)       Nil (红)
				t.leftRotate(n)
				// 因为左旋, 这里n.left是之前的父节点
				n = n.left
			} else if n == n.parent.left && n.parent == n.grandpa().right {
				// 如果目标在父节点左边, 且父节点在爷节点右边, 右旋
				t.rightRotate(n)
				// 因为右旋, 这里n.right是之前的父节点
				n = n.right
			}

			// 第二次旋转, 先把父元素染黑、爷元素染红
			n.parent.red = false
			n.grandpa().red = true

			if n == n.parent.left && n.parent == n.grandpa().left {
				// 当目标在父节点左侧且父也在爷左时
				//         黑             红               黑
				//        / \     染色    / \      右旋    / \
				//       红  黑    ->    黑  黑     -->  (红) 红
				//      /               /                     \
				//     (红)           (红)                     黑
				// 右旋其父
				t.rightRotate(n.parent)
			} else {
				// 否则就是述情况的镜像
				// 左旋其父
				t.leftRotate(n.parent)
			}
		}
	}

}

// Remove 删除
/** 删除逻辑
我们先简化问题, 把删除两个儿子节点的问题统一成删除一个儿子的问题;
对于二叉查找树，在删除带有两个非叶子儿子的节点的时候，
我们要么找到它左子树中的最大元素、要么找到它右子树中的最小元素，
并把它的值转移到要删除的节点中。
我们接着删除我们从中复制出值的那个节点，它必定有少于两个非叶子的儿子。
因为只是复制了一个值（没有复制颜色），不违反任何性质，
这就把问题简化为如何删除最多有一个儿子的节点的问题。
它不关心这个节点是最初要删除的节点还是我们从中复制出值的那个节点。
*/
func (t *RBTree) Remove(k int) bool {
	n := t.Find(k)
	if n == nil {
		return false
	}
	// 找到节点, 节点没有右节点, 只删除一个儿子成立
	if n.right == nil {
		t.RemoveOne(n)
	} else {
		// 找到它右子树中的最小元素
		// 最小值没有左树, 只删除一个儿子成立
		smallest := n.right
		for smallest.left != nil {
			smallest = smallest.left
		}
		// 用最小值替换目标值
		// PS: 因为只是复制了一个值（没有复制颜色），不违反任何性质
		n.key = smallest.key
		// 删除最小值的节点
		t.RemoveOne(smallest)
	}

	return true
}

// RemoveOne 删除只有一个儿子的节点
func (t *RBTree) RemoveOne(n *TreeNode) {
	child := n.left
	if n.right != nil {
		child = n.right
	}

	// n是root
	if n == t.root {
		// 无子
		if child == nil {
			n = nil
			t.root = nil
		} else {
			// 有子
			n = nil
			t.root = child
			// 根必须是黑
			child.red = false
		}

		return
	}

	// 用child代替n的位置
	p := n.parent
	if n == p.left {
		p.left = child
	} else {
		p.right = child
	}
	if child != nil {
		child.parent = p
	}

	/** 自平衡
	如果我们删除一个红色节点（此时该节点的儿子将都为叶子节点），
	它的父亲和儿子一定是黑色的。所以我们可以简单的用它的黑色儿子替换它，
	并不会破坏性质3和性质4。通过被删除节点的所有路径只是少了一个红色节点，
	这样可以继续保证性质5。另一种简单情况是在被删除节点是黑色而它的儿子是红色的时候。
	如果只是去除这个黑色节点，用它的红色儿子顶替上来的话，会破坏性质5，
	但是如果我们重绘它的儿子为黑色，则曾经通过它的所有路径将通过它的黑色儿子，
	这样可以继续保持性质5。
	1. 节点是红色或黑色。
	2. 根是黑色。
	3. 所有叶子都是黑色（叶子是NIL节点）。
	4. 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
	5. 从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
	*/

	// 删除节点是黑
	if n.red == false {
		// 删除节点 - 子节点为红
		if child != nil && child.red == true {
			// 不违反规则
			child.red = false
		} else {
			/**
			需要进一步讨论的是在要删除的节点和它的儿子二者都是黑色的时候，
			这是一种复杂的情况（这种情况下该结点的两个儿子都是叶子结点，
			否则若其中一个儿子是黑色非叶子结点，另一个儿子是叶子结点，
			那么从该结点通过非叶子结点儿子的路径上的黑色结点数最小为2，
			而从该结点到另一个叶子结点儿子的路径上的黑色结点数为1，违反了性质5）
			*/
			t.removeBlance(child)
		}
	}
	n = nil
}

func (t *RBTree) removeBlance(n *TreeNode) {
	// 前提 n 是 黑色

	if n == nil {
		return
	}

	if n.parent == nil {
		// n是新 root
		n.red = false
		t.root = n
		return
	}

	// 如果n的兄弟是红色
	// n代替了原n父节点的位置, 所以n.sibling是n的原叔父节点
	if n.sibling().red == true {
		//
		//         黑(P)         右旋+          黑(S)
		//        /    \     对调P、S颜色       /    \
		//     黑(n)   红(S)      -->       红(P)   黑(SR)
		//             /  \                /  \
		//         黑(SL) 黑(SR)        黑(n) 黑(SL)

		n.parent.red = true
		n.sibling().red = false
		if n == n.parent.left {
			t.leftRotate(n.sibling())
		} else {
			t.rightRotate(n.sibling())
		}
		// 这里做了一次 旋转, 但是根本上n的一侧还是删除了一个黑色节点, 违背了规则5
		// 所以还要有后续处理
	}

	if n.red == false && n.sibling().red == false &&
		(n.sibling().left == nil || n.sibling().left.red == false) &&
		(n.sibling().right == nil || n.sibling().right.red == false) {
		//
		//         P(黑)                    P(黑)
		//        /    \                   /    \
		//     n(黑)   S(黑)    -->     n(黑)   S(红)
		//             /  \                     /  \
		//         SL(黑) SR(黑)             SL(黑) SR(黑)

		/**
		N的父亲、S和S的儿子都是黑色的。
		在这种情形下，我们简单的重绘S为红色。
		结果是通过S的所有路径，它们就是以前不通过N的那些路径，都少了一个黑色节点。
		因为删除N的初始的父亲使通过N的所有路径少了一个黑色节点，这使事情都平衡了起来。
		但是，通过P的所有路径现在比不通过P的路径少了一个黑色节点，所以仍然违反性质5。
		要修正这个问题，我们要从头开始，在P上做重新平衡处理。
		*/
		n.sibling().red = true
		t.removeBlance(n.parent)
	} else if n.red == true && n.sibling().red == false &&
		(n.sibling().left == nil || n.sibling().left.red == false) &&
		(n.sibling().right == nil || n.sibling().right.red == false) {
		//
		//         P(红)                    P(黑)
		//        /    \                   /    \
		//     n(黑)   S(黑)    -->     n(黑)   S(红)
		//             /  \                     /  \
		//         SL(黑) SR(黑)             SL(黑) SR(黑)
		/**
		S和S的儿子都是黑色，但是N的父亲是红色。
		在这种情形下，我们简单的交换N的兄弟和父亲的颜色。
		这不影响不通过N的路径的黑色节点的数目，
		但是它在通过N的路径上对黑色节点数目增加了一，
		添补了在这些路径上删除的黑色节点。
		*/
		n.red = false
		n.sibling().red = true
	} else {
		if n.sibling().red == false {
			if n == n.parent.left &&
				n.sibling().left != nil && n.sibling().left.red == true &&
				(n.sibling().right == nil || n.sibling().right.red == false) {
				//          P                        P
				//        /    \                   /    \
				//     n(黑)   S(黑)    -->     n(黑)   SL(黑)
				//             /  \                        \
				//         SL(红) SR(黑)                   S(红)
				//											  \
				//                                            SR(黑)
				n.sibling().red = true
				n.sibling().left.red = false
				t.rightRotate(n.sibling().left)
			} else if n == n.parent.right &&
				(n.sibling().left == nil || n.sibling().left.red == false) &&
				n.sibling().right != nil && n.sibling().right.red == true {
				//           P                       P
				//        /    \                   /    \
				//      S(黑)  n(黑)    -->      SR(黑)  n(黑)
				//      /  \                    /
				//   SL(黑) SR(红)            S(红)
				//							  /
				//						   SL(黑)
				n.sibling().red = true
				n.sibling().right.red = false
				t.leftRotate(n.sibling().right)
			}
		}
		n.sibling().red = n.parent.red
		n.parent.red = false
		if n == n.parent.left {
			n.sibling().right.red = false
			t.leftRotate(n.sibling())
			//			 P(黑)      			      SL()
			//			/    \       		    	/    \
			//		 n(黑)    SL()     左旋(SL)    P(黑)   S(黑)
			//					\       -->		  /		   \
			//				    S(红)       	n(黑)		SR(黑)
			//						\
			//					  SR(黑)
		} else {
			n.sibling().left.red = false
			t.rightRotate(n.sibling())
			//			 P(黑)      			      SL()
			//			/    \       		    	/    \
			//		  SL()   n(黑)    右旋(SL)    P(黑)   S(黑)
			//		  /              	-->		  /		   \
			//	    S(红)       				n(黑)		SR(黑)
			//		/
			//	 SR(黑)
		}
	}

}
