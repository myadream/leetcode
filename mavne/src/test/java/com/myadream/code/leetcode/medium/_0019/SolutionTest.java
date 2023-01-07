package com.myadream.code.leetcode.medium._0019;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.junit.jupiter.api.Assertions;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.twoPointers", "tags.array", "difficulty.easy"})
class SolutionTest extends DataSetControl {

    Solution solution = new Solution();

    @NotNull
    @Override

    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1, 2, 3, 4, 5}, 2),
                        new int[]{1, 2, 3, 5}
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1}, 1),
                        new int[]{}
                )
        );

        dataSets.add(
                new DataSet(
                        new Assist(new int[]{1, 2}, 1),
                        new int[]{1}
                )
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> funs = new ArrayList<>();

        funs.add(
                item -> {

                    Assist assist = (Assist) item.getSample();
                    int[] target  = (int[])item.getTarget();

                    ListNode temp;
                    ListNode L1 = temp = new ListNode();
                    for (int i : assist.getListNode()) {
                        temp.next = new ListNode(i);
                        temp = temp.next;
                    }

                    ListNode listNode = solution.removeNthFromEndTwoPointers(L1, assist.getN());
                    int i = 0;
                    while (listNode != null) {
                        Assertions.assertEquals(listNode.val, target[i]);
                        i++;
                        listNode = listNode.next;
                    }

                    return true;
                });

        return funs;
    }
}