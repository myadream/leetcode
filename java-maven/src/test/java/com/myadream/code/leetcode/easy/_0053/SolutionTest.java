package com.myadream.code.leetcode.easy._0053;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertEquals;

@Test(groups = {"tags.array", "tags.divide-and-conquer", "tags.dynamic-programming", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();
    DivideAndConquer divideAndControl = new DivideAndConquer();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<DataSet>();
        dataSets.add(
                new DataSet(
                        new int[]{-2, 1, -3, 4, -1, 2, 1, -5, 4},
                        6
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{5, 4, -1, 7, 8},
                        23
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{1},
                        1
                )
        );
        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        //动态规划
        impls.add(
                dataSet -> {
                    assertEquals(dataSet.getTarget(), solution.maxSubArray((int[]) dataSet.getSample()));
                    return true;
                }
        );

        //分治
        impls.add(
                dataSet -> {
                    assertEquals(dataSet.getTarget(), divideAndControl.maxSubArray((int[]) dataSet.getSample()));
                    return true;
                }
        );

        return impls;
    }
}