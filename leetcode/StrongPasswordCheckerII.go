package leetcode

func strongPasswordCheckerII(password string) bool {
    if len(password) < 8 {
        return false
    }

    var hasLower, hasUpper, hasDigit, hasSpecial bool
    specialChars := "!@#$%^&*()-+"

    for i := 0; i < len(password); i++ {
        if i > 0 && password[i] == password[i-1] {
            return false
        }

        char := password[i]
        if char >= 'a' && char <= 'z' {
            hasLower = true
        } else if char >= 'A' && char <= 'Z' {
            hasUpper = true
        } else if char >= '0' && char <= '9' {
            hasDigit = true
        } else {
            for j := 0; j < len(specialChars); j++ {
                if char == specialChars[j] {
                    hasSpecial = true
                    break
                }
            }
        }
    }

    return hasLower && hasUpper && hasDigit && hasSpecial
}