package medium._1282;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;
import java.util.List;

@Test(groups = {"tags.recursion", "tags.linked-list", "tags.math", "difficulty.medium"})
class SolutionTest extends DataSetControl {


    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(new int[]{3, 3, 3, 3, 3, 1, 3}, new int[][]{{5}, {0, 1, 2}, {3, 4, 6}})
        );

        dataSets.add(
                new DataSet(new int[]{2, 1, 3, 3, 3, 2}, new int[][]{{1}, {0, 5}, {2, 3, 4}})
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> funs = new ArrayList<>();

        funs.add(
                item -> {
                    List<List<Integer>> res = (new Solution()).groupThePeople((int[]) item.getSample());
                    for (int i = 0; i < res.size(); i++) {
                        System.out.println(res.get(i));
                    }

                    return true;
                }
        );

        return funs;
    }
}