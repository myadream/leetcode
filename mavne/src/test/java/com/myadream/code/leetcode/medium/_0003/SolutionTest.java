package com.myadream.code.leetcode.medium._0003;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.String", "tags.hashTable", "tags.slidingWindow", "difficulty.medium"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        "abcabcbb",
                        3
                )
        );


        dataSets.add(
                new DataSet(
                        "bbbbb",
                        1
                )
        );

        dataSets.add(
                new DataSet(
                        "pwwkew",
                        3
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
                    Assert.assertEquals(solution.lengthOfLongestSubstring((String) dataSet.getSample()), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}