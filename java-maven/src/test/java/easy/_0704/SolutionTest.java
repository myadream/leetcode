package easy._0704;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.binarySearch", "tags.string", "difficulty.easy"})
class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();
        Assist assist = new Assist(new int[0], 0);

        dataSets.add(
                new DataSet(
                        "",
                        0,
                        assist.copy(new int[]{-1, 0, 5}, -1)
                )
        );

        dataSets.add(
                new DataSet(
                        "",
                        -1,
                        assist.copy(new int[]{-1, 0, 3, 5, 9, 12}, 13)
                )
        );

        dataSets.add(
                new DataSet(
                        "",
                        4,
                        assist.copy(new int[]{-1, 0, 3, 5, 9, 12}, 9)
                )
        );

        dataSets.add(
                new DataSet(
                        "",
                        -1,
                        assist.copy(new int[]{-1, 0, 3, 5, 9, 12}, 2)
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
                    Assist assist = (Assist) dataSet.getAssist();
                    Assert.assertEquals(solution.searchMid(assist.getNums(), assist.getTarget()), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}