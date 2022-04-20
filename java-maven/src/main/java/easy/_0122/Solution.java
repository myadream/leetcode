package easy._0122;

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

    public int maxProfit2(int[] prices) {

    }
}
