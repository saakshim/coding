package main

const ALPHABETS = 26


type TrieNode struct {
    
    children    [ALPHABETS]*TrieNode
    isEnd       bool
    keyWord     string
    Count       int
}

type Trie struct {
    
    root    *TrieNode
}

func initializeTrie() *Trie {
    
    return &Trie{
        root:&TrieNode{},
    }
}

func (t *Trie) insert(key string) {
    
    length := len(key)
    node := t.root
    var index int
    
    for i:=0; i < length; i++ {
        
        index = int(key[i] - 'a')
        if node.children[index] == nil {
            node.children[index] = &TrieNode{}
        }
        node = node.children[index]
    }
    
    node.keyWord = key
    node.Count += 1
    node.isEnd = true
    
}
