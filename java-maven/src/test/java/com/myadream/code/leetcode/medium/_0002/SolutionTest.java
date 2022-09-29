package com.myadream.code.leetcode.medium._0002;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

@Test(groups = {"tags.recursion", "tags.linked-list", "tags.math", "difficulty.medium"})
class SolutionTest extends DataSetControl {

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();
        Assist assist = new Assist(new int[]{}, new int[]{});

        dataSets.add(
                new DataSet(assist.copy(new int[]{2, 4, 3}, new int[]{5, 6, 4}), new int[]{7, 0, 8})
        );

        dataSets.add(
                new DataSet(assist.copy(new int[]{0}, new int[]{0}), new int[]{0})
        );

        dataSets.add(
                new DataSet(assist.copy(new int[]{9, 9, 9, 9, 9, 9, 9}, new int[]{9, 9, 9, 9}), new int[]{0})
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

                    Solution.ListNode temp;

                    Solution.ListNode L1 = temp = new Solution.ListNode();
                    for (int i : assist.getL1()) {
                        temp.next = new Solution.ListNode(i);
                        temp = temp.next;
                    }

                    Solution.ListNode L2 = temp = new Solution.ListNode();
                    for (int i : assist.getL2()) {
                        temp.next = new Solution.ListNode(i);
                        temp = temp.next;
                    }

                    Solution.ListNode L3 = temp = new Solution.ListNode();
                    for (int i : assist.getL2()) {
                        temp.next = new Solution.ListNode(i);
                        temp = temp.next;
                    }

                    Solution.ListNode listNode = (new Solution()).addTwoNumbers(L1.next, L2.next);

                    while (listNode.next != null) {
                        System.out.println(listNode.val);
                        listNode = listNode.next;

                    }

                                            System.out.println("========");

                    return true;
                }
        );

        return funs;
    }
}