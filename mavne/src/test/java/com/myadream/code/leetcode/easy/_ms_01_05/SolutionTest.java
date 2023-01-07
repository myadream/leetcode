package com.myadream.code.leetcode.easy._ms_01_05;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

@Test(groups = {"tags.two-pointers", "tags.string", "difficult-easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();
        Assist assist = new Assist("","");
        dataSets.add(
                new DataSet(
                        assist.copy("pale", "ple"),
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        assist.copy("pales","pal"),
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

                    if ((Boolean) dataSet.getTarget()) {
                        assertTrue(solution.oneEditAway(((Assist) dataSet.getSample()).getSample(), ((Assist) dataSet.getSample()).getEditString()));
                    } else {
                        assertFalse(solution.oneEditAway(((Assist) dataSet.getSample()).getSample(), ((Assist) dataSet.getSample()).getEditString()));
                    }

                    return true;
                }
        );

        return impls;
    }
}