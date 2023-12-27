use crate::common::wrapper::{DataCarrier, SourceData, TargetData};
use rust::easy::leetcode_2828::{case_one};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<String>, String>, TargetData<bool>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                vec![String::from("alice"), String::from("bob"), String::from("charlie")],
                String::from("abc"),
            ),
            TargetData::new(true),
        ),
    ]
}

#[test]
fn case_one_test()  {
    for data in data_set() {
        assert_eq!(
            case_one(data.source_data.value, data.source_data.assist),
            data.target_data.value
        )
    }
}