use rust::easy::leetcode_2660::IsWinner;

use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, Vec<i32>>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                vec![1,2,7,6],
                vec![6,7,0,0],
            ),
            TargetData::new(1),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![9, 7, 10, 7],
                vec![10, 2, 4, 10],
            ),
            TargetData::new(1),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![3, 6, 10, 8],
                vec![9, 9, 9, 9],
            ),
            TargetData::new(2),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![4, 10, 7, 9],
                vec![6, 5, 2, 3],
            ),
            TargetData::new(1),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![3, 5, 7, 6],
                vec![8, 10, 10, 2],
            ),
            TargetData::new(2),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![2, 3],
                vec![4, 1],
            ),
            TargetData::new(0),
        ),
    ]
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn case_one_test() {
        for data in data_set() {
            assert_eq!(
                IsWinner::case_one(data.source_data.value, data.source_data.assist),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two_test() {
        for data in data_set() {
            assert_eq!(
                IsWinner::case_two(data.source_data.value, data.source_data.assist),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_three_test() {
        for data in data_set() {
            assert_eq!(
                IsWinner::case_three(data.source_data.value, data.source_data.assist),
                data.target_data.value
            )
        }
    }
}
