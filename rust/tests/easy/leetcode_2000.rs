use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<String, char>, TargetData<String>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                String::from("abcdefd"),
                'd'
            ),
            TargetData::new(String::from("dcbaefd")),
        ),
        DataCarrier::new(
            SourceData::new(
                String::from("xyxzxe"),
                'z'
            ),
            TargetData::new(String::from("zxyxxe")),
        ),
        DataCarrier::new(
            SourceData::new(
                String::from("abcd"),
                'z'
            ),
            TargetData::new(String::from("abcd")),
        ),
    ]
}

#[cfg(test)]
mod tests {
    use rust::easy::leetcode_2000::ReversePrefix;
    use crate::easy::leetcode_2000::data_set;

    #[test]
    fn case_one() {
        for d in data_set() {
            assert_eq!(
                ReversePrefix::case_one(d.source_data.value, d.source_data.assist),
                d.target_data.value
            )
        }
    }

    #[test]
    fn case_two() {
        for d in data_set() {
            assert_eq!(
                ReversePrefix::case_two(d.source_data.value, d.source_data.assist),
                d.target_data.value
            )
        }
    }
}