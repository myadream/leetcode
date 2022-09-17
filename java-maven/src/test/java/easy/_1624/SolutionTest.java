package easy._1624;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.Assert;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.hashTable", "tags.string", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        "mgntdygtxrvxjnwksqhxuxtrv",
                        18
                )
        );

        dataSets.add(
                new DataSet(
                        "aa",
                        0
                )
        );

        dataSets.add(
                new DataSet(
                        "abca",
                        2
                )
        );

        dataSets.add(
                new DataSet(
                        "cbzxy",
                        -1
                )
        );

        dataSets.add(
                new DataSet(
                        "cabbac",
                        4
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
                    Assert.assertEquals(solution.maxLengthBetweenEqualCharactersViolent((String) dataSet.getSample()), (int) dataSet.getTarget());
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Assert.assertEquals(solution.maxLengthBetweenEqualCharactersHasTable((String) dataSet.getSample()), (int) dataSet.getTarget());
                    return true;
                }
        );

        impls.add(
                dataSet -> {
                    Assert.assertEquals(solution.maxLengthBetweenEqualCharactersString((String) dataSet.getSample()), (int) dataSet.getTarget());
                    return true;
                }
        );

        return impls;
    }
}