package com.myadream.code.leetcode.easy._0350;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;


@Test(groups = {"tags.array", "tags.hashTable", "tags.twoPointers", "tags.binarySearch", "tags.sorting", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();


    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<DataSet>();
        Assist assist = new Assist(new int[]{}, new int[]{});
        dataSets.add(
                new DataSet(
                        assist.copy(new int[]{1, 2, 2, 1}, new int[]{2, 2})
                        ,
                        new int[]{2, 2,}
                )
        );

        dataSets.add(
                new DataSet(
                        assist.copy(new int[]{4, 9, 5}, new int[]{9, 4, 9, 8, 4}),
                        new int[]{9, 4}

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
                    assertArrayEquals(
                            (int[]) dataSet.getTarget(),
                            solution.intersect(
                                    ((Assist) dataSet.getSample()).getSample(),
                                    ((Assist) dataSet.getSample()).getNumber()
                            ));

                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    assertArrayEquals(
                            (int[]) dataSet.getTarget(),
                            solution.intersect2(
                                    ((Assist) dataSet.getSample()).getSample(),
                                    ((Assist) dataSet.getSample()).getNumber()
                            ));

                    return true;
                }
        );

        return impls;
    }
}