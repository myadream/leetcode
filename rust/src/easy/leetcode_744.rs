pub struct NextGreatestLetter {}

impl NextGreatestLetter {
    pub fn case_one(letters: Vec<char>, target: char) -> char {
        let mut min_word = letters[letters.len() - 1];
        if min_word <= target {
            return letters[0];
        }

        for i in 0..letters.len() {
            if letters[i] > target && letters[i] < min_word {
                min_word = letters[i]
            }
        }

        min_word
    }

    pub fn case_two(letters: Vec<char>, target: char) -> char {
        let (mut l, mut r) = (0, letters.len() - 1);
        let mut min_word = letters[r];
        if min_word <= target {
            return letters[0];
        }

        while l < r {
            let mid = l + ((r - l) >> 1);
            if letters[mid] <= target {
                l = mid + 1;
            } else {
                r = mid;
                min_word = min_word.min(letters[mid]);
            }

        }

        min_word
    }
}