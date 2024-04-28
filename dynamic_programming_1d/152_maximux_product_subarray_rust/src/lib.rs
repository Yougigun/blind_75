#![allow(warnings)]
// pub fn add(left: usize, right: usize) -> usize {
//     left + right
// }

struct Solution;

impl Solution {
    pub fn max_product(nums: Vec<i32>) -> i32 {
        let mut res = nums[0];
        for i in 0..nums.len() {
            let mut temp = 1;
            for j in i..nums.len() {
                temp *= nums[j];
                res = res.max(temp);
            }
        }
        res
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // #[test]
    // fn it_works() {
    //     let result = add(2, 2);
    //     assert_eq!(result, 4);
    // }

    #[test]
    fn test_max_product() {
        let nums = vec![2, 3, -2, 4];
        let result = Solution::max_product(nums);
        assert_eq!(result, 6);

        let nums = vec![-2, 0, -1];
        let result = Solution::max_product(nums);
        assert_eq!(result, 0);
    }
}
