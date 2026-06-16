type Node struct {
	isEnd    bool
	children [26]*Node
}

type PrefixTree struct {
	children [26]*Node
}

func Constructor() PrefixTree {
	return PrefixTree{children: [26]*Node{}}
}

func (this *PrefixTree) Insert(word string) {
	current := &this.children
	for wordIndex, rune := range word {
		letterIndex := int(rune - 'a')
		if (*current)[letterIndex] == nil {
			(*current)[letterIndex] = &Node{children: [26]*Node{}}
		}
		if wordIndex == len(word)-1 {
			(*current)[letterIndex].isEnd = true
		}

		current = &(*current)[letterIndex].children
	}
}

func (this *PrefixTree) Search(word string) bool {
	current := &this.children
	for wordIndex, rune := range word {
		letterIndex := int(rune - 'a')
		if (*current)[letterIndex] == nil {
			return false
		}
		if wordIndex == len(word)-1 {
			return (*current)[letterIndex].isEnd
		}

		current = &(*current)[letterIndex].children
	}
	return false
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	current := &this.children
	for wordIndex, rune := range prefix {
		letterIndex := int(rune - 'a')
		if (*current)[letterIndex] == nil {
			return false
		}
		if wordIndex == len(prefix)-1 {
			return true
		}

		current = &(*current)[letterIndex].children
	}
	return false
}