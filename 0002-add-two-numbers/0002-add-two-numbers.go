/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Formual => a + b + carry
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := new(ListNode)
    head := result
    carry := 0
    for(l1 != nil || l2 != nil) {
        a, b := 0, 0
        
        if(l1 != nil) {
            a = l1.Val
        }
        
        if(l2 != nil) {
            b = l2.Val
        }
        
        sum := a+ b + carry
        carry = sum / 10
        result.Next = &ListNode{Val: sum%10}
        result = result.Next
        
        if(l1 != nil) {
            l1 = l1.Next
        }
        
        if(l2 != nil) {
            l2 = l2.Next
        }
        
        if(carry != 0) {
            result.Next = &ListNode{Val: carry}
        }
    }
    
    return head.Next
}