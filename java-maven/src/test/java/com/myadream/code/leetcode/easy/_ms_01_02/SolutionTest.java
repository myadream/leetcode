package com.myadream.code.leetcode.easy._ms_01_02;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.hashTable", "tags.string", "tags.Sorting", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new Assist("asvnpzurz",
                                "urzsapzvn"),
                        true
                )
        );


        dataSets.add(
                new DataSet(
                        new Assist("abb", "aabb"),
                        false
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("a", "ab"),
                        false
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("abc", "bca"),
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("abc", "bad"),
                        false
                )
        );


        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {

        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        //hash table count
        impls.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getSample();
                    if ((boolean) dataSet.getTarget()) {
                        Assert.assertTrue(solution.CheckPermutationHashTableCount(assist.getS1(), assist.getS2()));
                    } else {
                        Assert.assertFalse(solution.CheckPermutationHashTableCount(assist.getS1(), assist.getS2()));
                    }
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getSample();
                    if ((boolean) dataSet.getTarget()) {
                        Assert.assertTrue(solution.CheckPermutationHashTable(assist.getS1(), assist.getS2()));
                    } else {
                        Assert.assertFalse(solution.CheckPermutationHashTable(assist.getS1(), assist.getS2()));
                    }
                    return true;
                }
        );

        return impls;
    }
}