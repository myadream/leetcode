use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<i32, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(2, 0),
            TargetData::new(2),
        ),
        DataCarrier::new(
            SourceData::new(3, 0),
            TargetData::new(3),
        ),

    ]
}

#[cfg(test)]
mod test {
    use rust::easy::leetcode_70::ClimbingStairs;
    use crate::easy::leetcode_70::data_set;

    #[test]
    fn case_one() {
        for data in data_set() {
            assert_eq!(
                ClimbingStairs::case_one(data.source_data.value),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two() {
        for data in data_set() {
            assert_eq!(
                ClimbingStairs::case_two(data.source_data.value),
                data.target_data.value
            )
        }
    }
}