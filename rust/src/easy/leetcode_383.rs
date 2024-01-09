use std::collections::HashMap;

pub struct RansomNote {}

impl RansomNote {
    pub fn case_one(ransom_note: String, magazine: String) -> bool {
        let mut sums: HashMap<char, i32> = HashMap::new();
        for word in magazine.chars().into_iter() {
            let w = sums.entry(word).or_insert(0);
            *w += 1;
        }

        for word in ransom_note.chars().into_iter() {
            match sums.get_mut(&word) {
                Some(w) => {
                    *w -= 1;
                    if *w < 0 {
                        return false
                    }
                },
                None => {
                    return false
                }
            }
        }

        true
    }

    pub fn case_two(ransom_note: String, magazine: String) -> bool {
        let mut sums = vec![0; 26];
        for word in magazine.into_bytes() {
            sums[(word - b'a') as usize] += 1;
        }

        for word in ransom_note.into_bytes() {
            sums[(word - b'a') as usize] -= 1;
            if sums[(word - b'a') as usize] <= 0 {
                return false;
            }
        }

        true
    }
}