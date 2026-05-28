package leetcode

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	i := 0

	for i < len(words) {
		lineLen := len(words[i])
		j := i + 1

		for j < len(words) && lineLen+1+len(words[j]) <= maxWidth {
			lineLen += 1 + len(words[j])
			j++
		}

		numWords := j - i
		numSpaces := maxWidth - lineLen + (numWords - 1)
		isLastLine := j == len(words)

		var line strings.Builder

		if numWords == 1 || isLastLine {
			for k := i; k < j; k++ {
				if k > i {
					line.WriteByte(' ')
				}
				line.WriteString(words[k])
			}
			for line.Len() < maxWidth {
				line.WriteByte(' ')
			}
		} else {
			gaps := numWords - 1
			spacePerGap := numSpaces / gaps
			extra := numSpaces % gaps

			for k := i; k < j; k++ {
				line.WriteString(words[k])
				if k < j-1 {
					spaces := spacePerGap
					if k-i < extra {
						spaces++
					}
					for s := 0; s < spaces; s++ {
						line.WriteByte(' ')
					}
				}
			}
		}

		result = append(result, line.String())
		i = j
	}

	return result
}
