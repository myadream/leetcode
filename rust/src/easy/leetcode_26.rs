pub struct RemoveElement {}

impl RemoveElement {
    pub fn case_one(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut l= 0;
        while l < nums.len() {
            if nums[l] == val {
                nums.remove(l);
            } else {
                l += 1;
            }
        }

        nums.len() as i32
    }

    pub fn case_two(nums: &mut Vec<i32>, val: i32) -> i32 {
        let (mut i, mut j) = (0, nums.len() - 1);

        while i < j + 1 {
            if nums[i] == val {
                nums[i] = nums[j];
                j -= 1
            } else {
                i += 1
            }
        }
        
        return i as i32
    }
}