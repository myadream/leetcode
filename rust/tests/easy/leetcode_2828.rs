use crate::common::wrapper::{DataCarrier, DataCarrierTrait, SourceData, SourceDataTrait, TargetData, TargetDataTrait};
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
            case_one(data.source_data().value().clone(), data.source_data().assist().clone().clone()),
            data.target_data().value().clone()
        )
    }
}