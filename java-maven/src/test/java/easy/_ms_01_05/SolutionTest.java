package easy._ms_01_05;

import easy.DataSet;
import easy.DataSetControl;
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
        Assist assist = new Assist("");
        dataSets.add(
                new DataSet(
                        "pale",
                        true,
                        assist.copy("ple")
                )
        );

        dataSets.add(
                new DataSet(
                        "pales",
                        false,
                        assist.copy("pal")
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
                        assertTrue(solution.oneEditAway((String) dataSet.getSample(), ((Assist) dataSet.getTarget()).getEditString()));
                    } else {
                        assertFalse(solution.oneEditAway((String) dataSet.getSample(), ((Assist) dataSet.getTarget()).getEditString()));
                    }

                    return true;
                }
        );

        return impls;
    }
}