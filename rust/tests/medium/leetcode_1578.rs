use rust::medium::leetcode_1578::{MinConst};

use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<String, Vec<i32>>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(String::from("aaabbbabbbb"), vec![3,5,10,7,5,3,5,5,4,8,1]),
            TargetData::new(26),
        ),
        DataCarrier::new(
            SourceData::new(String::from("aabaa"), vec![1, 2, 3, 4, 1]),
            TargetData::new(2),
        ),
        DataCarrier::new(
            SourceData::new(String::from("abaac"), vec![1, 2, 3, 4, 5]),
            TargetData::new(3),
        ),
        DataCarrier::new(
            SourceData::new(String::from("abc"), vec![1, 2, 3]),
            TargetData::new(0),
        ),

    ]
}

#[test]
fn test_case_one() {
    for data in data_set() {
        assert_eq!(
            MinConst::case_one(data.source_data.value, data.source_data.assist),
            data.target_data.value
        )
    }
}

#[test]
fn test_case_two() {
    for data in data_set() {
        assert_eq!(
            MinConst::case_two(data.source_data.value, data.source_data.assist),
            data.target_data.value
        )
    }
}