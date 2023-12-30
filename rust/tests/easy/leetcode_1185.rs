use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

struct DateFormat {
    pub year: i32,
    pub month: i32,
    pub day: i32,
}

impl DateFormat {
    pub fn new(year: i32, month: i32, day: i32) -> Self {
        Self {
            year,
            month,
            day,
        }
    }
}


fn data_set() -> Vec<DataCarrier<SourceData<DateFormat, i32>, TargetData<usize>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                DateFormat::new(2000, 2, 29),
                0,
            ),
            TargetData::new(2),
        ),
        DataCarrier::new(
            SourceData::new(
                DateFormat::new(2019, 8, 31),
                0,
            ),
            TargetData::new(6),
        ),
        DataCarrier::new(
            SourceData::new(
                DateFormat::new(1999, 7, 18),
                0,
            ),
            TargetData::new(0),
        ),
        DataCarrier::new(
            SourceData::new(
                DateFormat::new(1993, 8, 15),
                0,
            ),
            TargetData::new(0),
        ),
    ]
}


#[cfg(test)]
mod tests {
    use rust::easy::leetcode_1185::DayOfTheWeek;
    use super::*;

    fn get_week_string(day: usize) -> &'static str {
        let days = [
            "Sunday",
            "Monday",
            "Tuesday",
            "Wednesday",
            "Thursday",
            "Friday",
            "Saturday",
        ];

        days[day]
    }

    #[test]
    fn case_one() {
        for data in data_set() {
            let date = data.source_data.value;
            assert_eq!(
                DayOfTheWeek::case_one(date.day, date.month, date.year),
                get_week_string(data.target_data.value)
            )
        }
    }
}