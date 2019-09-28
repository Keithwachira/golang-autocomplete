/*script to
  search for strings
in a dictionary
using Trie data structure

By Keith Wachira */

package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

const alphaSize = 26

///using an or a file to test
func main() {
	///executionTime(time.Now(),"main_test")
	root := TrieNode{}
	dictionary := []string{"cat", "a", "aa", "dog","o","gdo", "museum", "photosynthesis", "typewriter"}
	////dictionary := ReadFileData("words_alpha.txt")
	insertFromArray(dictionary, &root)

	searchTerms := []string{"cat", "dog"}

	searchMultipleTerms(searchTerms, &root)

	///DisplayTrieContent(&root,present,level,output[:])

}

///all words in the dictionary are  in small letter
///thus can  fit in latin-1 use a slice instead of map for speed
type TrieNode struct {
	children       [alphaSize]*TrieNode
	IsCompleteWord bool ///is it a complete word or just a prefix
}




//return new trie node initialized
// to nil
func getNode() *TrieNode {
	node := &TrieNode{}
	node.IsCompleteWord = false

	for i := 0; i < alphaSize; i++ {
		node.children[i] = nil
	}

	node.IsCompleteWord = false
	return node
}

///we will use this in the insertion of new nodes
///that is we will load all the dictionary
///in the trie node using insert
func insertNode(root *TrieNode, word string) {
	current := root

	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = getNode()
		}
		current = current.children[index]
	}

	current.IsCompleteWord = true

}

///helper function to make sure our function
///is optimal
func executionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

///this will recursively get all possible words in
//the dictionary
///output will help in testing
func SearchWords(root *TrieNode, present []bool, word string, output []string, position *int) {

	if root.IsCompleteWord == true {
		newPosition := *position + 1
		*position = newPosition

		// *position+=*position+1
		if *position >= len(output) {
			output = append(output, word)
		} else {
			output[*position] = word
		}

		log.Println(word)
	}

	for i := 0; i < alphaSize; i++ {
		if present[i] == true && root.children[i] != nil {
			c := string(i + 'a')
			///log.Println("look i am ",c)
			//log.Println("i have become",word+c)
			SearchWords(root.children[i], present, word+c, output, position)

		}

	}

}

///this is for testing
///to make sure our trie has
//inserted right data
func DisplayTrieContent(root *TrieNode, word []string, level int, output []string) {
	if root == nil {
		return
	}

	if root.IsCompleteWord {
		wp := strings.Join(word[:], "")
		output[level] = wp
		//output=append(output,wp)
		///log.Println(level)
	}

	for i := 0; i < alphaSize; i++ {
		if root.children[i] != nil {
			word[level] = string(i + 'a')
			DisplayTrieContent(root.children[i], word, level+1, output)

		}

	}

}














func insertFromArray(terms []string, root *TrieNode) {
	for _, value := range terms {
		insertNode(root, value)

	}

}

////this will open our file with
///dictionary words and read it
///return array instead of inserting
////directly from file to allow testing
func ReadFileData(path string) []string {

	///defer executionTime(time.Now(), "ReadFile")

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Error oppening dictionary file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	/// return each line stripped of any trailing end-of-line
	scanner.Split(bufio.ScanLines)

	// root Node of Trie which need to be empty

	////loop all lines and insert all
	///words into our trie
	var data []string
	for scanner.Scan() {
		data = append(data, strings.TrimSpace(scanner.Text()))
		////insertNode(root, scanner.Text())

	}
	/// getWords("cat",root)
	return data

}

func searchMultipleTerms(terms []string, root *TrieNode) {

	for _, value := range terms {
		position := 1
		output := make([]string, 26, 26)
		getWords(value, root, output, &position)
	}

}

///this will search for
///all words that can be formed from
//a single word
func getWords(term string, root *TrieNode, output []string, position *int) {
	////convert our search tetm to rune
	///char := []rune(term)
	log.Printf("Phrases from %s:", term)
	present := make([]bool, alphaSize, alphaSize)

	///log.Println(len(term))
	for i := 0; i < len(term); i++ {

		present[ term[i]-'a'] = true

	}
	word := ""
	temp := root
	for i := 0; i < alphaSize; i++ {
		if present[i] == true && temp.children[i] != nil {
			word := word + string(i+'a')
			SearchWords(temp.children[i], present, word, output, position)
			word = ""
		}

	}

}

