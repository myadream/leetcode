package easy

data class DataSet(val sample: Any, val target: Any, val assist: Any) {
    constructor(sample: Any, target: Any) : this(sample, target, Unit)

    constructor(sample: Any) : this(sample, Unit, Unit)
}
