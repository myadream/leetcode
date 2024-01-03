use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<String, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                String::from("2019-01-09"),
                0,
            ),
            TargetData::new(9),
        ),
        DataCarrier::new(
            SourceData::new(
                String::from("2019-02-10"),
                0,
            ),
            TargetData::new(41),
        ),
    ]
}

#[cfg(test)]
mod tests {
    use rust::easy::leetcode_1154::DayOfYear;
    use super::*;

    #[test]
    fn case_oen() {
       for date in data_set() {
           assert_eq!(
               DayOfYear::case_oen(
                   date.source_data.value
               ),
               date.target_data.value
           )
       }
    }
}