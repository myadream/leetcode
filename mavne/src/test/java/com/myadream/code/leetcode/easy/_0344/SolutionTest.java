package com.myadream.code.leetcode.easy._0344;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.binarySearch", "tags.string", "tags.recursion", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new char[]{'h', 'e', 'l', 'l', 'o'},
                        new char[]{'o', 'l', 'l', 'e', 'h'}
                )
        );

        dataSets.add(
                new DataSet(
                        new char[]{'H', 'a', 'n', 'n', 'a', 'h'},
                        new char[]{'h', 'a', 'n', 'n', 'a', 'H'}
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
                    solution.reverseString((char[]) dataSet.getSample());
                    Assert.assertEquals((char[]) dataSet.getSample(), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}