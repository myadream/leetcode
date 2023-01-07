package com.myadream.code.leetcode.medium._0019;

public class Solution {
    public ListNode removeNthFromEndTwoPointers(ListNode head, int n) {
        ListNode dummy = new ListNode(0, head);
        ListNode first = head;
        ListNode seond = dummy;

        for (int i = 0; i < n; i++) {
            first = first.next;
        }

        while (first != null) {
            first = first.next;
            seond = seond.next;
        }

        seond.next = seond.next.next;
        return dummy.next;
    }
}
