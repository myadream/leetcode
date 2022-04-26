package easy._0118;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> list = new ArrayList<>();

        for (int i = 0; i < numRows; i++) {
            ArrayList<Integer> temp = new ArrayList<Integer>();
            for (int j = 0; j < i + 1; j++) {
                temp.add(j);
            }
        }
    }
}
