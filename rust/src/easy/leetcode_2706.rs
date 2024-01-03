pub struct ByChoco {}

impl ByChoco {
    pub fn case_one(mut prices: Vec<i32>, money: i32) -> i32 {
        if prices.len() < 2 {
            return money;
        }

        prices.sort_unstable();

        // 计算总价
        let remainder = money - (prices[0] + prices[1]);

        // 返回结果
        if remainder >= 0 { remainder } else { money }
    }

    pub fn case_two(mut prices: Vec<i32>, money: i32) -> i32 {
        if prices.len() < 2 {
            return money;
        }

        prices.sort_unstable();

        // 计算总价
        let remainder: i32 = prices.into_iter().take(2).sum();

        // 返回结果
        if money >= remainder { money - remainder } else { money }
    }

    pub fn case_three(prices: Vec<i32>, money: i32) -> i32 {
        if prices.len() < 2 {
            return money;
        }

        let (mut i, mut j) = (i32::MAX, i32::MAX);
        for p in prices {
            if p < i {
                (i, j) = (p, i);
                continue
            }

            if p < j {
                j = p
            }
        }

        if i + j > money {
            return money;
        }

        return money - (i + j);
    }
}