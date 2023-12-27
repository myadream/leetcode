// https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/

use std::ops::Index;

pub struct MinConst {}

impl MinConst {

    // 容易超过时间
    // O(n2) 在子项目中多次寻找是否重复
    pub fn case_one(colors: String, needed_time: Vec<i32>) -> i32 {
        let mut index = 0;
        let len = needed_time.len();
        let mut any = 0;

        while index < len {
            let c1 = colors.chars().nth(index).unwrap();
            let mut max_time = 0;
            let mut sum = 0;

            while index < len && c1 == colors.chars().nth(index).unwrap() {
                max_time = needed_time[index].max(max_time);
                sum += needed_time[index];

                index+=1;
            }

            any += sum - max_time
        }

        any
    }

    // 通过交换最大值，在进行比较
    pub fn case_two(colors: String, mut needed_time: Vec<i32>) -> i32 {
        let mut   sum= 0;
        for index in 1..colors.len() {
            if colors.as_bytes().index(index) == colors.as_bytes().index(index - 1) {
                sum += needed_time[index].min(needed_time[index - 1]);
                needed_time[index] = needed_time[index].max(needed_time[index - 1]);
            }
        }

        sum
    }

}
