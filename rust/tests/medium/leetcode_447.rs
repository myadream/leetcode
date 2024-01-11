use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<Vec<i32>>, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(vec![vec![0, 0], vec![1, 0], vec![2, 0]], 0),
            TargetData::new(2),
        ),
        DataCarrier::new(
            SourceData::new(vec![vec![1, 1], vec![2, 2], vec![3, 3]], 0),
            TargetData::new(2),
        ),
        DataCarrier::new(
            SourceData::new(vec![vec![1, 1]], 0),
            TargetData::new(0),
        ),
    ]
}

#[cfg(test)]
mod test {
    use rust::medium::leetcode_447::NumberOfBoomerangs;

    use crate::medium::leetcode_447::data_set;

    #[test]
    fn case_one() {
        for data in data_set() {
            assert_eq!(
                NumberOfBoomerangs::case_one(data.source_data.value),
                data.target_data.value
            );
        }
    }
}