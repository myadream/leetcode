use rust::medium::leetcode_53::*;
use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(vec![-2,1,-3,4,-1,2,1,-5,4], 0),
            TargetData::new(6),
        ),
        DataCarrier::new(
            SourceData::new(vec![1], 0),
            TargetData::new(1),
        ),
        DataCarrier::new(
            SourceData::new(vec![5,4,-1,7,8], 0),
            TargetData::new(23),
        ),
    ]
}

#[cfg(test)]
mod test {
    use rust::medium::leetcode_53::case_one;
    use crate::medium::leetcode_53::data_set;

    #[test]
    fn case_one_test() {
        for data in data_set() {
            assert_eq!(case_one(
                data.source_data.value
            ), data.target_data.value);
        }
    }
}

