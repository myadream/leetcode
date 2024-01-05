use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<i32>, i32>, TargetData<Vec<i32>>>> {
    vec![
        DataCarrier::new(
            SourceData::new(vec![10,6,8,5,11,9], 0),
            TargetData::new(vec![3,1,2,1,1,0])
        ),
        DataCarrier::new(
            SourceData::new(vec![5,1,2,3,10], 0),
            TargetData::new(vec![4,1,1,1,0])
        ),
    ]

}

#[cfg(test)]
mod test {
    use rust::hard::lettcode_1944::CanSeePersonsCount;
    use crate::hard::lettcode_1944::data_set;

    #[test]
    fn case_one() {
        for data in data_set() {
            assert_eq!(
                CanSeePersonsCount::case_one(data.source_data.value),
                data.target_data.value
            )
        }
    }
}