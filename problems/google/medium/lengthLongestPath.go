import (
	"strings"
)

// https://leetcode.com/problems/longest-absolute-file-path/
/*
Suppose we abstract our file system by a string in the following manner:

The string "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext" represents:

dir
    subdir1
    subdir2
        file.ext
The directory dir contains an empty sub-directory subdir1 and a sub-directory subdir2 containing a file file.ext.

The string "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext" represents:

dir
    subdir1
        file1.ext
        subsubdir1
    subdir2
        subsubdir2
            file2.ext
The directory dir contains two sub-directories subdir1 and subdir2. subdir1 contains a file file1.ext and an empty second-level sub-directory subsubdir1. subdir2 contains a second-level sub-directory subsubdir2 containing a file file2.ext.

We are interested in finding the longest (number of characters) absolute path to a file within our file system. For example, in the second example above, the longest absolute path is "dir/subdir2/subsubdir2/file2.ext", and its length is 32 (not including the double quotes).

Given a string representing the file system in the above format, return the length of the longest absolute path to file in the abstracted file system. If there is no file in the system, return 0.

Note:
The name of a file contains at least a . and an extension.
The name of a directory or sub-directory will not contain a ..
Time complexity required: O(n) where n is the size of the input string.

Notice that a/aa/aaa/file1.txt is not the longest file path, if there is another path aaaaaaaaaaaaaaaaaaaaa/sth.png.
*/
/*
这个的思路是对的 split 再indexLastByte 但是没有想到用stack去简化
*/
func lengthLongestPath(input string) int {
	if input == "" {
		return 0
	}
	if strings.IndexByte(input, '.') >= 0 && strings.LastIndexByte(input, '\t') <= -1 {
		return len(input)
	}
	inputArray := strings.Split(input, "\n")
	res := len(inputArray[0])

	for start := 1; start < len(inputArray); start++ {
		cur := inputArray[start]

		end := start + 1
		curIndex := strings.LastIndexByte(cur, '\t')
		length := len(cur)
		if curIndex >= 0 {
			length = len(cur[curIndex:])
		}

		if start == 1 && strings.LastIndexByte(cur, '.') >= 0 {
			res = max(res, res+length)
			// fmt.Println("cur", res)
			continue
		}
		if end < len(inputArray) && cur[0] == '\t' && cur[1] != '\t' {
			// fmt.Println("********")
			// // fmt.Println(start)
			// fmt.Println(cur)
			// fmt.Println(len(cur[curIndex:]))

			for end < len(inputArray) && (end-start) < len(inputArray[end]) && inputArray[end][end-start] == '\t' {

				next := inputArray[end]
				// fmt.Println(next)

				nextIndex := strings.LastIndexByte(next, '\t')
				if nextIndex >= 0 {
					length += len(next[nextIndex:])
				} else {
					length += len(next)
				}

				// fmt.Println(len(next[nextIndex:]))
				if strings.IndexByte(next, '.') >= 0 {
					res = max(res, len(inputArray[0])+length)
					// fmt.Println("next", res)
					break
				}
				end++
			}
		} else {
			start = end
			continue
		}
	}
	if res == len(inputArray[0]) {
		return 0
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// stack
/*

public int lengthLongestPath(String input) {
        Deque<Integer> stack = new ArrayDeque<>();
        stack.push(0); // "dummy" length
        int maxLen = 0;
        for(String s:input.split("\n")){
            int lev = s.lastIndexOf("\t")+1; // number of "\t"
            while(lev+1<stack.size()) stack.pop(); // find parent
            int len = stack.peek()+s.length()-lev+1; // remove "/t", add"/"
            stack.push(len);
            // check if it is file
            if(s.contains(".")) maxLen = Math.max(maxLen, len-1);
        }
        return maxLen;
    }

An even shorter and faster solution using array instead of stack:

public int lengthLongestPath(String input) {
    String[] paths = input.split("\n");
    int[] stack = new int[paths.length+1];
    int maxLen = 0;
    for(String s:paths){
        int lev = s.lastIndexOf("\t")+1, curLen = stack[lev+1] = stack[lev]+s.length()-lev+1;
        if(s.contains(".")) maxLen = Math.max(maxLen, curLen-1);
    }
    return maxLen;
}
*/
