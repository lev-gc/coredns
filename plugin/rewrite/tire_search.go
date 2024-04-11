package rewrite

type TrieNode struct {
	isInit   bool
	key      bool
	children *[2]TrieNode
	meta     interface{}
}

func newTrieTree() *TrieNode {
	return &TrieNode{true, false, nil, nil}
}

// add - add meta to Trie tree
func (t *TrieNode) add(pattern []bool, metaData interface{}) {
	node := t
	for _, key := range pattern {
		if node.children == nil {
			node.children = &[2]TrieNode{}
		}
		var index int
		if key {
			index = 1
		} else {
			index = 0
		}
		if !node.children[index].isInit {
			node = &node.children[index]
			node.key = key
			node.isInit = true
		} else {
			node = &node.children[index]
		}
	}
	node.meta = metaData
}

// search check if the input matched in the tree
func (t *TrieNode) search(inputBoolArr []bool) (isMatch bool, metaData interface{}) {
	if t == nil || !t.isInit || t.children == nil {
		// trie tree not init
		return false, nil
	}
	node := t
	for _, key := range inputBoolArr {
		var index int
		if key {
			index = 1
		} else {
			index = 0
		}
		if node.children == nil || node.children[index].isInit == false {
			if node.meta != nil {
				return true, node.meta
			} else {
				return false, nil
			}
		}
		node = &node.children[index]
	}
	// if not return in the loop, that it is full match with one record in meta.
	return true, node.meta
}
