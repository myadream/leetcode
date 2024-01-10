use std::collections::HashMap;

pub struct IntersectionOfTwoArraysIi {}

impl IntersectionOfTwoArraysIi {
    pub fn case_one(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        if nums1.len() < nums2.len() {
            return Self::case_one(nums2, nums1);
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
                }
                None => {}
            }
        }

        ans
    }

    pub fn case_two(mut nums1: Vec<i32>, mut nums2: Vec<i32>) -> Vec<i32> {
        if nums1.len() < nums2.len() {
            return Self::case_two(nums2, nums1);
        }

        nums2.sort();
        nums1.sort();

        let (mut r, mut l) = (0, 0);
        let mut ans = vec![];

        while r < nums2.len() && l < nums1.len() {
            if nums1[l] == nums2[r] {
                ans.push(nums1[l]);
                l += 1;
                r += 1;
            } else if nums1[l] > nums2[r] {
                r += 1;
            } else {
                l += 1;
            }
        }

        ans
    }

}