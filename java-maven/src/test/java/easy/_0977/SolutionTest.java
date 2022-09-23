package easy._0977;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.twoPointers", "tags.array", "tags.sorting", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[]{-4, -1, 0, 3, 10},
                        new int[]{0, 1, 9, 16, 100}
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{-7, -3, 2, 3, 11},
                        new int[]{4, 9, 9, 49, 121}
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
                    Assert.assertEquals(solution.sortedSquaresSorting((int[]) dataSet.getSample()), dataSet.getTarget());
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Assert.assertEquals(solution.sortedSquareTwoPointers((int[]) dataSet.getSample()), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}