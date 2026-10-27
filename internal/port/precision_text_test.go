//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestPrecisionText — degrees, arcminutes and arcseconds written back as the loaded dump writes precisions
package port

import "testing"

func TestPrecisionText(t *testing.T) {
	for want, got := range map[string]string{
		"0.0001":               precisionText(1e-4, 1),
		"0.0000010":            precisionText(1e-6, 1),
		"0.016666666666667":    precisionText(1, 60),
		"0.00027777777777778":  precisionText(1, 3600),
		"0.000027777777777778": precisionText(0.1, 3600),
		"10":                   precisionText(10, 1),
	} {
		if got != want {
			t.Fatalf("%q want %q", got, want)
		}
	}
}
