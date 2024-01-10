use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, Vec<i32>>, TargetData<Vec<i32>>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                vec![4,7,9,7,6,7],
                vec![5,0,0,6,1,6,2,2,4],
            ),
            TargetData::new(
                vec![4, 6]
            ),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![4, 9, 5],
                vec![9, 4, 9, 8, 4],
            ),
            TargetData::new(
                vec![4, 9]
            ),
        ),
        DataCarrier::new(
            SourceData::new(
                vec![1, 2, 2, 1],
                vec![2, 2],
            ),
            TargetData::new(
                vec![2, 2]
            ),
        ),

    ]
}

#[cfg(test)]
mod test {
    use rust::easy::leetcode_350::IntersectionOfTwoArraysIi;

    use crate::easy::leetcode_350::data_set;

    #[test]
    fn case_one() {
        for data in data_set() {
            assert_eq!(
                IntersectionOfTwoArraysIi::case_one(data.source_data.value, data.source_data.assist),
                data.target_data.value
            );
        }
    }

    #[test]
    fn case_two() {
        for data in data_set() {
            assert_eq!(
                IntersectionOfTwoArraysIi::case_two(data.source_data.value, data.source_data.assist),
                data.target_data.value
            );
        }
    }
}