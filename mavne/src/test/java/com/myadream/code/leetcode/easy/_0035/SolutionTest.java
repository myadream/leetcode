package com.myadream.code.leetcode.easy._0035;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;
@Test(groups = {"tags.binarySearch", "tags.array","difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1, 3, 5, 6}, 5),
                        2
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1, 3, 5, 6}, 2),
                        1
                )
        );
        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1, 3, 5, 6}, 7),
                        4
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
                    Assert.assertEquals(solution.searchInsert(assist.getNums(), assist.getTarget()), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}