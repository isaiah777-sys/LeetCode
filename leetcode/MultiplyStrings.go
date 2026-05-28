package leetcode

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    result := make([]int, len(num1)+len(num2))

    for i := len(num1) - 1; i >= 0; i-- {
        for j := len(num2) - 1; j >= 0; j-- {
            digit1 := int(num1[i] - '0')
            digit2 := int(num2[j] - '0')

            product := digit1 * digit2
            p1 := i + j
            p2 := i + j + 1

            sum := product + result[p2]

            result[p2] = sum % 10
            result[p1] += sum / 10
        }
    }

    var resultString []byte
    startIndex := 0
    if result[0] == 0 {
        startIndex = 1
    }

    for i := startIndex; i < len(result); i++ {
        resultString = append(resultString, byte(result[i]+'0'))
    }

    return string(resultString)
}