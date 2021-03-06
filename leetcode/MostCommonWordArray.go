package main

/**
Problem : https://leetcode.com/problems/most-common-word/

Description :

Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.
Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive.  The answer is in lowercase.

Example:

Input: 
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"

*/

import (
	"log"
	"regexp"
	"strings"
)

func main() {
	para := "a."
	banned := []string{""}
	log.Printf("maxWord=%s", mostCommonWord(para, banned))
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

func maxOccuringWord(tp *TrieNode) bool {
    
    if tp == nil {
        return false
    }
    
    for _, child := range tp.children {
        
        if child == nil {
            continue
        }
        if maxCt < child.Count {
            maxCt = child.Count
            maxWord = child.keyWord
            log.Printf("maxWord=%s", maxWord)
        }
        
        //preorder traversal : recurse on child nodes
        maxOccuringWord(child)
    }
    return true
}


func mostCommonWord(paragraph string, banned []string) string {
   
    reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

    if len(paragraph) == 0 {
        return ""
    }
    paragraph = reg.ReplaceAllString(paragraph, " ")
    trie := initializeTrie()
    
    
    paraSplits := strings.Split(paragraph, " ")    
    for _, s := range paraSplits {  
        s = strings.ToLower(s)
        if !contains(s, banned) {
             trie.insert(s)
        }     
    }
    
    maxOccuringWord(trie.root)
    log.Printf("maxWord=%s", maxWord)
    return maxWord
}

func contains(k string, arr []string) bool {
    
    for _, b := range arr {
        if b == k {
            return true
        }
    }
    return false
}