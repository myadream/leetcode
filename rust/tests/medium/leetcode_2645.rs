use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<String, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                String::from("b"),
                0
            ),
            TargetData::new(2)
        ),
        DataCarrier::new(
            SourceData::new(
                String::from("aaa"),
                0
            ),
            TargetData::new(6)
        ),
        DataCarrier::new(
            SourceData::new(
                String::from("abc"),
                0
            ),
            TargetData::new(0)
        ),
    ]
}

#[cfg(test)]
mod test {
    use rust::medium::leetcode_2645::AddMinimum;
    use crate::medium::leetcode_2645::data_set;

    #[test]
    fn case_one() {
        for data in data_set() {
            assert_eq!(
                AddMinimum::case_one(data.source_data.value),
                data.target_data.value
            )
        }
    }
}