package easy._0557;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.string", "tags.twoPointers", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        "Let's take LeetCode contest",
                        "s'teL ekat edoCteeL tsetnoc"
                )
        );

        dataSets.add(
                new DataSet(
                        "God Ding",
                        "doG gniD"
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
                    Assert.assertEquals(solution.reverseWordsString((String) dataSet.getSample()), dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}