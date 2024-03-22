use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(vec![3,2,2,3], 3),
            TargetData::new(2),
        ),
        DataCarrier::new(
            SourceData::new(vec![0,1,2,2,3,0,4,2], 2),
            TargetData::new(5),
        ),
   
    ]
}

#[cfg(test)]
mod test {
    use rust::easy::leetcode_26::RemoveElement;
    use crate::easy::leetcode_26::data_set;

    #[test]
    fn case_one() {
        for mut data in data_set() {
            assert_eq!(
                RemoveElement::case_one(data.source_data.value.as_mut(), data.source_data.assist),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two() {
        for mut data in data_set() {
            assert_eq!(
                RemoveElement::case_two(data.source_data.value.as_mut(), data.source_data.assist),
                data.target_data.value
            )
        }
    }


}