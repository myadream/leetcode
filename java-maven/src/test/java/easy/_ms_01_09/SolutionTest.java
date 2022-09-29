package easy._ms_01_09;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.jupiter.api.Assertions;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.string", "tags.stringMatching", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new Assist("waterbottle", "erbottlewat"),
                        true
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist("aa", "aba"),
                        false
                )
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        //暴力
        impls.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getSample();
                    if ((boolean) dataSet.getTarget()) {
                        Assertions.assertTrue(solution.isFlipedString(assist.getS1(), assist.getS2()));
                    } else {
                        Assertions.assertFalse(solution.isFlipedString(assist.getS1(), assist.getS2()));
                    }
                    return true;
                }
        );

        return impls;
    }
}