
pub fn case_one(words: Vec<String>, s: String) -> bool {
    if words.len() != s.len() {
        return false
    }

    let mut str = String::new();
    for i in 0..words.len() {
        str.push(words[i].chars().nth(0).unwrap());
    }

    str == s && s != ""
}

pub fn case_two(words: Vec<String>, s: String) -> bool {
    s == words
        .iter()
        .map(|x| x.chars().next().unwrap()) //map转换
        .collect::<String>()
}