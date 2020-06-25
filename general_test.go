package stringsearchtree

import (
	"fmt"
	"testing"
)

func TestStringSearch(t *testing.T) {
	n := CreateSearchTree()

	n.AddWordString("hello")
	n.AddWordString("hell")
	n.AddWordString("dog")
	n.AddWordString("dogs")

	fmt.Println(n)
}

func TestRegexMaker(t *testing.T) {
	n := CreateSearchTree()

	n.AddWordString("fsly")
	n.AddWordString("pqq")
	n.AddWordString("pt")
	n.AddWordString("app")
	n.AddWordString("worl")

	r := n.StartBuildRegex(true, []rune{' ', '$'}, []rune{',', ' ', '.'}, true)
	m, err := r.FindStringMatch("fsly is the best stock to trade compared to $pqq, but I would consider $pt as good as well or app")
	if err != nil {
		panic(err)
	}
	
	for m != nil && err == nil{
		fmt.Println(m.Groups()[1].String())
		m, err = r.FindNextMatch(m)
	}

	/* for m, err = r.FindNextMatch(m); err == nil && m != nil; {
		fmt.Println(m.String())
	} */

}
