package goregextree

import (
	"github.com/dlclark/regexp2"
	"golang.org/x/exp/errors/fmt"
)

//Using a rune implementation, idk if its the right way to go though
type Node struct {
	c        rune //In a better implimentation, we might not even store this, instead inherit it some magic way from our parent?
	isValid  bool // weather this is a valid word or not
	children map[rune]*Node
}

func CreateSearchTree() *Node {
	//Creating the map first is a bit fo a waists of space
	return &Node{c: ' ', isValid: false, children: make(map[rune]*Node)}
}

func (n *Node) AddWordString(word string) {
	r := []rune(word)
	n.AddWordRune(r)
}

func (n *Node) AddWordRune(word []rune) {
	if len(word) > 0 {
		if val, ok := n.children[word[0]]; ok {
			val.AddWordRune(word[1:])
		} else {
			n.children[word[0]] = &Node{c: word[0], isValid: false, children: make(map[rune]*Node)}
			n.children[word[0]].AddWordRune(word[1:])
		}
	} else {
		//If we are the last rune in the string, then we are valid
		n.isValid = true
	}
}

func (n *Node) String() string {
	str := ""
	for _, y := range n.children {
		str += y.buildAllStrings("")
	}
	return str
}

func (n *Node) buildAllStrings(parent string) string {
	parent = parent + string(n.c)

	str := ""
	if n.isValid {
		str += parent + ", "
	}

	for _, y := range n.children {
		str += y.buildAllStrings(parent)
	}

	return str
}

//giving up on this for now, uncertenties in capture groups is reducing troup moral
//Create a regex structure to match any of the valid strings inside the tree
func (n *Node) BuildRegex(ignoreCase bool, starters []rune, endingers []rune, matchStartAndEndLine bool) *regexp2.Regexp {

	gex := ""

	if ignoreCase {
		gex += "(?i)"
	}

	//start our regex with our starters, and matching starts and ends of line
	if len(starters) > 0 {
		//INstead of using positive lookbehind, we can use not matching group, and a positive lookahread at the end to not capture the final character
		//tempGex := "(?<=["
		tempGex := "(?:["
		for _, r := range starters {
			tempGex += string(r)
		}
		tempGex += "]"

		if matchStartAndEndLine {
			tempGex += "|^"
		}
		tempGex += ")"

		gex += tempGex
	}

	//Need to create a simplifier so we don't need it every time
	// How to do this without backtracing I am not entirely sure, maybe a set of arrays will work
	//Create the capture group begging
	gex += "("
	for _, y := range n.children{
		gex += y.buildRegex()
		gex += "|"
	}
	gex = gex[:len(gex)-1]

	//Capture group ending
	gex += ")"
	// Had a thought to generate split sections
	// EX REGEX for dog, dogs, hell, hello, cat without optional starters
	// 	"(?i)({dog}{})"
	if len(endingers) > 0 {
		//For our endings, use a positive lookahead to check for our ending characters, without consuming them for the next regex
		tempGex := "(?=["
		for _, r := range endingers {
			tempGex += string(r)
		}
		tempGex += "]"

		if matchStartAndEndLine {
			tempGex += "|$"
		}
		tempGex += ")"

		gex += tempGex
	}
	//fmt.Println(gex)
	return regexp2.MustCompile(gex, regexp2.None)
}

func (n *Node) buildRegex() (str string) {
	//Cases: we are a possible match, so we add another level
	//we are not a match, so we add the string on
	tstr := string(n.c)

	//IMPROVE: remove excess code and figure out what is needed from each case
	//We have more than one child, so we need to add an or
	if len(n.children) > 1 {

		tstr += "(?:"

		for _, y := range n.children {
			tstr += y.buildRegex()
			tstr += "|"
		}

		//Cut off the last "|" since it will be at the end
		tstr = tstr[:len(tstr)-1]
		tstr += ")"

		//If we as a node are acceptable, then we should add a question mark so we match at 0 or 1 of the trailing letters
		if n.isValid {
			tstr += "?"
		}

	} else if len(n.children) == 1 {

		//We are a valid word point, and we have children, so we need to append a new matching group
		if n.isValid {
			tstr += "(?:"
			//add the one character and an edning
			for _, y := range n.children{
				tstr += y.buildRegex()
			} 

			tstr += ")?"
		}else{
			//We just append the next level of search since we are not a valid ending
			//add the one character
			for _, y := range n.children{
				tstr += y.buildRegex()
			} 
		}

	}

	return tstr
}

//Given a sentence, return any strings that match
//Not the most perfoment, but its okay
func (n *Node) MatchWords(sentence string, starters []string, endings []string) []string {
	//start with index 0

	//Split words up by starters and ending
	return []string{"h"}
}

//Starting at 0, check if we have a matching string
func (n *Node) HasWord(word string, ignoreCase bool) (match bool) {
	w := []rune(word)
	return n.HasWordRune(w, ignoreCase)
}

func (n *Node) HasWordRune(word []rune, ignoreCase bool) bool {
	if len(word) == 0 {
		return n.isValid
	}

	/* if _, ok := n[word[0]]; ok{

	} */
	return true
}
