package com.myadream.code.leetcode.easy._0020;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;


@Test(groups = {"tags.stack", "tags.string", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

               dataSets.add(
                new DataSet(
                        "(([]){})",
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        "()",
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        "()[]{}",
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        "(]",
                        false
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
                    if ((boolean) dataSet.getTarget()) {
                        Assert.assertTrue(solution.isValid((String) dataSet.getSample()));
                    } else {
                        Assert.assertFalse(solution.isValid((String) dataSet.getSample()));
                    }
                    return true;
                }
        );

        return impls;
    }
}