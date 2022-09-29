package easy._0905;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.*;

@Test(groups = {"tags.array", "tags.twoPointers", "tags.Sorting", "difficulty.easy"})
class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[]{3, 1, 2, 4},
                        new int[]{2, 4, 3, 1}
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{0},
                        new int[]{0}
                )
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        //暴力
        impls.add(
                dataSet -> {
                    assertArrayEquals((int[]) dataSet.getTarget(), solution.sortArrayByParity((int[]) dataSet.getSample()));
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    assertArrayEquals((int[]) dataSet.getTarget(), solution.sortArrayByParity_2((int[]) dataSet.getSample()));
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    assertArrayEquals((int[]) dataSet.getTarget(), solution.sortArrayByParity_2_1((int[]) dataSet.getSample()));
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    assertArrayEquals((int[]) dataSet.getTarget(), solution.sortArrayByParity_3((int[]) dataSet.getSample()));
                    return true;
                }
        );
        return impls;
    }
}