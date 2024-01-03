pub struct DayOfTheWeek {}

impl DayOfTheWeek {
    pub fn case_one(day: i32, month: i32, year: i32) -> &'static str {
        let week = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"];
        let month_days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30];

        // 1971 未非闰年 1969 为闰， 这个通过计算当前有多少不同年份
        let mut days = 365 * (year - 1971) + (year - 1969) / 4;
        days += month_days[0..(month -1) as usize ].into_iter().sum::<i32>() + day;

        if (year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)) && month >= 3 {
            days += 1;
        }

        week[((days + 3) % 7) as usize]
    }
}