package easy

import org.testng.annotations.Test

abstract class DataSetControl {
    abstract fun buildDataSet(): ArrayList<DataSet>

    abstract fun buildImpl(): ArrayList<(DataSet) -> Any>

    @Test
    fun run() {
        for (f: (DataSet) -> Any in buildImpl()) {
            for (item: DataSet in buildDataSet()) {
                f.invoke(item.copy());
            }
        }
    }
}