package medium._0167;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.Assert;
import org.junit.jupiter.api.Assertions;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.array", "tags.twoPointers", "tags.binarySearch", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();


        dataSets.add(
                new DataSet(
                        new Assist(new int[]{5, 25, 75}, 100),
                        new int[]{2, 3}
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{2,7,11,15}, 9),
                        new int[]{1, 2}
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{2,3,4}, 6),
                        new int[]{1, 3}
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{-1, 0}, -1),
                        new int[]{1, 2}
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
                    Assist assist = (Assist) dataSet.getSample();
                    Assertions.assertArrayEquals((int[]) dataSet.getTarget(), solution.twoSum(assist.getNums(), assist.getTarget()));
                    return true;
                }
        );
        return impls;
    }
}