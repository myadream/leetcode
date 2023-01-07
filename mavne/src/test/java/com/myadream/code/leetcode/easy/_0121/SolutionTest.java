package com.myadream.code.leetcode.easy._0121;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertEquals;

@Test(groups = {"tags.array", "tags.dynamic-programming", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<DataSet>();

        dataSets.add(
                new DataSet(
                        new int[] {
                                7,1,5,3,6,4
                        },
                        5
                )
        );

        dataSets.add(
                new DataSet(
                        new int[] {
                                7,6,4,3,1
                        },
                        0
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
                    assertEquals(dataSet.getTarget(), solution.maxProfit((int[]) dataSet.getSample()));

                    return true;
                }
        );

        return impls;
    }
}