package easy._0905;

import easy.DataSet;
import easy.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Tesvbt;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.*;

@Test(groups = {"tags.array", "tags.two-pointers", "tags.Sorting", "difficulty.easy", "d-single"})
class SolutionTest extends DataSetControl {

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[]{3, 1, 2, 4},
                        new int[]{2, 4, 3, 1}
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{0},
                        new int[]{0}
                )
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        return impls;
    }
}