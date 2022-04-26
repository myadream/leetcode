package easy._0566;

//
class Solution {
    public int[][] matrixReshape(int[][] mat, int r, int c) {
        int m = mat.length;
        int n = mat[0].length;
        int count = m * n;

        if (r * c != count) {
            return mat;
        }

        int[][] arr = new int[r][c];
        for (int i = 0; i < count; ++i) {
            arr[i / c][i % c] = mat[i / n][i % n];
        }

        return arr;
    }
}

