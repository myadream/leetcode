#include "maxProfit.h"

int main(void) {
    int prices[] = {7,1,5, 3,6,4};
    int len = 6;
    int target = 7;
    int result = maxProfit(prices, len);

    if (target != result) {
        printf("result: %d", result);

        return 0;
    }

    int prices1[] = {1,2,3,4,5};
    len = 5;
    target = 4;
    result = maxProfit(prices1, len);
    if (target != result) {
        printf("result: %d", result);

        return 0;
    }

}