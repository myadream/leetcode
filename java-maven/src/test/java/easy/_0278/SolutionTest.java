package easy._0278;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;


@Test(groups = {"tags.binarySearch", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new Assist(5, 4),
                        4
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(1, 1),
                        1
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
                    solution.setBad(assist.getBad());
                    Assert.assertEquals(solution.firstBadVersion(assist.getN()), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}