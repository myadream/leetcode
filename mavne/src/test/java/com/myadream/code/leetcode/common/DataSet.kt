package com.myadream.code.leetcode.common

data class DataSet(val sample: Any, val target: Any) {
    constructor(sample: Any) : this(sample, Unit)
}
