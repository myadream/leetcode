// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
//
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
// 测试二
pub fn case_one(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut total = 0;
    let (mut l1, mut l2) = (l1.as_ref(), l2.as_ref());
    let mut l3 = Some(Box::new(ListNode { val: 0, next: None }));
    let mut head = l3.as_mut();

    while l1.is_some() || l2.is_some() || total > 0 {
        if let Some(node) = l1 {
            total += node.val;
            l1 = node.next.as_ref();
        }

        if let Some(node) = l2 {
            total += node.val;
            l2 = node.next.as_ref();
        }

        head.as_mut().unwrap().next = Some(Box::new(ListNode { val: total % 10, next: None }));
        head = head.unwrap().next.as_mut();

        total /= 10;
    }

    l3.unwrap().next // 返回头节点的下一个节点，即实际结果
}

pub struct CaseTwo {

}

impl CaseTwo {
    pub fn case(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        CaseTwo::helper(l1, l2, 0)
    }

    fn helper(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>, carry: i32) -> Option<Box<ListNode>> {
        if l1.is_none() && l2.is_none() && carry == 0 {
            return None
        }

        let l1 = l1.unwrap_or(Box::new(ListNode::new(0)));
        let l2 = l2.unwrap_or(Box::new(ListNode::new(0)));

        let sum = l1.val + l2.val + carry;
        let mut node = ListNode::new(sum % 10);
        node.next = CaseTwo::helper(l1.next, l2.next, sum / 10);

        Some(Box::new(node))
    }
}

pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }
}