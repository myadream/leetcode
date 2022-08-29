package easy._1470;

import easy.DataSet;
import easy.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

@Test(groups = {"tags.array", "difficulty.easy"})
class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @org.junit.jupiter.api.Test
    void main() {
        this.run();
    }

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[]{},
                        new int[]{2, 3, 5, 4, 1, 7},
                        new Assist(new int[]{2, 5, 1, 3, 4, 7}, 3)
                )
        );


        dataSets.add(
                new DataSet(
                        new int[]{},
                        new int[]{1, 4, 2, 3, 3, 2, 4, 1},
                        new Assist(new int[]{1, 2, 3, 4, 4, 3, 2, 1}, 4)
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{},
                        new int[]{1, 2, 1, 2},
                        new Assist(new int[]{1, 1, 2, 2}, 2)
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
                    assertArrayEquals(solution.shuffle(assist.getNums(), assist.getN()), (int[]) dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}