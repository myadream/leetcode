package medium._0036;

public class Solution {
    public boolean isValidSudoku(char[][] board) {

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (board[i][j] != '.' && !isValid(i, j, board)) {
                    return false;
                }
            }
        }

        return true;
    }

    private boolean isValid(int i, int j, char[][] board) {
        for (int row = i + 1; row < board.length; row++) {
            if (board[i][j] == board[row][j]) {
                return false;
            }
        }

        for (int col = j + 1; col < board[0].length; col++) {
            if (board[i][j] == board[i][col]) {
                return false;
            }
        }

        for (int row = i + 1; row < (i / 3) * 3 + 3; row++) {
            for (int col = (j / 3) * 3; col < (j / 3) * 3 + 3; col++) {
                if (board[i][j] == board[row][col] && (i != row && j != col)) {
                    return false;
                }
            }
        }

        return true;
    }

}
