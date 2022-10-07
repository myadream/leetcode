package com.myadream.code.leetcode.medium._0870;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.jupiter.api.Assertions;

import java.util.ArrayList;

class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<DataSet>();

        dataSets.add(
                new DataSet(
                        new Assist(
                                new int[]{2, 7, 11, 15},
                                new int[]{1, 10, 4, 11}
                        ),
                        new int[]{2, 11, 7, 15}
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(
                                new int[]{12, 24, 8, 32},
                                new int[]{13, 25, 32, 11}
                        ),
                        new int[]{24, 32, 8, 12}
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
                    Assertions.assertEquals(dataSet.getTarget(), solution.advantageCount(assist.getNums1(), assist.getNums2()));
                    return true;
                }
        );

        return funs;
    }
}