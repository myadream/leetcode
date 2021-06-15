#include "maxProfit.h"

int maxProfit(int* prices, int pricesSize){
    int i = 1, ans = 0;

    for (;i < pricesSize; ++i) {
        if (prices[i] > prices[i - 1])
            ans += prices[i] - prices[i - 1];
    }

    return ans;
}