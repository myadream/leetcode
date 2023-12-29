pub struct ReversePrefix {}

impl ReversePrefix {
    pub fn case_one(word: String, ch: char) -> String {
        match word.find(ch) {
            Some(i) => word[..=i].chars().rev().collect::<String>() + &word[i + 1..],
            None => word
        }
    }

    pub fn case_two(word: String, ch: char) -> String {
        let mut word: Vec<char>  = word.chars().collect();

        for (idx, c) in word.iter_mut().enumerate() {
            if ch != *c {
                continue;
            }

            let mut  r = idx;
            let mut  l = 0;
            while l < r {
                word.swap(l, r);
                l += 1;
                r -= 1;
            }
            break;
        }

        word.iter().collect()
    }
}