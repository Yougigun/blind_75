#![allow(warnings)]

use std::str::Bytes;

struct Solution;
impl Solution {
    // Initialize a 2D boolean array dp where dp[i][j] is true if the substring s[i..j]
    // is a palindrome, and false otherwise.

    // Every single character is a palindrome by itself, so we set dp[i][i] = true for all i.

    // A two-character substring s[i..i+1] is a palindrome if s[i] == s[i+1], so we
    // check and set dp[i][i+1] = true if s[i] == s[i+1].

    // For substrings longer than 2 characters, we use the
    // formula: dp[i][j] = dp[i+1][j-1] && s[i] == s[j].
    // This means a substring s[i..j] is a palindrome if s[i] == s[j] and the
    // substring s[i+1..j-1] is also a palindrome.
    // Note that we need to check all substrings of length 3, 4, 5, ..., n.
    // Because we have the two-character substrings already covered then we can calculate the longer substrings.
    // 2->3->4->5->...->n

    // We keep track of the start and length (or end) of the longest palindromic substring found so far.
    // When we find a longer palindrome, we update these variables.

    // Finally, we return the longest palindromic substring using the start and length/end values.

    // Analysis
    // time complexity: O(n^2)
    // space complexity: O(n^2)
    pub fn longest_palindrome(s: String) -> String {
        let n = s.len();
        if n == 0 {
            return "".to_string();
        }
        // make a 2d array of size n x n for dp
        let mut dp = vec![vec![false; n]; n];
        let mut start = 0;
        let mut max_len = 1;

        // every single character is a palindrome by itself
        for i in 0..n {
            dp[i][i] = true;
        }
        // turn string into bytes
        let s_chars = s.chars().collect::<Vec<char>>();
        // check for two-character substrings
        for i in 0..(n - 1) {
            if s_chars[i] == s_chars[i + 1] {
                dp[i][i + 1] = true;
                start = i;
                max_len = 2;
            }
        }

        // check for substring of length 3 and more
        for len in 3..=n {
            // start index and end index
            for i in 0..=n - len {
                let j = i + len - 1;
                let start_char = s_chars[i];
                let end_char = s_chars[j];
                if start_char == end_char && dp[i + 1][j - 1] {
                    dp[i][j] = true;
                    if len > max_len {
                        start = i;
                        max_len = len;
                    }
                }
            }
        }
        return s[start..(start + max_len)].into();
    }

    // loop based solution
    pub fn longest_palindrome_loop(s: String) -> String {
        let n = s.len();
        if n == 0 {
            return "".into();
        } else if n == 1 {
            return s;
        };

        let s_vec = s.as_bytes();

        let mut start = 0;
        let mut max_len = 1;
        // check for add length substring, i is the center of the palindrome
        for i in 0..n {
            // check if it is palindrome
            let (mut l, mut r) = (i, i);
            while s_vec[l] == s_vec[r] {
                if r - l + 1 > max_len {
                    start = l;
                    max_len = r - l + 1;
                }
                if l == 0 || r == n - 1 {
                    break;
                }
                l -= 1;
                r += 1;
            }
        }

        // check for even length substring, i is the center of the palindrome
        for i in 0..=(n - 2) {
            let (mut l, mut r) = (i, i + 1);
            while s_vec[l] == s_vec[r] {
                if r - l + 1 > max_len {
                    start = l;
                    max_len = r - l + 1;
                }
                if l == 0 || r == n - 1 {
                    break;
                }
                l -= 1;
                r += 1;
            }
        }
        return s[start..(start + max_len)].into();
    }

}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_longest_palindrome() {
        let s = "babad".to_string();
        let expected = "bab".to_string();
        assert_eq!(Solution::longest_palindrome(s), expected);

        let s = "cbbd".to_string();
        let expected = "bb".to_string();
        assert_eq!(Solution::longest_palindrome(s), expected);

        let s = "a".to_string();
        let expected = "a".to_string();
        assert_eq!(Solution::longest_palindrome(s), expected);

        let s = "ac".to_string();
        let expected = "a".to_string();
        assert_eq!(Solution::longest_palindrome(s), expected);
    }

    // the test cases for the loop based solution
    #[test]
    fn test_longest_palindrome_loop() {
        let s = "babad".to_string();
        let expected = "bab".to_string();
        assert_eq!(Solution::longest_palindrome_loop(s), expected);

        let s = "cbbd".to_string();
        let expected = "bb".to_string();
        assert_eq!(Solution::longest_palindrome_loop(s), expected);

        let s = "a".to_string();
        let expected = "a".to_string();
        assert_eq!(Solution::longest_palindrome_loop(s), expected);

        let s = "ac".to_string();
        let expected = "a".to_string();
        assert_eq!(Solution::longest_palindrome_loop(s), expected);
    }
}
