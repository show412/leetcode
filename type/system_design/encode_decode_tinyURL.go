/*
 * @Author: hongwei.sun
 * @Date: 2024-04-18 20:25:33
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-18 22:29:41
 * @Description: file content
 */
//  https://leetcode.com/problems/encode-and-decode-tinyurl/
/*
Note: This is a companion problem to the System Design problem: Design TinyURL.
TinyURL is a URL shortening service where you enter a URL such as https://leetcode.com/problems/design-tinyurl and it returns a short URL such as http://tinyurl.com/4e9iAk. Design a class to encode a URL and decode a tiny URL.

There is no restriction on how your encode/decode algorithm should work. You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.

Implement the Solution class:

Solution() Initializes the object of the system.
String encode(String longUrl) Returns a tiny URL for the given longUrl.
String decode(String shortUrl) Returns the original long URL for the given shortUrl. It is guaranteed that the given shortUrl was encoded by the same object.
 

Example 1:

Input: url = "https://leetcode.com/problems/design-tinyurl"
Output: "https://leetcode.com/problems/design-tinyurl"

Explanation:
Solution obj = new Solution();
string tiny = obj.encode(url); // returns the encoded tiny url.
string ans = obj.decode(tiny); // returns the original url after decoding it.
 

Constraints:

1 <= url.length <= 104
url is guranteed to be a valid URL.
*/
import (
    "strconv"
)
type Codec struct {
    m map[int]string
	urlID int
}


func Constructor() Codec {
    return Codec {
		m: make(map[int]string),
		urlID: 0,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.urlID++
	this.m[this.urlID] = longUrl
	return "http://tinyurl.com/" + strconv.FormatInt(int64(this.urlID), 16)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    urlid, _ := strconv.Atoi(shortUrl[len("http://tinyurl.com/"):])
	return this.m[urlid]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
