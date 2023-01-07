package com.myadream.code.leetcode.easy._0122;

import java.util.ArrayList;
import java.util.Comparator;

public class Solution {
    /**
     * 暴力
     *
     * @param prices
     * @return
     */
    public int maxProfit(int[] prices) {
        int len = prices.length;
        int maxProfit = 0;

        for (int i = 0; i < len - 1; i++) {
            for (int j = i + 1; j < len; j++) {
                int profit = prices[j] - prices[i];
                if (profit > maxProfit) {
                    maxProfit = profit;
                }
            }
        }

        return maxProfit;
    }

    /**
     * 单循环
     *
     * @param prices
     * @return
     */
    public int maxProfit2(int[] prices) {
        int len = prices.length;
        int maxProfit = 0;
        int minProfit = Integer.MAX_VALUE;

        for (int i = 0; i < len; i++) {
            if (minProfit > prices[i]) {
                minProfit = prices[i];
            } else if ((prices[i] - minProfit) > maxProfit) {
                maxProfit = prices[i] - minProfit;
            }
        }

        return maxProfit;
    }

    /**
     * 单循环
     *
     * @param prices
     * @return
     */
    public int maxProfit3(int[] prices) {
        ArrayList<Integer> profit = new ArrayList<>();
        profit.add(0);

        for (int i = 1; i < prices.length; i++) {
            if (prices[0] > prices[i]) {
                prices[0] = prices[i];
            } else {
                profit.add(prices[i] - prices[0]);
            }

        }
        profit.sort(Comparator.reverseOrder());

        return profit.get(0);
    }

    /**
     * 使用 Math 作为比较
     *
     * @param prices
     * @return
     */
    public int maxProfit4(int[] prices) {
        int maxProfit = 0;
        for (int i = 1; i < prices.length; i++) {
            prices[0] = Math.min(prices[0], prices[i]);
            maxProfit = Math.max(prices[i] - prices[0], maxProfit);
        }

        return maxProfit;
    }
}
