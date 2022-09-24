package easy._0283;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.twoPointers", "tags.array", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[]{0, 1, 0, 3, 12},
                        new int[]{1, 3, 12, 0, 0}
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{0},
                        new int[]{0}
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{1, 2},
                        new int[]{1, 2}
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
                    solution.moveZeroesArray((int[]) dataSet.getSample());
                    Assert.assertEquals(dataSet.getSample(), dataSet.getTarget());
                    return true;
                }
        );


        impls.add(
                dataSet -> {
                    solution.moveZeroesTwoPointers((int[]) dataSet.getSample());
                    Assert.assertEquals((int[]) dataSet.getSample(), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}