package medium._0002;

import java.util.List;

class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int beyond = 0;
        ListNode sum = new ListNode();
        return sum(l1, l2 , sum, beyond);
    }

    private ListNode sum(ListNode l1, ListNode l2, ListNode head, int beyond) {
        //双节点都为空
        if (l1 == null && l2 == null) {
            return head;
        }

        int amount = beyond;
        if (l1 != null) {
            amount += l1.val;
        }

        if (l2 != null) {
           amount += l2.val;
        }

        if (amount >= 10) {
            amount >>= 1;
            beyond = 1;
        } else {
            beyond = 0;
        }

        head.val = amount;
        head.next = new ListNode(0);

        return sum(l1.next, l2.next, head, beyond);
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
