package anigi

import (
	"testing"
)

func TestPercentToPace_Range(t *testing.T) {
	for i := 0; i <= 100; i++ {
		pace := percentToPace(float64(i))
		t.Logf("percent=%d, pace=%f", i, pace)
		if pace < 0.01 || pace > 2.0 {
			t.Errorf("percent=%d, pace=%f, want in [0,2]", i, pace)
		}
	}
}
