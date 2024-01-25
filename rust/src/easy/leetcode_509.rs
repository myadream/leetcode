
pub struct  Fib {

}

impl Fib {
    pub fn case_one(mut n: i32) -> i32 {

        let (mut ans, mut sum) = (0, 1);
        while n > 0 {
            (ans, sum) = (sum, ans + sum);
            n -= 1;
        }

        ans
    }

    pub fn case_two(n: i32) -> i32 {
        if n == 0 {
            return 0
        }

        if n <= 1 {
            return  1
        }

        return Self::case_two(n - 1) + Self::case_two(n - 2)
    }
}