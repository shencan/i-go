package tire

import (
	"fmt"
)

const MaxCap = 26 // a-z 每一层级分支数

type Tire struct {
	next map[rune]*Tire
	// 标记到此是否为一个完整单词
	isWord bool
}

// NewTire Init Tire
func NewTire() Tire {
	root := new(Tire)
	root.next = make(map[rune]*Tire, MaxCap)
	root.isWord = false
	return *root
}
func (t *Tire) InsertMany(words []string) {
	for _, v := range words {
		t.Insert(v)
	}
}

// Insert 添加字符串到 Tire 中
func (t *Tire) Insert(word string) {
	tire := t
	for _, v := range word {
		if tire.next[v] == nil {
			node := new(Tire)
			//子节点数量为26
			node.next = make(map[rune]*Tire, MaxCap)
			//初始化节点单词标志为假
			node.isWord = false
			tire.next[v] = node
		}
		tire = tire.next[v]
	}
	tire.isWord = true
}

/** Returns if the word is in the trie. */
func (t *Tire) Search(word string) bool {
	tire := t
	for _, v := range word {
		if tire.next[v] == nil {
			return false
		}
		tire = tire.next[v]
	}
	return tire.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Tire) StartsWith(prefix string) bool {
	tire := t
	for _, v := range prefix {
		if tire.next[v] == nil {
			return false
		}
		tire = tire.next[v]
	}
	return true
}

func main() {
	t := NewTire()
	t.Insert("Hello")
	fmt.Print(t.Search("Hello"), "\n")
	fmt.Print(t.Search("Hallo"), "\n")
}
