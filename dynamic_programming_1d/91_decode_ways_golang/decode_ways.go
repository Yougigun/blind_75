package decode_ways

// s = "226"
// A = 1, B = 2,..., Z = 26
// 2(0) -> 2(1) -> 6(2)
// 2(0) -> 26(1)
// 22(0) -> 6(2)

// f(n) = the all number of ways to decode the s[n:]
// formulation
// f(n) = f(n+1)+f(n+2). plus f(n+2) only when s[n:n+2] in ("10",...,"26")
// and cause when n==len(s) - 2. will cause dp[n+2] out of bound. so tha
// f(n) = f(n+1) + f(n+2), when n == len(s)-2 and s[n:n+1] is in j to z.(10 to 26) then f(n+2) == 1
// boundary
// f(len(s) -1) = 1 when s[len(s)-1:] != "0"
// f(len(s)) = 1
// using bottom-up, init dp[i]=0 for all i excluding i == len(s) -1
// return f(0)
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 && s[0:1] != "0" {
		return 1
	} else if len(s) == 1 && s[0:1] == "0" {
		return 0
	}
	// set of j...z
	jToZ := map[string]struct{}{
		"10": {}, "11": {}, "12": {}, "13": {},
		"14": {}, "15": {}, "16": {}, "17": {},
		"18": {}, "19": {}, "20": {}, "21": {},
		"22": {}, "23": {}, "24": {}, "25": {}, "26": {}}
	dp := make([]int, len(s)+1)
	if s[len(s)-1:] != "0" {
		dp[len(s)-1] = 1
	}
	dp[len(s)] = 1 // for n == len(s)-2
	for i := len(s) - 2; i >= 0; i-- {
		if s[i:i+1] == "0" {
			continue
		}
		sum := dp[i+1]
		if _, ok := jToZ[s[i:i+2]]; ok {
			sum += dp[i+2]
		}
		dp[i] = sum
	}
	return dp[0]
}
