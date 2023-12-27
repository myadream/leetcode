#[derive(Debug)]
pub struct DataCarrier<S, T>
{
    pub source_data: S,
    pub target_data: T,
}

impl<S, T> DataCarrier<S, T>
{
    /// 构造函数
    pub fn new(source_data: S, target_data: T) -> Self {
        Self {
            source_data,
            target_data,
        }
    }
}

pub struct SourceData<V, A>
{
    pub value: V,
    pub assist: A,
}

impl<V, A> SourceData<V, A>
{
    /// 构造函数
    pub fn new(value: V, assist: A) -> Self {
        Self {
            value,
            assist,
        }
    }
}

pub struct TargetData<V>
{
    pub value: V,
}

impl<V> TargetData<V>
{
    /// 构造函数
    pub fn new(value: V) -> Self {
        Self {
            value,
        }
    }
}


