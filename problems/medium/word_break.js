Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false

https://leetcode.com/problems/word-break/description/

/**
 * @param {string} s
 * @param {string[]} wordDict
 * @return {boolean}
 */
var wordBreak = function(s, wordDict) {
    if (wordDict.length === 0) return false;
    if (wordDict.length === 1) return s === wordDict[0];
    
    let queue = [''];
    let memo = new Map();
    
    while (queue.length > 0) {
        const val = queue.shift();
        
        for (let word of wordDict) {
            const searchWord = `${val}${word}`;
            const startsWith = s.indexOf(searchWord) === 0;
            if (searchWord === s) return true;
            else if (!memo.has(searchWord) && startsWith) {
                memo.set(searchWord, true);
                queue.push(searchWord);
            }
        }
    }
    
    return false;
};
