# https://leetcode.com/problems/word-ladder/description/
# @param {String} begin_word
# @param {String} end_word
# @param {String[]} word_list
# @return {Integer}

# "hot"
# "dog"
# ["hot","dog","dot"]

# https://leetcode.com/problems/word-ladder/description/
# @param {String} begin_word
# @param {String} end_word
# @param {String[]} word_list
# @return {Integer}
def ladder_length(begin_word, end_word, word_list)
	return 0 if !word_list.include?(end_word)
	return 2 if begin_word.length == 1
	queue = [begin_word]
	visited = {}
	visited[begin_word] = true
	result = []
	distance = 0
	# distance = word_list.include?(begin_word) ? 1 : 0
	while(!queue.empty?)
    word = queue.pop
    path = [word]
		# queue.pop
		i = 0
		while i < word.length
				('a'..'z').each do |c|
            tmp = word.clone
						tmp[i] = c
						if tmp == end_word && word_list.include?(tmp)
							# p tmp
							# p end_word
							# p queue
							# p visited
							distance += 1
              # path << tmp
              # p path
              # result << path
							return distance
							# result << distance
							# distance = 0
						end

						if word_list.include?(tmp) && visited[tmp].nil?
							# p tmp

							# p visited
							# p distance

							queue << tmp
							visited[tmp] = true
							distance += 1
              # path << tmp
              # next
              # p queue
						end
				end
				i+=1
		end
		# p queue
		# p visited
		# p distance
	end
	# p result
	# return result.empty? ? 0 : result.min
  return distance
end


# bool BFS(Node& Vs, Node& Vd){
#     queue<Node> Q;
#     Node Vn, Vw;
#     int i;

#     //初始状态将起点放进队列Q
#     Q.push(Vs);
#     hash(Vw) = true;//设置节点已经访问过了！

#     while (!Q.empty()){//队列不为空，继续搜索！
#         //取出队列的头Vn
#         Vn = Q.front();

#         //从队列中移除
#         Q.pop();

#         while(Vw = Vn通过某规则能够到达的节点){
#             if (Vw == Vd){//找到终点了！
#                 //把路径记录，这里没给出解法
#                 return true;//返回
#             }

#             if (isValid(Vw) && !visit[Vw]){
#                 //Vw是一个合法的节点并且为白色节点
#                 Q.push(Vw);//加入队列Q
#                 hash(Vw) = true;//设置节点颜色
#             }
#         }
#     }
#     return false;//无解
# }
