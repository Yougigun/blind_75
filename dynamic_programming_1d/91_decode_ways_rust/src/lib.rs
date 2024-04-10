#![allow(warnings)]

struct Solution;
// https://leetcode.com/problems/decode-ways/
impl Solution {
    // Thoughts:
    // we can use recursion to solve this problem.
    // And when we are doing recursion, we can use memorization to avoid duplicate calculation.
    // m[string] = possible ways to decode the string.

    // because we can use top-down approach of dp, we can also think of bottom-up approach.
    // dp[i]: possible ways to decode the string from 0 to i.
    // dp[i] = dp[i-1] + dp[i-2](only if i-1..=i is a valid number)
    // dp[0] = 1 if s[0] != "0"

    pub fn num_decodings(s: String) -> i32 {
        let s = s.chars().collect::<Vec<char>>();
        if s[0] == '0' {
            return 0;
        }
        let n = s.len();
        let mut dp = vec![0; n];
        dp[0] = if s[0] != '0' { 1 } else { 0 };
        for i in 1..n {
            let mut ways = if s[i] != '0' { dp[i - 1] } else { 0 };
            if s[i - 1] == '1' {
                if i > 2 {
                    ways += dp[i - 2];
                } else {
                    ways += 1;
                }
            } else if s[i - 1] == '2' && s[i] <= '6' {
                if i > 2 {
                    ways += dp[i - 2];
                } else {
                    ways += 1;
                }
            }
            dp[i] = ways;
        }

        dp[n - 1]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_num_decodings() {
        let s = "12".to_string();
        let result = Solution::num_decodings(s);
        assert_eq!(result, 2);

        let s = "226".to_string();
        let result = Solution::num_decodings(s);
        assert_eq!(result, 3);

        let s = "0".to_string();
        let result = Solution::num_decodings(s);
        assert_eq!(result, 0);

        let s = "06".to_string();
        let result = Solution::num_decodings(s);
        assert_eq!(result, 0);
        // add "10" case
        let s = "10".to_string();
        let result = Solution::num_decodings(s);
        assert_eq!(result, 1);
        // add "012" case
        let s = "012".to_string();
        let result = Solution::num_decodings(s);
        assert_eq!(result, 0);
    }
}
