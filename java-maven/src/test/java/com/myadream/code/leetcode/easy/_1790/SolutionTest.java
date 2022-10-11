package com.myadream.code.leetcode.easy._1790;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.jupiter.api.Assertions;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.hashTable", "tags.strings", "tags.count", "tags.counting", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();


        dataSets.add(
                new DataSet(
                        new Assist("npv", "zpn"),
                        false
                )
        );


        dataSets.add(
                new DataSet(
                        new Assist("caa", "aaz"),
                        false
                )
        );


        dataSets.add(
                new DataSet(
                        new Assist("aa", "ac"),
                        false
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("ab", "ba"),
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("bank", "kanb"),
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("attack", "defend"),
                        false
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("kelb", "kelb"),
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("abcd", "dcba"),
                        false
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
                    if ((boolean) dataSet.getTarget()) {
                        Assertions.assertTrue(solution.areAlmostEqualCount(assist.getS1(), assist.getS2()));
                    } else {
                        Assertions.assertFalse(solution.areAlmostEqualCount(assist.getS1(), assist.getS2()));
                    }

                    return true;
                }
        );

        return funs;
    }
}