// use rust::medium::leetcode_2::*;
use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                vec![2],
                3,
            ),
            TargetData::new(3),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![1, 2, 2],
                3,
            ),
            TargetData::new(0),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![3,2,3],
                3,
            ),
            TargetData::new(3),
        ),
    ]
}

#[cfg(test)]
mod tests {
    use rust::easy::leetcode_2706::ByChoco;
    use super::*;

    #[test]
    fn case_one() {
        for data in data_set() {
            assert_eq!(
                ByChoco::case_one(data.source_data.value, data.source_data.assist),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two() {
        for data in data_set() {
            assert_eq!(
                ByChoco::case_two(data.source_data.value, data.source_data.assist),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_three() {
        for data in data_set() {
            assert_eq!(
                ByChoco::case_three(data.source_data.value, data.source_data.assist),
                data.target_data.value
            )
        }
    }
}

