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

fn data_set() -> Vec<DataCarrier<SourceData<Option<Box<ListNode>>, Option<Box<ListNode>>>, TargetData<Option<Box<ListNode>>>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                make_node_list(vec![1,2]),
                make_node_list(vec![1,2]),
            ),
            TargetData::new(make_node_list(vec![1,2]),),
        ),
    ]
}


#[test]
fn case_one_test()  {
    // for data in data_set() {
        // assert_eq!(
        //     case_one(data.source_data().value().clone(), data.source_data().assist().clone()),
        //     data.target_data().value().clone()
        // )
    // }
}