# https://leetcode.com/problems/word-ladder/description/
# @param {String} begin_word
# @param {String} end_word
# @param {String[]} word_list
# @return {Integer}

# "hit"
# "cog"
# ["hot","dot","dog","lot","log","cog"]

# "hot"
# "dog"
# ["hot","dog","dot"]

# "hot"
# "dog"
# ["hot","dog"]

# "hot"
# "dog"
# ["hot","cog","dog","tot","hog","hop","pot","dot"]

# "cet"
# "ism"
# ["kid","tag","pup","ail","tun","woo","erg","luz","brr","gay","sip","kay","per","val","mes","ohs","now","boa","cet","pal","bar","die","war","hay","eco","pub","lob","rue","fry","lit","rex","jan","cot","bid","ali","pay","col","gum","ger","row","won","dan","rum","fad","tut","sag","yip","sui","ark","has","zip","fez","own","ump","dis","ads","max","jaw","out","btu","ana","gap","cry","led","abe","box","ore","pig","fie","toy","fat","cal","lie","noh","sew","ono","tam","flu","mgm","ply","awe","pry","tit","tie","yet","too","tax","jim","san","pan","map","ski","ova","wed","non","wac","nut","why","bye","lye","oct","old","fin","feb","chi","sap","owl","log","tod","dot","bow","fob","for","joe","ivy","fan","age","fax","hip","jib","mel","hus","sob","ifs","tab","ara","dab","jag","jar","arm","lot","tom","sax","tex","yum","pei","wen","wry","ire","irk","far","mew","wit","doe","gas","rte","ian","pot","ask","wag","hag","amy","nag","ron","soy","gin","don","tug","fay","vic","boo","nam","ave","buy","sop","but","orb","fen","paw","his","sub","bob","yea","oft","inn","rod","yam","pew","web","hod","hun","gyp","wei","wis","rob","gad","pie","mon","dog","bib","rub","ere","dig","era","cat","fox","bee","mod","day","apr","vie","nev","jam","pam","new","aye","ani","and","ibm","yap","can","pyx","tar","kin","fog","hum","pip","cup","dye","lyx","jog","nun","par","wan","fey","bus","oak","bad","ats","set","qom","vat","eat","pus","rev","axe","ion","six","ila","lao","mom","mas","pro","few","opt","poe","art","ash","oar","cap","lop","may","shy","rid","bat","sum","rim","fee","bmw","sky","maj","hue","thy","ava","rap","den","fla","auk","cox","ibo","hey","saw","vim","sec","ltd","you","its","tat","dew","eva","tog","ram","let","see","zit","maw","nix","ate","gig","rep","owe","ind","hog","eve","sam","zoo","any","dow","cod","bed","vet","ham","sis","hex","via","fir","nod","mao","aug","mum","hoe","bah","hal","keg","hew","zed","tow","gog","ass","dem","who","bet","gos","son","ear","spy","kit","boy","due","sen","oaf","mix","hep","fur","ada","bin","nil","mia","ewe","hit","fix","sad","rib","eye","hop","haw","wax","mid","tad","ken","wad","rye","pap","bog","gut","ito","woe","our","ado","sin","mad","ray","hon","roy","dip","hen","iva","lug","asp","hui","yak","bay","poi","yep","bun","try","lad","elm","nat","wyo","gym","dug","toe","dee","wig","sly","rip","geo","cog","pas","zen","odd","nan","lay","pod","fit","hem","joy","bum","rio","yon","dec","leg","put","sue","dim","pet","yaw","nub","bit","bur","sid","sun","oil","red","doc","moe","caw","eel","dix","cub","end","gem","off","yew","hug","pop","tub","sgt","lid","pun","ton","sol","din","yup","jab","pea","bug","gag","mil","jig","hub","low","did","tin","get","gte","sox","lei","mig","fig","lon","use","ban","flo","nov","jut","bag","mir","sty","lap","two","ins","con","ant","net","tux","ode","stu","mug","cad","nap","gun","fop","tot","sow","sal","sic","ted","wot","del","imp","cob","way","ann","tan","mci","job","wet","ism","err","him","all","pad","hah","hie","aim","ike","jed","ego","mac","baa","min","com","ill","was","cab","ago","ina","big","ilk","gal","tap","duh","ola","ran","lab","top","gob","hot","ora","tia","kip","han","met","hut","she","sac","fed","goo","tee","ell","not","act","gil","rut","ala","ape","rig","cid","god","duo","lin","aid","gel","awl","lag","elf","liz","ref","aha","fib","oho","tho","her","nor","ace","adz","fun","ned","coo","win","tao","coy","van","man","pit","guy","foe","hid","mai","sup","jay","hob","mow","jot","are","pol","arc","lax","aft","alb","len","air","pug","pox","vow","got","meg","zoe","amp","ale","bud","gee","pin","dun","pat","ten","mob"]


# https://leetcode.com/problems/word-ladder/description/
# @param {String} begin_word
# @param {String} end_word
# @param {String[]} word_list
# @return {Integer}
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
	distance = 1
	while(!queue.empty?)
    current_queue = queue.clone
    q_length = current_queue.length
    queue = []
    (0..q_length - 1).each do |q|
      word = current_queue.pop
      length = word.length
      (0..length-1).each do |i|
          ('a'..'z').each do |c|
              tmp = word.clone
              tmp[i] = c
              if tmp == end_word && word_list.include?(tmp)
                distance += 1
                return distance
              end
              if word_list.include?(tmp) && visited[tmp] != true && tmp != word
                queue.push(tmp)
                visited[word] = true
              end
          end
      end
    end
    distance += 1 if !queue.empty?
	end
  return distance == 1 ? 0 : distance
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
