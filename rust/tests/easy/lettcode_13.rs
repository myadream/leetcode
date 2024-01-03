use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<String, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                String::from("III"),
                0,
            ),
            TargetData::new(3),
        ),

        DataCarrier::new(
            SourceData::new(
                String::from("IV"),
                0,
            ),
            TargetData::new(4),
        ),

        DataCarrier::new(
            SourceData::new(
                String::from("IX"),
                0,
            ),
            TargetData::new(9),
        ),

        DataCarrier::new(
            SourceData::new(
                String::from("LVIII"),
                0,
            ),
            TargetData::new(58),
        ),

        DataCarrier::new(
            SourceData::new(
                String::from("MCMXCIV"),
                0,
            ),
            TargetData::new(1994),
        ),

    ]
}

#[cfg(test)]
mod tests {
    use rust::easy::leetcode_13::RomanToInt;
    use super::*;

    #[test]
    fn case_oen() {
       for date in data_set() {
           assert_eq!(
               RomanToInt::case_one(
                   date.source_data.value
               ),
               date.target_data.value
           )
       }
    }
}