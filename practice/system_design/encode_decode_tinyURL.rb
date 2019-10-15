# TinyURL is a URL shortening service where you enter a URL such as https://leetcode.com/problems/design-tinyurl and it returns a short URL such as http://tinyurl.com/4e9iAk.

# Design the encode and decode methods for the TinyURL service. There is no restriction on how your encode/decode algorithm should work. You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.
# https://leetcode.com/problems/encode-and-decode-tinyurl/description/

# Encodes a URL to a shortened URL.
#
# @param {string} longUrl
# @return {string}
# 在 solution 里有很多种解法 https://leetcode.com/articles/encode-and-decode-tinyurl/
$map = {}
$i = 0
def encode(longUrl)
  match = longUrl.match(/((http[s]?|ftp):\/\/([0-9a-zA-Z.]*)+\/)(.*)?/)
  # hostUrl = match[1]
  adUrl=match[4]
  $i += 1
  $map[$i] = match[0]
  tinyUrl = "http://tinyurl.com/" + $i.to_s(16)
  return tinyUrl
end

# Decodes a shortened URL to its original URL.
#
# @param {string} shortUrl
# @return {string}
def decode(shortUrl)
  match = shortUrl.match(/(http:\/\/tinyurl.com\/)(.*)?/)
  originUrl = $map[match[2].hex]
  return originUrl
end


# Your functions will be called as such:
# decode(encode(url))
