package com.myadream.code.leetcode.easy._0876;

public class Solution {
    public ListNode middleNodeArray(ListNode head) {
        int mid = 0;
        ListNode[] ans = new ListNode[100];


        while (head != null) {
            ans[mid++] = head;
            head = head.next;
        }

        return ans[mid / 2];
    }

    //快慢
    public ListNode middleNodeTwoPointers(ListNode head)
    {
        ListNode slow = head;
        ListNode fast = head;

        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }

        return slow;
    }
}
