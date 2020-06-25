# Regex Tree Creator

Regex Tree Creator creates a regex search pattern for all the words built into this tree.


## Installation

```bash
go get github.com/alexander-e-andrews/GoRegexTree
```

Note: Uses github.com/dlclark/regexp2 for forward look ahead.

## Usage

```go
import github.com/alexander-e-andrews/GoRegexTree

tree := goregextree.CreateSearchTree()

tree.AddWordString("foo")
tree.AddWordString("bar")
tree.AddWordString("foobar")
tree.AddWordString("bars")

re := tree.BuildRegex(true, []rune{' '}, []rune{',', ' ', '.', ';'}, true)
//re is a regexp2 for the string (?i)(?:[ ]|^)(foo(?:bar)?|bar)(?=[, .;]|$)
```