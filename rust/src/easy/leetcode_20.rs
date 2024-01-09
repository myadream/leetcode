use std::collections::VecDeque;

pub struct ValidParentheses {}

impl ValidParentheses {
    pub fn case_one(mut s: String) -> bool {
        for _ in 0..(s.len() >> 1) {
            s = s.replace("()", "");
            s = s.replace("[]", "");
            s = s.replace("{}", "");
        }

        s.len() == 0
    }

    pub fn case_two(s: String) -> bool {
        if s.len() <= 1 {
            return false
        }

        let mut  stack: VecDeque<char> = VecDeque::new();

        for word in s.chars() {
            if word == '[' {
                stack.push_back(']')
            } else if word == '{' {
                stack.push_back('}')
            } else if word == '(' {
                stack.push_back(')')
            } else if stack.is_empty() || stack.pop_back().unwrap() != word {
                return  false
            }
        }

        stack.is_empty()
    }
}