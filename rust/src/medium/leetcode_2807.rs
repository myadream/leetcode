#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

pub struct InsertGreatestCommonDivisors {

}

impl InsertGreatestCommonDivisors {
    pub fn case_one(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if let Some(mut curr) = head {
            if let Some(next) = curr.next {
                let mut node = ListNode::new(Self::gcd(curr.val, next.val));

                node.next = Self::case_one(Some(next.clone()));
                curr.next = Some(Box::new(node));

                Some(curr)
            } else {
                Some(curr)
            }
        } else {
            head
        }
    }

    pub fn case_two(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut prov = head?;
        let head = Self::case_two(prov.next.take());
        if let Some(last) = head {
            let mut node = Box::new(ListNode::new(Self::gcd(prov.val, last.val)));
            node.next = Some(last);
            prov.next = Some(node);
        }

        Some(prov)
    }


    fn gcd(x: i32, y: i32) -> i32 {
        if y == 0 {
            return x
        }

        return Self::gcd(y, x % y)
    }
}