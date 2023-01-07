package com.myadream.code.leetcode.easy._0455;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.jupiter.api.Assertions;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.array", "tags.greedy", "tags.twoPointers", "tags.binarySearch", "tags.sorting", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();
        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1, 2, 3}, new int[]{1, 1}),
                        1
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1, 2}, new int[]{1, 2, 3}),
                        2
                )
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> funs = new ArrayList<>();

        funs.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getSample();
                    Assertions.assertEquals(solution.findContentChildren(assist.getG(), assist.getS()), dataSet.getTarget());

                    return true;
                }
        );

        return funs;
    }
}