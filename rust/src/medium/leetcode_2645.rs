pub struct AddMinimum {

}

impl AddMinimum {
    pub fn case_one(word: String) -> i32 {
        let word = word.as_bytes();
        let mut count = 1;

        for i in 1..word.len() {
            if word[i - 1] >= word[i] {
                count += 1;
            }
        }

        return count * 3 - word.len() as i32
    }
}