import "sort"

// https://leetcode.com/problems/design-search-autocomplete-system
/*
Design a search autocomplete system for a search engine.
Users may input a sentence (at least one word and end with a special character '#').
For each character they type except '#',
you need to return the top 3 historical hot sentences that have prefix the same as the part of sentence already typed. Here are the specific rules:

The hot degree for a sentence is defined as the number of times a user typed the exactly same sentence before.
The returned top 3 hot sentences should be sorted by hot degree (The first is the hottest one). If several sentences have the same degree of hot, you need to use ASCII-code order (smaller one appears first).
If less than 3 hot sentences exist, then just return as many as you can.
When the input is a special character, it means the sentence ends, and in this case, you need to return an empty list.
Your job is to implement the following functions:

The constructor function:

AutocompleteSystem(String[] sentences, int[] times): This is the constructor. The input is historical data. Sentences is a string array consists of previously typed sentences. Times is the corresponding times a sentence has been typed. Your system should record these historical data.

Now, the user wants to input a new sentence. The following function will provide the next character the user types:

List<String> input(char c): The input c is the next character typed by the user. The character will only be lower-case letters ('a' to 'z'), blank space (' ') or a special character ('#'). Also, the previously typed sentence should be recorded in your system. The output will be the top 3 historical hot sentences that have prefix the same as the part of sentence already typed.


Example:
Operation: AutocompleteSystem(["i love you", "island","ironman", "i love leetcode"], [5,3,2,2])
The system have already tracked down the following sentences and their corresponding times:
"i love you" : 5 times
"island" : 3 times
"ironman" : 2 times
"i love leetcode" : 2 times
Now, the user begins another search:

Operation: input('i')
Output: ["i love you", "island","i love leetcode"]
Explanation:
There are four sentences that have prefix "i". Among them, "ironman" and "i love leetcode" have same hot degree. Since ' ' has ASCII code 32 and 'r' has ASCII code 114, "i love leetcode" should be in front of "ironman". Also we only need to output top 3 hot sentences, so "ironman" will be ignored.

Operation: input(' ')
Output: ["i love you","i love leetcode"]
Explanation:
There are only two sentences that have prefix "i ".

Operation: input('a')
Output: []
Explanation:
There are no sentences that have prefix "i a".

Operation: input('#')
Output: []
Explanation:
The user finished the input, the sentence "i a" should be saved as a historical sentence in system. And the following input will be counted as a new search.


Note:

The input sentence will always start with a letter and end with '#', and only one blank space will exist between two words.
The number of complete sentences that to be searched won't exceed 100. The length of each sentence including those in the historical data won't exceed 100.
Please use double-quote instead of single-quote when you write test cases even for a character input.
Please remember to RESET your class variables declared in class AutocompleteSystem, as static/class variables are persisted across multiple test cases. Please see here for more details.
*/
type Trie struct {
	Cnt      int
	IsWord   bool
	Children map[byte]*Trie
}

type AutocompleteSystem struct {
	UserTrie  *Trie
	UserInput string
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	userTrie := &Trie{0, false, make(map[byte]*Trie)}
	for i := 0; i < len(sentences); i++ {
		sentence := sentences[i]
		insert(userTrie, sentence, times[i])
	}
	return AutocompleteSystem{userTrie, ""}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	// add sentence
	if c == '#' {
		insert(this.UserTrie, this.UserInput, 0)
		this.UserInput = ""
		return []string{}
	}
	// search
	this.UserInput += string(c)
	return this.search()
}

// helpers
func insert(userTrie *Trie, sentence string, time int) {
	current := userTrie
	for i := 0; i < len(sentence); i++ {
		charactor := sentence[i]
		var temp *Trie
		if value, existed := current.Children[charactor]; existed {
			temp = value
		} else {
			temp = &Trie{0, false, make(map[byte]*Trie)}
		}
		if i == len(sentence)-1 {
			temp.IsWord = true
			if time == 0 {
				temp.Cnt++
			} else {
				temp.Cnt = time
			}
		}
		current.Children[charactor] = temp
		current = temp
	}
}

func (this *AutocompleteSystem) search() []string {
	// find the target node
	query := this.UserInput
	targetNode := this.UserTrie
	for i := 0; i < len(query); i++ {
		charactor := query[i]
		if value, existed := targetNode.Children[charactor]; existed {
			targetNode = value
		} else {
			// it means it cant find the target
			return []string{}
		}
	}
	// iterate the target node children
	var results []Result
	// dfs
	var stack []*Stack
	stack = append(stack, &Stack{"", targetNode})
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop.trie.IsWord {
			temp := Result{query + pop.prefix, pop.trie.Cnt}
			results = append(results, temp)
		}
		// add children into stack
		for key, value := range pop.trie.Children {
			stack = append(stack, &Stack{pop.prefix + string(key), value})
		}
	}
	// sort
	descendingOccurence := func(r1, r2 *Result) bool {
		return r1.Occurence > r2.Occurence
	}
	acendingAlphabet := func(r1, r2 *Result) bool {
		return r1.Str < r2.Str
	}
	OrderedBy(descendingOccurence, acendingAlphabet).Sort(results)
	var sortedResults []string
	for i := 0; i < len(results) && i < 3; i++ {
		sortedResults = append(sortedResults, results[i].Str)
	}
	return sortedResults
}

// for the dfs
type Stack struct {
	prefix string
	trie   *Trie
}

// for the results
type Result struct {
	Str       string
	Occurence int
}

/*
sort by multiple keys...so cumbrous
*/
type lessFunc func(p1, p2 *Result) bool
type multiSorter struct {
	results []Result
	less    []lessFunc
}

func (ms *multiSorter) Sort(results []Result) {
	ms.results = results
	sort.Sort(ms)
}
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// implement the interface for sorting
func (ms *multiSorter) Len() int {
	return len(ms.results)
}
func (ms *multiSorter) Swap(i, j int) {
	ms.results[i], ms.results[j] = ms.results[j], ms.results[i]
}
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.results[i], &ms.results[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */
