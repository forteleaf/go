package SplitLine

func SplitLine(raw string) (line []string) {
	lines := make([]string, 1)
	temp := ""
	for _, ch := range raw {
		if ch == '\r' {
			continue
		} else if ch == '\n' {
			lines = append(lines, temp)
			temp = ""
			continue
		}
		temp = temp + string(ch)
	}
	if temp != "" {
		lines = append(lines, temp)
	}
	return
}
