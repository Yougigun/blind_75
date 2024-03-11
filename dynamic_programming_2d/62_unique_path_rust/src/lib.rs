#![allow(warnings)]
use std::collections::HashMap;
struct Solution;
impl Solution {
    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let mut mem = HashMap::<(i32, i32), i32>::new();
        Self::dp(m - 1, n - 1, &mut mem)
    }

    pub fn dp(m: i32, n: i32, mem: &mut HashMap<(i32, i32), i32>) -> i32 {
        if let Some(v) = mem.get(&(m, n)) {
            return v.to_owned();
        }
        if m == 0 || n == 0 {
            mem.insert((m, n), 1);
            return 1;
        }
        let ans = Self::dp(m - 1, n, mem) + Self::dp(m, n - 1, mem);
        mem.insert((m, n), ans);
        ans
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_unique_paths() {
        assert_eq!(Solution::unique_paths(3, 7), 28);
    }
}
