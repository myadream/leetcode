package easy._0122;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertEquals;

class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();
        dataSets.add(new DataSet(new int[]{7, 1, 5, 3, 6, 4}, 5));
        dataSets.add(new DataSet(new int[]{7, 6, 4, 3, 1}, 0));

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impl = new ArrayList<>();

        impl.add(dataSet -> {
            assertEquals(dataSet.getTarget(), solution.maxProfit((int[]) dataSet.getSample()));
            return true;
        });

        impl.add(dataSet -> {
            assertEquals(dataSet.getTarget(), solution.maxProfit2((int[]) dataSet.getSample()));
            return true;
        });

        impl.add(dataSet -> {
            assertEquals(dataSet.getTarget(), solution.maxProfit3((int[]) dataSet.getSample()));
            return true;
        });

        impl.add(dataSet -> {
            assertEquals(dataSet.getTarget(), solution.maxProfit4((int[]) dataSet.getSample()));
            return true;
        });

        return impl;
    }
}