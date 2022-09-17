package easy._0088;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertEquals;

@Test(groups = {"tags.array", "tags.two-pointers", "tags.Sorting", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<DataSet>();
        Assist assist = new Assist(0, 0, new int[]{});

        dataSets.add(
                new DataSet(
                        new int[]{1, 2, 3, 0, 0, 0},
                        new int[]{1, 2, 2, 3, 5, 6},
                        assist.copy(3, 2, new int[]{2, 5, 6})
                )
        );


        dataSets.add(
                new DataSet(
                        new int[]{1, 2, 3, 0, 0, 0},
                        new int[]{1, 2, 2, 3, 5, 6},
                        assist.copy(3, 2, new int[]{2, 5, 6})
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
                    int[] nums = (int[]) dataSet.getSample();
                    solution.merge((int[]) nums, assist.getM(), assist.getNums(), assist.getN());
                    assertArrayEquals((int[]) dataSet.getTarget(), nums);

                    return true;
                }
        );


        return impls;
    }
}