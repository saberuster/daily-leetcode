pub mod twosum;


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_twonum() {
        let a = twosum::Solution::two_sum(vec![2, 7, 11, 15], 9);
        assert_eq!(vec![0, 1], a)
    }
}
