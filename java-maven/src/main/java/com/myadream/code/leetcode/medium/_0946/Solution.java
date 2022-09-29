package com.myadream.code.leetcode.medium._0946;

import java.util.ArrayDeque;
import java.util.Deque;

public class Solution {
    public boolean validateStackSequences(int[] pushed, int[] popped) {
        //push(1),push(2),push(3),push(4),pop(4),pop(3),push(5),pop(5),pop(2)
        //1. push 推送 stack， 判断当前推送是否要 pop stack 重复
        //2. 重复情况 移除 并且移动 pop 要移除位置标记
        //3. 检查当前 push 是否还继续重复如重复重复 1， 2步骤。不重复push
        //4. 当推 push 位置 不等于 pop 位置 为不相等

        int pushFlag = 0;
        int popFlag = 0;
        int temp = 0;
        int n = pushed.length;

        while (pushFlag < n && popFlag != n) {
            temp = pushFlag;
            for (;temp > 0; temp--) {
                if (pushed[temp] == popped[popFlag]) {
                   popFlag++;
                   pushFlag--;
                }
            }
            pushFlag++;
        }

        return popFlag == n;
    }

    public boolean validateStackSequencesDeque(int[] pushed, int[] popped) {
        int n = pushed.length;
        Deque<Integer> stack = new ArrayDeque<>(n);
        for (int i = 0, j = 0; i < n; i++) {
            stack.push(pushed[i]);
            while (!stack.isEmpty() && stack.peek() == popped[j]) {
                stack.pop();
                j++;
            }
        }
        return stack.isEmpty();
    }
}
