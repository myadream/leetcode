package com.myadream.code.leetcode.easy._0088;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

@Test(groups = {"tags.array", "tags.two-pointers", "tags.Sorting", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<DataSet>();
        Assist assist = new Assist(new int[]{}, 0, 0, new int[]{});

        dataSets.add(
                new DataSet(
                        assist.copy(new int[]{1, 2, 3, 0, 0, 0}, 3, 2, new int[]{2, 5, 6}),
                        new int[]{1, 2, 2, 3, 5, 6}
                )
        );

        dataSets.add(
                new DataSet(
                        assist.copy(new int[]{1, 2, 3, 0, 0, 0}, 3, 2, new int[]{2, 5, 6}),
                        new int[]{1, 2, 2, 3, 5, 6}
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
                    solution.merge((int[]) assist.getSample(), assist.getM(), assist.getNums(), assist.getN());
                    assertArrayEquals((int[]) dataSet.getTarget(), assist.getNums());

                    return true;
                }
        );


        return impls;
    }
}