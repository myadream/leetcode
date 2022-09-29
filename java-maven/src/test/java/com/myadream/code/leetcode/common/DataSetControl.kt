package com.myadream.code.leetcode.common

import org.testng.annotations.Test

abstract class DataSetControl {
    abstract fun buildDataSet(): ArrayList<DataSet>

    abstract fun buildImpl(): ArrayList<(DataSet) -> Any>

    @org.junit.jupiter.api.Test
    fun main() {
        this.run();
    }

    @Test
    fun run() {
        for (f: (DataSet) -> Any in buildImpl()) {
            for (item: DataSet in buildDataSet()) {
                f.invoke(item.copy());
            }
        }
    }
}