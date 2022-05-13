package medium._0036;

import easy.DataSet;
import easy.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

@Test(groups = {"tags.array", "tags.hashTable", "tags.matrix", "difficulty.medium"})
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
                        new char[][]{
                                new char[]{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
                                new char[]{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
                                new char[]{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
                                new char[]{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
                                new char[]{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
                                new char[]{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
                                new char[]{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
                                new char[]{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
                                new char[]{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
                        },
                        true
                )
        );


        dataSets.add(
                new DataSet(
                        new char[][]{
                                new char[]{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
                                new char[]{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
                                new char[]{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
                                new char[]{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
                                new char[]{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
                                new char[]{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
                                new char[]{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
                                new char[]{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
                                new char[]{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
                        },
                        false
                )
        );


        dataSets.add(
                new DataSet(
                        new char[][]{
                                new char[]{'.', '.', '5', '.', '.', '.', '.', '.', '.'},
                                new char[]{'.', '.', '.', '8', '.', '.', '.', '3', '.'},
                                new char[]{'.', '5', '.', '.', '2', '.', '.', '.', '.'},
                                new char[]{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
                                new char[]{'.', '.', '.', '.', '.', '.', '.', '.', '9'},
                                new char[]{'.', '.', '.', '.', '.', '.', '4', '.', '.'},
                                new char[]{'.', '.', '.', '.', '.', '.', '.', '.', '7'},
                                new char[]{'.', '1', '.', '.', '.', '.', '.', '.', '.'},
                                new char[]{'2', '4', '.', '.', '.', '.', '9', '.', '.'}
                        },
                        false
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
                    if ((boolean) dataSet.getTarget()) {
                        assertTrue(solution.isValidSudoku((char[][]) dataSet.getSample()));
                    } else {
                        assertFalse(solution.isValidSudoku((char[][]) dataSet.getSample()));
                    }

                    return true;
                }
        );

        return impls;
    }
}