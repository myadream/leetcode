use rust::medium::leetcode_2::*;
use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn make_node_list(mut list_node: Vec<i32>) -> Option<Box<ListNode>> {
    if list_node.is_empty() {
        return None
    }

    let mut node = ListNode::new(list_node.pop().unwrap_or(0));
    node.next  = make_node_list(list_node);

    Some(Box::new(node))
}

fn output_node_list(nodes: Option<Box<ListNode>>) -> Vec<i32> {
    let mut output = Vec::new();
    let mut  nodes = nodes.as_ref();

    while nodes.is_some() {
        if let Some(node) = nodes {
            nodes = node.next.as_ref();
            output.push(node.val);
        }

    }

    output
}

fn data_set() -> Vec<DataCarrier<SourceData<Option<Box<ListNode>>, Option<Box<ListNode>>>, TargetData<Vec<i32>>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                make_node_list(vec![2,4,3]),
                make_node_list(vec![5,6,4]),
            ),
            TargetData::new(vec![7,0,8]),
        ),
        DataCarrier::new(
            SourceData::new(
                make_node_list(vec![0]),
                make_node_list(vec![0]),
            ),
            TargetData::new(vec![0]),
        ),
        DataCarrier::new(
            SourceData::new(
                make_node_list(vec![9,9,9,9,9,9,9]),
                make_node_list(vec![9,9,9,9]),
            ),
            TargetData::new(vec![8,9,9,9,0,0,0,1]),
        ),
    ]
}

#[cfg(test)]
mod test {
    use rust::medium::leetcode_2::{case_one, CaseTwo};
    use crate::medium::leetcode_2::{data_set, output_node_list};

    #[test]
    fn case_one_test()  {
        for data in data_set() {
            assert_eq!(
                output_node_list(case_one(data.source_data.value, data.source_data.assist)),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two_test()  {
        for data in data_set() {
            assert_eq!(
                output_node_list(CaseTwo::case(data.source_data.value, data.source_data.assist)),
                data.target_data.value
            )
        }
    }
}