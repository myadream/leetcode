pub struct ClimbingStairs {}

impl ClimbingStairs {
    pub fn case_one(n: i32) -> i32 {
        fn dfs(i: i32) -> i32 {
            if i <= 1 {
                return 1;
            }

            dfs(i - 1) + dfs(i - 2)
        }

        dfs(n - 1) + dfs(n - 2)
    }

    pub fn case_two(n: i32) -> i32 {
        let (mut f2, mut ans) = (0, 1);
        for _ in 0..n {
            let f1 = f2;
            f2 = ans;
            ans = f1 + f2;
        }

        ans
    }
}