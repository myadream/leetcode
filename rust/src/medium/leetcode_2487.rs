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
        // If the list is empty, return None
        if head.is_none() {
            return None;
        }

        // Unwrap the head node since we know it's not None
        let mut boxed_head = head.unwrap();

        // If the next node exists, process it recursively
        if let Some(next_node) = boxed_head.next {

            let node = Self::case_one(Some(next_node));

            // Check if the current node's value is less than the next node's value
            if boxed_head.val < node.as_ref().unwrap().val {
                return node;
            } else {
                boxed_head.next = node;
            }
        }

        // Return the current node
        Some(boxed_head)
    }
}