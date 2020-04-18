package main

import (
    "log"
    "regexp"
    "strings"
)

func main() {
    para := "a ab abc abcd"
    
    operations(para)
}

const ALPHABETS = 26

var maxCt int
var maxWord = ""

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

func (t *Trie) search(key string) bool {

    node := t.root
    length := len(key)
    var index int

    for i := 0; i < length; i++ {
        index = int(key[i] - 'a')
        if node.children[index] == nil {
            return false
        }
        log.Printf("index=%d; node=%+v", index, node)

        node = node.children[index]
    }

    if node != nil && node.isEnd {
        return true
    }
    return false
}

func (t *Trie) delete(key string) bool {

    node := t.root
    length := len(key)
    prevIndex := int(key[0]-'a')
    var index int
    prevLeaf := t.root
    var i int

    for i = 0; i < length; i++ {

        index = int(key[i] - 'a')
        if node.children[index] == nil {
            return false
        }
        if node.isEnd {
            log.Printf("isEnd; node=%+v", node)
            prevLeaf = node
            prevIndex = index
            log.Printf("i=%d; index=%d; string(key[i])=%s", i, index, string(key[i]))
        }
        node = node.children[index]
    }
 

    if node != nil && node.isEnd {
        log.Printf("node=%+v", node)
        
        node.Count -= 1
        node.isEnd = false
        node.keyWord = ""
        log.Printf("prevIndex=%d", prevIndex)
        
        // check if it is a prefix for other words
        if isPrefix(node) {
            // do not delete any nodes
            log.Printf("is prefix")
            log.Printf("prevLeaf=%+v", prevLeaf)
          
        } else {
            // if it is a leaf, check for previous leaf nodes
            log.Printf("is leaf")
            log.Printf("prevLeaf=%+v", prevLeaf)
            deleteFromIndex(prevLeaf, prevIndex)
            log.Printf("prevLeaf=%+v", prevLeaf)
        }
       
    }
    return true
}


func deleteFromIndex(node *TrieNode, index int) {

    var nextChild *TrieNode

    for true {
     
        nextChild = node.children[index]
        node.children[index] = nil
        
        if nextChild == nil {
            return
        }
        
    }
}

func isPrefix(node *TrieNode) bool {

    for i, _ := range node.children {
        
        if node.children[i] != nil {
            return true
        }
    }
    return false

}


func operations(paragraph string) {
   
    reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

    if len(paragraph) == 0 {
        return
    }
    paragraph = reg.ReplaceAllString(paragraph, " ")
    trie := initializeTrie()
    
    
    paraSplits := strings.Split(paragraph, " ")    
    for _, s := range paraSplits {  
        s = strings.ToLower(s)
        
        trie.insert(s)
             
    }

    wrd := "abc"
    if trie.search(wrd) {
        log.Printf("Found word %s", wrd)
    } else {
        log.Printf("Not Found word %s", wrd)
    }
 


    wrd = "abcd"
    if trie.delete(wrd) {
        log.Printf("Deleted word %s", wrd)
    }

    wrd = "abc"
    if trie.search(wrd) {
        log.Printf("Found word %s", wrd)
    } else {
        log.Printf("Not Found word %s", wrd)
    }

}

func contains(k string, arr []string) bool {
    
    for _, b := range arr {
        if b == k {
            return true
        }
    }
    return false
}


/** OUTPUT:

2009/11/10 23:00:00 index=0; node=&{children:[0xc0000c00f0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:false keyWord: Count:0}
2009/11/10 23:00:00 index=1; node=&{children:[<nil> 0xc0000c01e0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:a Count:1}
2009/11/10 23:00:00 index=2; node=&{children:[<nil> <nil> 0xc0000c02d0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:ab Count:1}
2009/11/10 23:00:00 Found word abc
2009/11/10 23:00:00 isEnd; node=&{children:[<nil> 0xc0000c01e0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:a Count:1}
2009/11/10 23:00:00 i=1; index=1; string(key[i])=b
2009/11/10 23:00:00 isEnd; node=&{children:[<nil> <nil> 0xc0000c02d0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:ab Count:1}
2009/11/10 23:00:00 i=2; index=2; string(key[i])=c
2009/11/10 23:00:00 isEnd; node=&{children:[<nil> <nil> <nil> 0xc0000c03c0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:abc Count:1}
2009/11/10 23:00:00 i=3; index=3; string(key[i])=d
2009/11/10 23:00:00 node=&{children:[<nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:abcd Count:1}
2009/11/10 23:00:00 prevIndex=3
2009/11/10 23:00:00 is leaf
2009/11/10 23:00:00 prevLeaf=&{children:[<nil> <nil> <nil> 0xc0000c03c0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:abc Count:1}
2009/11/10 23:00:00 prevLeaf=&{children:[<nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:abc Count:1}
2009/11/10 23:00:00 Deleted word abcd
2009/11/10 23:00:00 index=0; node=&{children:[0xc0000c00f0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:false keyWord: Count:0}
2009/11/10 23:00:00 index=1; node=&{children:[<nil> 0xc0000c01e0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:a Count:1}
2009/11/10 23:00:00 index=2; node=&{children:[<nil> <nil> 0xc0000c02d0 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>] isEnd:true keyWord:ab Count:1}
2009/11/10 23:00:00 Found word abc

**/
