use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(vec![0,0,1,1,1,2,2,3,3,4], 0),
            TargetData::new(5),
        ),
        DataCarrier::new(
            SourceData::new(vec![1,1,2], 0),
            TargetData::new(2),
        ),
   
    ]
}

#[cfg(test)]
mod test {
    use rust::easy::leetcode_25::RemoveDuplicates;
    use crate::easy::leetcode_25::data_set;

    #[test]
    fn case_one() {
        for mut data in data_set() {
            assert_eq!(
                RemoveDuplicates::case_one(data.source_data.value.as_mut()),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two() {
        for mut data in data_set() {
            assert_eq!(
                RemoveDuplicates::case_two(data.source_data.value.as_mut()),
                data.target_data.value
            )
        }
    }
}