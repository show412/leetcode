# https://leetcode.com/problems/word-ladder/description/
def ladder_length(begin_word, end_word, word_list)
    visited = [begin_word]
    distance = 1
    while !visited.include?(end_word)
        to_add = []
        visited.each do |word|
            i = 0
            while i<word.length
                ('a'..'z').each do |c|
                    tmp = word.dup
                    tmp[i] = c
                    if word_list.include?(tmp)
                        to_add << tmp
                        word_list.delete(tmp)
                    end
                end
                i+=1
            end
        end
        distance +=1
        return 0 if to_add.empty?
        visited = to_add
    end
    return distance
end
