pub struct RemoveDuplicates {}

impl RemoveDuplicates {
    pub fn case_one(nums: &mut Vec<i32>) -> i32 {
        let mut l= 1;
        while l < nums.len() {
            if nums[l] == nums[l - 1] {
                nums.remove(l);
            } else {
                l += 1;
            }
        }

        nums.len() as i32
    }

    pub fn case_two(nums: &mut Vec<i32>) -> i32 {
        if nums.len() == 0 {
            return 0
        }
        
        let mut count = 0;

        for i in 0..nums.len() {
            if nums[i] != nums[count] {
                count += 1;
                nums[count] = nums[i];
            }
        }

        (count + 1) as i32
    }
}