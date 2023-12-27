use std::cmp::Ordering::{Greater, Less, Equal};

pub struct IsWinner {}

impl IsWinner {
    pub fn case_one(player1: Vec<i32>, player2: Vec<i32>) -> i32 {
        let (mut player1_sum, mut player2_sum) = (0, 0);
        for index in 0..player1.len() {
            if index > 0 && player1[index - 1] == 10 || (index > 1 && player1[index - 2] == 10) {
                player1_sum += player1[index];
            }

            if index > 0 && player2[index - 1] == 10 || (index > 1 && player2[index - 2] == 10) {
                player2_sum += player2[index];
            }

            player1_sum += player1[index];
            player2_sum += player2[index];
        }

        if player1_sum > player2_sum {
            return 1;
        } else if player1_sum < player2_sum {
            return 2;
        }

        0
    }

    pub fn case_two(player1: Vec<i32>, player2: Vec<i32>) -> i32 {
        let fun = |player: Vec<i32>| -> i32 {
            let mut pre1 = 0;
            let mut pre2 = 0;
            let mut res = 0;
            for i in player.into_iter() {
                res += i;
                if pre1 == 10 || pre2 == 10 {
                    res += i;
                }
                pre2 = pre1;
                pre1 = i;
            }
            res
        };


        let (player1_sum, player2_sum) = (fun(player1), fun(player2));
        if player1_sum > player2_sum {
            return 1;
        } else if player1_sum < player2_sum {
            return 2;
        }

        0
    }

    pub fn case_three(player1: Vec<i32>, player2: Vec<i32>) -> i32 {
        let fun = |player: &Vec<i32>| -> i32 {
            let mut res: i32 = 0;
            let mut pos = usize::MAX;
            for i in 0..player.len() {
                if pos != usize::MAX && i - pos <= 2 {
                    res += player[i]
                }

                res += player[i];
                if player[i] == 10 {
                    pos = i;
                }
            }
            res
        };


        let (player1_sum, player2_sum) = (fun(&player1), fun(&player2));

        match player1_sum.cmp(&player2_sum) {
            Greater => 1,
            Less => 2,
            Equal => 0
        }
    }
}