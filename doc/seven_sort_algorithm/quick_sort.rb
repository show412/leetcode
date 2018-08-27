def quick_sort(left, right)
  return if(left > right)
  # set the pivot
  temp = $array[left];
  i = left
  j = right
  # take the item which is bigger than pivot to the right side of pivot
  # take the item which is smaller than pivot to the left side of pivot
  while(i != j)
    while($array[j] >= temp && i < j  )
      j -= 1
    end
    while ($array[i] <= temp && i < j)
      i += 1
    end
    if (i < j)
      t=$array[i]
      $array[i]=$array[j]
      $array[j]=t
    end
  end
  # put the pivot to the right position
  $array[left]=$array[i]
  $array[i]=temp
  # recursive of quick sort to the left part of pivot and the right part of pivot
  quick_sort(left, i-1);
  quick_sort(i+1, right);
end
# example
# $array = [6,  1,  2, 7,  9,  3,  4,  5, 10,  8]
# quick_sort(0, $array.length - 1)
# p $array
