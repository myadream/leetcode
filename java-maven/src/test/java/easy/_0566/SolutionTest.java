package easy._0566;

import easy.DataSet;
import easy.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @Test
    void main() {
        this.run();
    }


    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> datasets = new ArrayList<>();
        Assist assist = new Assist(0, 0);

        datasets.add(new DataSet(
                new int[][]{{1, 2}, {3, 4}},
                new int[][]{{1, 2, 3, 4}},
                assist.copy(1, 4)
        ));

        datasets.add(new DataSet(
                new int[][]{{1, 2}, {3, 4}},
                new int[][]{{1, 2}, {3, 4}},
                assist.copy(2, 4)
        ));

        return datasets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        impls.add(
                dataSet -> {
                    Assist assist = (Assist) dataSet.getAssist();
                    assertArrayEquals((int[][]) dataSet.getTarget(), solution.matrixReshape(
                            (int[][]) dataSet.getSample(),
                            assist.getR(),
                            assist.getC()
                    ));

                    return true;
                }
        );

        return impls;
    }
}