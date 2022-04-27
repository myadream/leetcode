package medium._0036;

import easy.DataSet;
import easy.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;

class SolutionTest extends DataSetControl {
    Solution solution = new Solution();

    @Test
    void main() {
        this.run();
    }

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new String[][]{
                                {"5", "3", ".", ".", "7", ".", ".", ".", "."},
                                {"6", ".", ".", "1", "9", "5", ".", ".", "."},
                                {".", "9", "8", ".", ".", ".", ".", "6", "."},
                                {"8", ".", ".", ".", "6", ".", ".", ".", "3"},
                                {"4", ".", ".", "8", ".", "3", ".", ".", "1"},
                                {"7", ".", ".", ".", "2", ".", ".", ".", "6"},
                                {".", "6", ".", ".", ".", ".", "2", "8", "."},
                                {".", ".", ".", "4", "1", "9", ".", ".", "5"},
                                {".", ".", ".", ".", "8", ".", ".", "7", "9"}
                        },
                        true
                )
        );


        dataSets.add(
                new DataSet(
                        new String[][]{
                                {"8","3",".",".","7",".",".",".","."},
                                {"6",".",".","1","9","5",".",".","."},
                                {".","9","8",".",".",".",".","6","."},
                                {"8",".",".",".","6",".",".",".","3"},
                                {"4",".",".","8",".","3",".",".","1"},
                                {"7",".",".",".","2",".",".",".","6"},
                                {".","6",".",".",".",".","2","8","."},
                                {".",".",".","4","1","9",".",".","5"},
                                {".",".",".",".","8",".",".","7","9"}
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
//        impls.add(
//                dataSet -> {
//                    ass
//
//                    return true;
//                }
//        )

        return impls;
    }
}