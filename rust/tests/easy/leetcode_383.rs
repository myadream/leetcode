use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<String, String>, TargetData<bool>>> {
    vec![
        DataCarrier::new(
            SourceData::new(String::from("aaaaa"), String::from("aaaaa")),
            TargetData::new(true)
        ),
        DataCarrier::new(
            SourceData::new(String::from("a"), String::from("b")),
            TargetData::new(false)
        ),
        DataCarrier::new(
            SourceData::new(String::from("aa"), String::from("b")),
            TargetData::new(false)
        ),
        DataCarrier::new(
            SourceData::new(String::from("aa"), String::from("aab")),
            TargetData::new(true)
        ),

    ]
}

#[cfg(test)]
mod test {
    use rust::easy::leetcode_383::RansomNote;
    use crate::easy::leetcode_383::data_set;

    #[test]
    fn case_oen () {
        for data in data_set() {
            assert_eq!(
                RansomNote::case_one(data.source_data.value, data.source_data.assist),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two () {
        for data in data_set() {
            assert_eq!(
                RansomNote::case_two(data.source_data.value, data.source_data.assist),
                data.target_data.value
            )
        }
    }
}