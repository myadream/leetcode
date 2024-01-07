use rust::medium::leetcode_2807::ListNode;
use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn make_list_node(mut nums: Vec<i32>) -> Option<Box<ListNode>> {
    if nums.is_empty() {
        return None;
    }

    let mut node = ListNode::new(nums.pop().unwrap_or(0));
    node.next = make_list_node(nums);

    Some(Box::new(node))
}

fn output_list_node(mut head: Option<Box<ListNode>>) -> Vec<i32> {
    let mut nums = Vec::new();

    while head.is_some() {
        if let Some(node) = head {
            nums.push(node.val);
            head = node.next
        }
    }

    nums.into_iter().rev().collect()
}

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, i32>, TargetData<Vec<i32>>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                vec![18,6,10,3],
                0,
            ),
            TargetData::new(vec![18,6,6,2,10,1,3]),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![7],
                0,
            ),
            TargetData::new(vec![7]),
        ),
    ]
}

#[cfg(test)]
mod test {
    use rust::medium::leetcode_2807::InsertGreatestCommonDivisors;
    use crate::medium::leetcode_2807::{data_set, make_list_node, output_list_node};

    #[test]
    fn case_one() {
        for data in data_set() {
            let node = make_list_node(data.source_data.value);
            let res = output_list_node(InsertGreatestCommonDivisors::case_one(node));
            assert_eq!(
                res,
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two() {
        for data in data_set() {
            let node = make_list_node(data.source_data.value);
            let res = output_list_node(InsertGreatestCommonDivisors::case_two(node));
            assert_eq!(
                res,
                data.target_data.value
            )
        }
    }
}