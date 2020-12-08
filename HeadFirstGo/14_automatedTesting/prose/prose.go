package prose

import "strings"

// JoinWithCommas is ...
func JoinWithCommas(phase []string) string {

	var resultJoin string

	// resultJoin = strings.Join(phase, ",")
	subPhase1 := phase[:len(phase)-1]
	lastElem := phase[len(phase)-1]

	resultJoin = strings.Join(subPhase1, ", ")
	resultJoin += " and " + lastElem

	return resultJoin
}
