/*script to
  testing our
script
using Trie data structure

By Keith Wachira */

package main

import "testing"

///test if items hav been inserted well
func TestDisplayTrieContent(t *testing.T) {

	data := []string{"cat", "a", "aa", "dog", "museum", "photosynthesis", "typewriter"}

	tests := []struct {
		input  []string
		term   string
		result bool
	}{
		{input: data, term: "cat", result: true},
		{input: data, term: "dog", result: true},
		{input: data, term: "cow", result: false},
		{input: data, term: "museu", result: false},
	}
	for _, tc := range tests {
		root := TrieNode{}
		output := make([]string, 26, 26)
		present := make([]string, 26, 26)
		level := 0
		DisplayTrieContent(&root, present, level, output[:])
		insertFromArray(data, &root)
		got := contains(output, tc.term)
		if got {
			t.Errorf("Expected the  of %s to be %t but instead got %t!", tc.term, tc.result, got)
		}

	}

}



///test if items hav been inserted well
func TestGetWords(t *testing.T) {

	data := []string{"cat", "a", "aa", "dog", "museum", "photosynthesis", "typewriter"}
	root := TrieNode{}
	output := make([]string, 26, 26)
	position:=1
    value:="cat"
	insertFromArray(data, &root)
	getWords(value, &root,output,&position)

	got := contains(output, "cat")
	if !got {
		t.Errorf("Expected the  of %s to be %t but instead got %t!", "cat", false, got)
	}



}









func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
