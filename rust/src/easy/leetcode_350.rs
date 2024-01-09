use std::collections::HashMap;

pub struct IntersectionOfTwoArraysIi {

}

impl IntersectionOfTwoArraysIi {
    pub fn case_one(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {

        if nums1.len() < nums2.len() {
            return Self::case_one(nums2, nums1)
        }

        let mut sum = HashMap::new();
        let mut ans = vec![];

        for num in nums1.into_iter() {
            *sum.entry(num).or_insert(0) += 1;
        }

        for num in nums2.into_iter() {
            match sum.get_mut(&num) {
                Some(tmp) => {
                    if *tmp >= 0 {
                        *tmp -= 1;
                        ans.push(num)
                    }
                },
                None => {}
            }
        }

        ans
    }

    pub fn case_two
}