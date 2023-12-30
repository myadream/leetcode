pub struct DayOfYear {}

impl DayOfYear {
    pub fn case_oen(date: String) -> i32 {
        let month_days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30];
        let mut date = date.split_terminator("-");
        let year = date.next().unwrap().parse::<i32>();
        let month = date[5..7].parse::<i32>().unwrap();
        let day = date[8..].parse::<i32>().unwrap();

        let mut sum = month_days[..(month - 1) as usize].iter().sum::<i32>() + day;
        if (year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)) && month >= 3 {
            sum += 1;
        }

        sum
    }
}