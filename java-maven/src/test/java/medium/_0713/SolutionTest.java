package medium._0713;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertEquals;

@Test(groups = {"tags-array", "tags-slidingWindow", "difficulty.medium"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSet = new ArrayList<>();
        Assist assist = new Assist(0);

        dataSet.add(
                new DataSet(
                        new int[]{10, 5, 2, 6},
                        8,
                        assist.copy(100)
                )
        );

        dataSet.add(
                new DataSet(
                        new int[]{1, 2, 3},
                        0,
                        assist.copy(0)
                )
        );

        return dataSet;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();
        impls.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getAssist();
                    assertEquals(dataSet.getTarget(), solution.numSubarrayProductLessThanK((int[]) dataSet.getSample(), assist.getK()));

                    return true;
                }
        );

        return impls;
    }
}