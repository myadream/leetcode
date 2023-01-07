package com.myadream.code.leetcode.easy._0217;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

@Test(groups = {"tags.array", "tags.hashTable", "tags.sorting", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<DataSet>();

        dataSets.add(
                new DataSet(
                        new int[]{1, 2, 3, 1},
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{1, 2, 3, 4},
                        false
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
                        true
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
                    Boolean target = (Boolean) dataSet.getTarget();

                    if (target) {
                        assertTrue(solution.containsDuplicate((int[]) dataSet.getSample()));
                    } else {
                        assertFalse(solution.containsDuplicate((int[]) dataSet.getSample()));
                    }

                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Boolean target = (Boolean) dataSet.getTarget();

                    if (target) {
                        assertTrue(solution.containsDuplicate2((int[]) dataSet.getSample()));
                    } else {
                        assertFalse(solution.containsDuplicate2((int[]) dataSet.getSample()));
                    }

                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Boolean target = (Boolean) dataSet.getTarget();

                    if (target) {
                        assertTrue(solution.containsDuplicate3((int[]) dataSet.getSample()));
                    } else {
                        assertFalse(solution.containsDuplicate3((int[]) dataSet.getSample()));
                    }

                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Boolean target = (Boolean) dataSet.getTarget();

                    if (target) {
                        assertTrue(solution.containsDuplicate4((int[]) dataSet.getSample()));
                    } else {
                        assertFalse(solution.containsDuplicate4((int[]) dataSet.getSample()));
                    }

                    return true;
                }
        );

        return impls;
    }
}