# @param {Integer} num_courses
# @param {Integer[][]} prerequisites
# @return {Boolean}
def can_finish(num_courses, prerequisites)
  edges = {}
  indgree = Array.new(num_courses)
  (0..prerequisites.length-1).each do |i|
    edges[prerequisites[i][1]] = [] if edges[prerequisites[i][1]].nil?
    edges[prerequisites[i][1]].push(prerequisites[i][0])
    indgree[prerequisites[i][0]] = 0 if indgree[prerequisites[i][0]].nil?
    indgree[prerequisites[i][0]] += 1
  end
  queue = []
  (0..num_courses-1).each do |i|
    queue.push(i) if indgree[i].to_i.zero?
  end
  num = 0
  p indgree
  p queue
  p edges
  until queue.empty?
    n = queue.pop
    num += 1
    if !edges[n].nil?
      edges[n].each do |i|
        indgree[i] -= 1 if !indgree[i].nil? && indgree[i].to_i > 0
        queue.push(i) if indgree[i].to_i.zero?
      end
    end
  end
  return true if num == num_courses
  false
end