use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<String, i32>, TargetData<bool>>> {
    vec![
        DataCarrier::new(
            SourceData::new(String::from("()"), 0),
            TargetData::new(true),
        ),
        DataCarrier::new(
            SourceData::new(String::from("()[]{}"), 0),
            TargetData::new(true),
        ),
        DataCarrier::new(
            SourceData::new(String::from("(]"), 0),
            TargetData::new(false),
        ),
        DataCarrier::new(
            SourceData::new(String::from("}{"), 0),
            TargetData::new(false),
        ),
        DataCarrier::new(
            SourceData::new(String::from("}"), 0),
            TargetData::new(false),
        ),
    ]
}

#[cfg(test)]
mod test {
    use rust::easy::leetcode_20::ValidParentheses;
    use crate::easy::leetcode_20::data_set;

    #[test]
    fn case_one() {
        for data in data_set() {
            assert_eq!(
                ValidParentheses::case_one(data.source_data.value),
                data.target_data.value
            );
        }
    }

    #[test]
    fn case_two() {
        for data in data_set() {
            assert_eq!(
                ValidParentheses::case_two(data.source_data.value),
                data.target_data.value
            );
        }
    }
}