use std::collections::LinkedList;

pub struct CanSeePersonsCount {}

impl CanSeePersonsCount {
    pub fn case_one(heights: Vec<i32>) -> Vec<i32> {
        let mut sums: Vec<i32> = vec![0; heights.len()];
        let mut st = LinkedList::new();

        st.push_back(i32::MAX);

        for i in (0..heights.len()).rev() {
            while st.back() <= Option::from(&heights[i]) {
                st.pop_back();
                sums[i] += 1
            }

            if st.len() > 1 {
                sums[i] += 1;
            }

            st.push_back(heights[i])
        }

        sums
    }
}