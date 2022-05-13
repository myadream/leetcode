package easy._0053;

import org.psjava.formula.Max;

import java.util.Stack;

public class DivideAndConquer {
    protected class Status {
        public int lSum, rSum, mSum, iSum;

        public Status(int lSum, int rSum, int mSum, int iSum) {
            this.lSum = lSum;
            this.rSum = rSum;
            this.mSum = mSum;
            this.iSum = iSum;
        }
    }

    public int maxSubArray(int[] nums) {
        return sum(nums, 0, nums.length - 1).mSum;
    }

    protected Status sum(int[] nums, int l, int r) {
        if (r == l) {
           return new Status(nums[l], nums[l], nums[l], nums[l]);
        }

        int m = (l + r) >> 1;
        Status lStatus = sum(nums, l, m);
        Status rStatus = sum(nums, m + 1, r);

        return pushUp(lStatus, rStatus);
    }

    //求和
    protected Status  pushUp(Status l, Status r) {
        Status sum = new Status(0, 0, 0, 0);

        sum.iSum = l.iSum + r.iSum;
        sum.rSum = Math.max(r.rSum, l.rSum + r.rSum);
        sum.lSum = Math.max(l.lSum, r.lSum + l.lSum);
        sum.mSum = Math.max(Math.max(l.mSum, r.mSum), l.rSum+ r.lSum);

        return sum;
    }

}
