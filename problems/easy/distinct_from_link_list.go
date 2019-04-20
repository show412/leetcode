func distinct(head *SinglyLinkedListNode) *SinglyLinkedListNode {
    // Write your code here
    m := make(map[int32]bool)
    returnList := &SinglyLinkedListNode{}
    returnList.next = head
    preList := returnList
    for preList.next != nil {
        if _, ok := m[preList.next.data]; ok  {
            preList.next = preList.next.next
        } else {
            m[preList.next.data] = true
            preList = preList.next
        }
    }
    return returnList.next
}
