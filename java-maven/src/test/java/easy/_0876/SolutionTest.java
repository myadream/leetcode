package easy._0876;

import common.DataSet;
import common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.Assert;
import org.junit.jupiter.api.Assertions;
import org.testng.annotations.Test;

import java.util.ArrayList;
import java.util.List;

@Test(groups = {"tags.linkedList", "tags.twoPointers", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[]{1, 2, 3, 4, 5},
                        3
                )
        );

        dataSets.add(
                new DataSet(
                        new int[]{1, 2, 3, 4, 5, 6},
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
                    ListNode head = new ListNode();
                    ListNode node = head;
                    for (int val : (int[]) dataSet.getSample()) {
                        node.next = new ListNode(val);
                        node = node.next;

                    }

                    Assertions.assertEquals((int) dataSet.getTarget(), solution.middleNodeArray(head.next).val);
                    return true;
                }
        );

                impls.add(
                dataSet -> {
                    ListNode head = new ListNode();
                    ListNode node = head;
                    for (int val : (int[]) dataSet.getSample()) {
                        node.next = new ListNode(val);
                        node = node.next;

                    }

                    Assertions.assertEquals((int) dataSet.getTarget(), solution.middleNodeTwoPointers(head.next).val);
                    return true;
                }
        );

        return impls;
    }
}