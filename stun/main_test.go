package test

import (
	"testing"
)

func Test_test(t *testing.T) {
	type args struct {
		t *testing.T
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			test(tArgs.t)

		})
	}
}
