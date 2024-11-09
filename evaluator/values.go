package evaluator

// "To qualify for the leaderboard your value for A + B + C must be less than 50"
const MaxValue = 50

type Values struct {
	A uint
	B uint
	C uint
}

func (e *Evaluator) GetValues() []Values {
	return e.values
}

// generateAllValues generates all combinations for A, B, and C up to the max value
func generateAllValues() []Values {
	valuesList := make([]Values, 0)
	for a := uint(1); a < MaxValue; a++ {
		for b := uint(1); b < MaxValue; b++ {
			if a == b {
				continue
			}
			for c := uint(1); c < MaxValue; c++ {
				sumCondition := a != c && b != c && a+b+c < MaxValue
				if sumCondition {
					valuesList = append(valuesList, Values{A: a, B: b, C: c})
				}
			}
		}
	}

	return valuesList
}
