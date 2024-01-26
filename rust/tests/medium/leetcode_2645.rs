use crate::common::wrapper::{DataCarrier, SourceData, TargetData};

fn data_set() -> Vec<DataCarrier<SourceData<String, i32>, TargetData<i32>>> {
    vec![
        DataCarrier::new(
            SourceData::new(
                String::from("b"),
                0
            ),
            TargetData::new(2)
        ),
        DataCarrier::new(
            SourceData::new(
                String::from("aaa"),
                0
            ),
            TargetData::new(6)
        ),
        DataCarrier::new(
            SourceData::new(
                String::from("abc"),
                0
            ),
            TargetData::new(0)
        ),
    ]
}