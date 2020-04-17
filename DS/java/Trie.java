public class Trie {
    
    static final int ALPHABETS = 26;
    
    static class TrieNode {
        
        TrieNode[] children = new TrieNode[ALPHABETS];
        
        boolean isEnd; // leaf node (end of word)
        int count;
        TrieNode() {
            isEnd = false;
            count = 0;
            for (int i=0; i < ALPHABETS; i++) {
                children[i] = null;
            }
        }
    }
    
    static TrieNode root;
    
    public static void insert(String key) {
        
        int l = key.length();
        if (l == 0)
            return;
            
        int index;
        
        TrieNode pCrawl = root;
        
        for (int i=0; i < l; i++) {
            
            index = key.charAt(i) - 'a';
            if (pCrawl.children[index] == null) {
                pCrawl.children[index] = new TrieNode();
            }
            
            pCrawl = pCrawl.children[index];
        }
        
        // inserted each alphabet in the word; 
        // set end here; increment count
        pCrawl.isEnd = true;
        pCrawl.count += 1;
    }
    
    public static boolean search(String key) {
        
        int l = key.length();
        TrieNode pCrawl = root;
        int index;
        
        for (int i=0; i < l; i++) {
            
            index = key.charAt(i) - 'a';
            if (pCrawl.children[index] == null) {
                return false;
            }
            
            pCrawl = pCrawl.children[index];
        }
        
        return (pCrawl != null && pCrawl.isEnd == true);
    }
       
    // Driver 
    public static void main(String args[]) 
    { 
        // Input keys (use only 'a' through 'z' and lower case) 
        String keys[] = {"the", "a", "the", "answer", "any", 
                         "bye", "bye", "their"}; 
       
        String output[] = {"Not present in trie", "Present in trie"}; 
       
       
        root = new TrieNode(); 
       
        // Construct trie 
        int i; 
        for (i = 0; i < keys.length ; i++) 
            insert(keys[i]); 
       
        // Search for different keys 
        if(search("the") == true) 
            System.out.println("the --- " + output[1]); 
        else System.out.println("the --- " + output[0]); 
          
        if(search("these") == true) 
            System.out.println("these --- " + output[1]); 
        else System.out.println("these --- " + output[0]); 
          
        if(search("bye") == true) 
            System.out.println("bye --- " + output[1]); 
        else System.out.println("bye --- " + output[0]); 
          
        if(search("thaw") == true) 
            System.out.println("thaw --- " + output[1]); 
        else System.out.println("thaw --- " + output[0]); 
        
    } 
}