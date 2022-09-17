package medium._0073;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

@Test(groups = {"tags.array", "tags.hashTable", "tags.matrix", "difficulty.medium"})
class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[][]{
                                new int[]{1, 1, 1},
                                new int[]{1, 0, 1},
                                new int[]{1, 1, 1}
                        },
                        new int[][]{
                                new int[]{1, 0, 1},
                                new int[]{0, 0, 0},
                                new int[]{1, 0, 1}
                        }
                )
        );

        dataSets.add(
                new DataSet(
                        new int[][]{
                                new int[]{0, 1, 2, 0},
                                new int[]{3, 4, 5, 2},
                                new int[]{1, 3, 1, 5}
                        },
                        new int[][]{
                                new int[]{0, 0, 0, 0},
                                new int[]{0, 4, 5, 0},
                                new int[]{0, 3, 1, 0}
                        }
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
                    solution.setZeroes((int[][]) dataSet.getSample());
                    assertArrayEquals((int[][]) dataSet.getTarget(), (int[][]) dataSet.getSample());
                    return true;
                }
        );

        return impls;
    }
}