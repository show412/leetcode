// IMPORT LIBRARY PACKAGES NEEDED BY YOUR PROGRAM
// SOME FUNCTIONALITY WITHIN A PACKAGE MAY BE RESTRICTED
// DEFINE ANY FUNCTION NEEDED
// FUNCTION SIGNATURE BEGINS, THIS FUNCTION IS REQUIRED
function cellCompete(states, days) {
  // WRITE YOUR CODE HERE
  for (var i = 0; i < days; i++) {
    var pre = states;
    for (var j = 0; j < pre.length; j++) {
      var left = 0;
      var right = 0;
      if (j != 0) {
        left = pre[j-1];
      }
      if (j != pre.length -1 ) {
        right = pre[j+1]
      }
      if (left != right) {
        states[j] = 1
      } else {
        states[j] = 0
      }
    }
  }
  return states
}
// FUNCTION SIGNATURE ENDS
