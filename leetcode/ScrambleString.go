package leetcode

func isScramble(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }
    if len(s1) != len(s2) {
        return false
    }

    memo := map[string]bool{}
    return solve(s1, s2, memo)
}

func solve(s1, s2 string, memo map[string]bool) bool {
    if s1 == s2 {
        return true
    }

    key := s1 + "#" + s2
    if val, ok := memo[key]; ok {
        return val
    }

    n := len(s1)
    freq := [26]int{}
    for i := 0; i < n; i++ {
        freq[s1[i]-'a']++
        freq[s2[i]-'a']--
    }
    for _, f := range freq {
        if f != 0 {
            memo[key] = false
            return false
        }
    }
    for split := 1; split < n; split++ {
        noSwap := solve(s1[:split], s2[:split], memo) &&
            solve(s1[split:], s2[split:], memo)
        swapped := solve(s1[:split], s2[n-split:], memo) &&
            solve(s1[split:], s2[:n-split], memo)
        if noSwap || swapped {
            memo[key] = true
            return true
        }
    }
    memo[key] = false
    return false
}