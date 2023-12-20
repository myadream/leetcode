#[derive(Debug)]
pub struct DataCarrier<S, T>
    where
        S: SourceDataTrait,
        T: TargetDataTrait,
{
    source_data: S,
    target_data: T,
}

pub trait DataCarrierTrait {
    type SourceData: SourceDataTrait;
    type TargetData: TargetDataTrait;

    fn source_data(&self) -> &Self::SourceData;
    fn target_data(&self) -> &Self::TargetData;
}

impl<S, T> DataCarrier<S, T>
    where
        S: SourceDataTrait,
        T: TargetDataTrait,
{
    /// 构造函数
    pub fn new(source_data: S, target_data: T) -> Self {
        Self {
            source_data,
            target_data,
        }
    }
}

impl<S, T> DataCarrierTrait for DataCarrier<S, T>
    where
        S: SourceDataTrait,
        T: TargetDataTrait,
{
    type SourceData = S;
    type TargetData = T;

    /// 获取源数据
    fn source_data(&self) -> &Self::SourceData {
        &self.source_data
    }

    /// 获取目标数据
    fn target_data(&self) -> &Self::TargetData {
        &self.target_data
    }
}


pub struct SourceData<V, A>
    where
        V:  Clone,
        A: Clone,
{
    value: V,
    assist: A,
}

impl<V, A> SourceData<V, A>
    where
        V: Clone,
        A: Clone,
{
    /// 构造函数
    pub fn new(value: V, assist: A) -> Self {
        Self {
            value,
            assist,
        }
    }
}


pub trait SourceDataTrait {
    type Value:  Clone;
    type Assist:  Clone;

    fn value(&self) -> &Self::Value;
    fn assist(&self) -> &Self::Assist;
}


impl<V, A> SourceDataTrait for SourceData<V, A>
    where
        V:  Clone,
        A:  Clone,
{
    type Value = V;
    type Assist = A;

    fn value(&self) -> &Self::Value {
        &self.value
    }

    fn assist(&self) -> &Self::Assist {
        &self.assist
    }
}

pub struct TargetData<V>
    where
        V:  Clone,
{
    value: V,
}

impl<V> TargetData<V>
    where
        V: Clone,
{
    /// 构造函数
    pub fn new(value: V) -> Self {
        Self {
            value,
        }
    }
}


pub trait TargetDataTrait {
    type Value:  Clone;

    fn value(&self) -> &Self::Value;
}


impl<V> TargetDataTrait for TargetData<V>
    where
        V: Clone,
{
    type Value = V;

    fn value(&self) -> &Self::Value {
        &self.value
    }
}