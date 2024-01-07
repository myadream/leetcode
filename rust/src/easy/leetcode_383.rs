use std::collections::{HashMap, VecDeque};

pub struct RansomNote {

}

impl RansomNote {
    pub fn case_one(ransom_note: String, magazine: String) -> bool {
        let mut sums: HashMap<char, i32> = HashMap::new();
        for word in ransom_note.chars() {
            sums.insert(word, 1);
        }

        true
    }
}