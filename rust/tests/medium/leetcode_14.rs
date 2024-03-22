use rust::medium::leetcode_2::*;
use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, i32>, TargetData<Vec<Vec<i32>>>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                vec![-1,0,1,2,-1,-4],
                0
            ),
            TargetData::new(vec![vec![7,0,8]]),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![],
                0
            ),
            TargetData::new(vec![]),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![0],
                0
            ),
            TargetData::new(vec![]),
        ),
        
    ]
}

#[cfg(test)]
mod test {
    use rust::medium::leetcode_14::ThreeSum;
    use crate::medium::leetcode_14::{data_set};

    #[test]
    fn case_one_test()  {
        for data in data_set() {
            assert_eq!(
                ThreeSum::case_one(data.source_data.value),
                data.target_data.value
            )
        }
    }

}