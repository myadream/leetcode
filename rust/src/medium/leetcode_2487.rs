pub struct RemoveNodes {}

#[derive(PartialEq, Eq, Clone, Debug)]
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
impl RemoveNodes {
    pub fn case_one(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if let Some(mut cur) = head {
            cur.next = RemoveNodes::case_one(cur.next);
             if cur.val < cur.next.unwrap().val {
                 return head.unwrap().next;
            } else {
                 return Some(Box::new(cur));
            }
        }


        None
    }
}