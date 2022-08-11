package medium._0002;

import easy.DataSet;
import easy.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.apache.commons.lang3.ArrayUtils;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.*;

@Test(groups = {"tags.recursion", "tags.linked-list", "tags.math", "difficulty.medium"})
class SolutionTest extends DataSetControl {


    @org.junit.jupiter.api.Test
    void main() {
        this.run();
    }

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();
        Assist assist = new Assist(new int[]{}, new int[]{});

        dataSets.add(
                new DataSet(new int[]{},new int[]{7,0,8}, assist.copy(new int[]{2,4,3}, new int[]{5,6,4}))
        );

        dataSets.add(
                new DataSet(new int[]{},new int[]{0}, assist.copy(new int[]{0}, new int[]{0}))
        );

        dataSets.add(
                new DataSet(new int[]{},new int[]{0}, assist.copy(new int[]{9,9,9,9,9,9,9}, new int[]{9,9,9,9}))
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> funs = new ArrayList<>();

        funs.add(
                item -> {

                    Assist assist = (Assist) item.getAssist();
                    int[] A1 = assist.getL1();
                    int[] A2 = assist.getL2();
                    int[] A3 = (int[]) item.getTarget();

                    ArrayUtils.reverse(A1);
                    ArrayUtils.reverse(A2);
                    ArrayUtils.reverse(A3);

                    Solution.ListNode L1 = new Solution.ListNode(0, new Solution.ListNode());
                    for (int i:assist.getL1()) {
                        L1.next.val = i;
                        L1.next.next = new Solution.ListNode();
                        L1.next = L1.next.next;
                    }

                    while (L1.val != 0) {
                        L1 = L1.next;
                        System.out.println(L1.val);
                    }

                    Solution.ListNode L2 = new Solution.ListNode();
                    for (int i:assist.getL2()) {
                       L2.val = i;
                       L2.next = new Solution.ListNode();
                       L2 = L2.next;
                    }

                    Solution.ListNode L3 = new Solution.ListNode();
                    for (int i:(int[])item.getTarget()) {
                        L3.val = i;
                        L3.next = new Solution.ListNode();
                        L3 = L3.next;
                    }

//                    L1.next = null;
//                    L2.next = null;
//                    L3.next = null;

                    Solution.ListNode listNode = (new Solution()).addTwoNumbers(L1, L2);

                    while (listNode.next != null) {

                        System.out.println(listNode.val + "--"+L3.val);
//                        assertEquals(listNode.val, L3.val);

                       listNode = listNode.next;
                       L3 = L3.next;
                    }

                    return true;
                }
        );

        return funs;
    }
}