pub struct MinConst {}

impl MinConst {
    pub fn case_one(nums: Vec<i32>, x: i32) -> i64 {
        let n = nums.len();
        let mut copy: Vec<i32> = nums.clone();
        let mut sum = MinConst::sum(&copy);
        let x = x as usize;

        for i in 1..n {
            for j in 0..n {
                copy[j] = copy[j].min(nums[(i + j) % n]);
            }

            sum =  sum.min((i * x + MinConst::sum(&copy) as usize) as i64);
        }

        sum
    }

    pub fn sum(nums: &[i32]) -> i64 {
        let mut sum: i64 = 0;
        for &num in nums.iter() {
            sum += num as i64
        }

        sum
    }
}