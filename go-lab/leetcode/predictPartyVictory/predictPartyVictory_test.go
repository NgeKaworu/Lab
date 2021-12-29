package predictpartyvictory

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	tests := []struct {
		senate, res string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
	}

	for _, v := range tests {
		actual := predictPartyVictory(v.senate)
		if actual != v.res {
			t.Errorf("predictPartyVictory(%v) = %v; expected %v", v.senate, actual, v.res)
		}
	}
}
