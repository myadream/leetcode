use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<Vec<char>, char>, TargetData<char>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                vec!['e', 'e', 'e', 'k', 'q', 'q', 'q', 'v', 'v', 'y'],
                'e',
            ),
            TargetData::new('k'),
        ),
        DataCarrier::new(
            SourceData::new(
                vec!['c', 'f', 'j'],
                'g',
            ),
            TargetData::new('j'),
        ),
        DataCarrier::new(
            SourceData::new(
                vec!['c', 'f', 'j'],
                'a',
            ),
            TargetData::new('c'),
        ),
        DataCarrier::new(
            SourceData::new(
                vec!['c', 'f', 'j'],
                'c',
            ),
            TargetData::new('f'),
        ),
        DataCarrier::new(
            SourceData::new(
                vec!['x', 'x', 'y', 'y'],
                'z',
            ),
            TargetData::new('x'),
        ),
    ]
}

#[cfg(test)]
mod test {
    use rust::easy::leetcode_744::NextGreatestLetter;
    use crate::easy::leetcode_744::data_set;

    #[test]
    fn case_one() {
        for data in data_set() {
            assert_eq!(
                NextGreatestLetter::case_one(
                    data.source_data.value,
                    data.source_data.assist,
                ),
                data.target_data.value
            )
        }
    }

    #[test]
    fn case_two() {
        for data in data_set() {
            println!("{:?}, {:?}", data.source_data, data.target_data);
            assert_eq!(
                NextGreatestLetter::case_two(
                    data.source_data.value,
                    data.source_data.assist,
                ),
                data.target_data.value
            )
        }
    }
}