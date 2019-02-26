use std::collections::HashMap;

pub struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut m = HashMap::new();
        for (i, n) in nums.iter().enumerate() {
            match m.get(&(target - *n)) {
                Some(ii) => {
                    return vec![*ii as i32, i as i32];
                }
                None => {
                    m.insert(*n, i);
                }
            }
        }
        return Vec::default();
    }
}