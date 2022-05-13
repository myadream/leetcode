package easy._0933;

import java.util.ArrayDeque;
import java.util.Queue;

//todo 待复习
public class RecentCounter2 {
   int left, right;
   int[] times = new int[10005];

    public RecentCounter2() {
        left = 0;
        right = 0;
    }

    public int ping(int t) {
        times[right++] = t;

        while (times[left] < (t - 3000)) {
            left++;
        }

        return right - left;
    }
}
