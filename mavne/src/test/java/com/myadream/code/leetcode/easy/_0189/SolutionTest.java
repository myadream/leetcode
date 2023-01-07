package com.myadream.code.leetcode.easy._0189;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;
@Test(groups = {"tags.twoPointers", "tags.array", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {

        ArrayList<DataSet> dataSets = new ArrayList<>();
        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1,2,3,4,5,6,7}, 3),
                        new int[]{5,6,7,1,2,3,4}
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{-1,-100,3,99}, 2),
                        new int[]{3,99,-1,-100}
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
                    solution.rotateArraySplice(assist.getNums(), assist.getK());
                    Assert.assertEquals(assist.getNums(), dataSet.getTarget());
                    return true;
                }
        );



        return impls;
    }
}