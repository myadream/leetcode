package com.myadream.code.leetcode.easy._0001;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

@Test(groups = {"tags.array", "tags.hashTable", "difficulty.easy"})
class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new Assist( new int[]{2, 7, 11, 14},9),
                        new int[]{0, 1}
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
                    Assist assist = (Assist) dataSet.getSample();
                    assertArrayEquals(solution.twoSum(assist.getSample(), assist.getTarget()), (int[]) dataSet.getTarget());
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getSample();
                    assertArrayEquals(solution.twoSum2(assist.getSample(), assist.getTarget()), (int[]) dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}