use rust::medium::leetcode_53::*;

#[test]
fn case_one_test() {
    assert_eq!(case_one(vec![-2,1,-3,4,-1,2,1,-5,4]), 6);
    assert_eq!(case_one(vec![1]), 1);
    assert_eq!(case_one(vec![5,4,-1,7,8]), 23);
}