package medium._0946;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.stack", "tags.array", "tags.simulation", "difficulty.medium"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();
        Assist assist = new Assist(new int[0], new int[0]);

        dataSets.add(
                new DataSet(
                        new int[]{},
                        true,
                        assist.copy(
                                new int[]{1, 2, 3, 4, 5},
                                new int[]{4, 5, 3, 2, 1}
                        )
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{},
                        false,
                        assist.copy(
                                new int[]{1, 2, 3, 4, 5},
                                new int[]{4, 3, 5, 1, 2}
                        )
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
                    Assert.assertEquals(solution.validateStackSequencesDeque(assist.getPushed(), assist.getPoped()), dataSet.getTarget());
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getAssist();
                    Assert.assertEquals(solution.validateStackSequences(assist.getPushed(), assist.getPoped()), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}