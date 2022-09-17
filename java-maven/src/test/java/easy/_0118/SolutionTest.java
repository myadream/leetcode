package easy._0118;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        5,
                        new int[][]{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
                )
        );

        dataSets.add(
                new DataSet(
                        1,
                        new int[][]{{1}}
                )
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        impls.add(
                dataSet -> {
                    assertArrayEquals(
                            (int[][]) dataSet.getTarget(),
                            solution.generate((Integer) dataSet.getSample()).toArray()
                    );

                    return true;
                }
        );

        return impls;
    }
}