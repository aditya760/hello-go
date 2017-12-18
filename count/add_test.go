package count

import "testing"

func TestAddthis(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name     string
		args     args
		wantTemp int
	}{
		{"counttest1", args{a: 2, b: 2}, 5},
		{"counttest2", args{a: 4, b: 2}, 6},
		{"counttest3", args{a: 9, b: 12}, 21},
	}
	for _, tt := range tests {
		if gotTemp := Addthis(tt.args.a, tt.args.b); gotTemp != tt.wantTemp {
			t.Errorf("%q. Addthis() = %v, want %v", tt.name, gotTemp, tt.wantTemp)
		}
	}
}
