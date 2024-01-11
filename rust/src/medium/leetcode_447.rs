use std::collections::HashMap;

pub struct NumberOfBoomerangs {}

impl NumberOfBoomerangs {
    pub fn case_one(points: Vec<Vec<i32>>) -> i32 {
        let mut cnt = HashMap::new();
        let mut ans = 0;
        for p in &points {
            for q in &points {
                let dst = (p[0] - q[0]) * (p[0] - q[0]) + (p[1] - q[1]) * (p[1] - q[1]);
                let v = cnt.entry(dst).or_insert(0);

                ans += *v * 2;
                *v += 1;
            }

            cnt.clear();
        }

        ans
    }
}