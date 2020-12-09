package prose

import "strings"

// JoinWithCommas is ...
func JoinWithCommas(phase []string) string {

	var resultJoin string

	// 문자열 길이 조건
	if len(phase) == 0 {
		resultJoin = ""
	} else if len(phase) == 1 {
		resultJoin = phase[0]
	} else {
		subPhase1 := phase[:len(phase)-1]
		lastElem := phase[len(phase)-1]

		resultJoin = strings.Join(subPhase1, ", ")
		resultJoin += " and " + lastElem
	}

	// 일반 케이스
	// subPhase1 := phase[:len(phase)-1]
	// lastElem := phase[len(phase)-1]

	// resultJoin = strings.Join(subPhase1, ", ")
	// resultJoin += " and " + lastElem

	return resultJoin
}
