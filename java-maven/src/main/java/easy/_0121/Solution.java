package easy._0121;

public class Solution {
    public int maxProfit(int[] prices) {
        int maxPrice = 0;
        int minPrice = prices[0];

        for (int i = 1; i < prices.length; i++) {
            minPrice = Math.min(minPrice, prices[i]);
            maxPrice = Math.max(prices[i] - minPrice, maxPrice);
        }

        return maxPrice;
    }
}
