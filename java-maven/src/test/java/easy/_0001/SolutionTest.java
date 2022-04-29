package easy._0001;

import easy.DataSet;
import easy.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

@Test(groups = {"tags.array", "tags.hashTable", "difficulty.easy"})
class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[]{2, 7, 11, 14},
                        new int[]{0, 1},
                        new Assist(9)
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
                    assertArrayEquals(solution.twoSum((int[]) dataSet.getSample(), assist.getTarget()), (int[]) dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}