package com.myadream.code.leetcode.medium._0002;

class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode tail;
        ListNode head = tail = new ListNode();
        int beyond = 0;
        while (l1 != null || l2 != null || tail.val > 0) {

            System.out.println((l1 != null ? l1.val : 0) +"---"+ (l2 != null ? l2.val : 0) +"---"+ beyond);

            tail.val += l1 != null ? l1.val : 0;
            tail.val += l2 != null ? l2.val : 0;

            if (tail.val >= 10) {
                tail.val %= 10;
                beyond = 1;
            } else {
                beyond = 0;
            }

            //双节点都为空
            l1 = l1 == null ? null : l1.next;
            l2 = l2 == null ? null : l2.next;
            tail.next = new ListNode(beyond);
            tail = tail.next;
        }

        return head;
    }

    public static class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }
}
