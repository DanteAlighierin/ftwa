package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestKeyHandler(t *testing.T) {
	tests := []struct {
		name string
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			KeyHandler()

		})
	}
}

func TestGoRoutineChannel(t *testing.T) {
	tests := []struct {
		name string

		want1 chan int
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := GoRoutineChannel()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GoRoutineChannel got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func Test_getPort(t *testing.T) {
	tests := []struct {
		name string

		want1 string
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := getPort()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getPort got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func Test_generator(t *testing.T) {
	tests := []struct {
		name string
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator()

		})
	}
}

func Test_uploadFile(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
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

			uploadFile(tArgs.w, tArgs.req)

		})
	}
}

func Test_setupRoutes(t *testing.T) {
	tests := []struct {
		name string
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupRoutes()

		})
	}
}

func Test_redirectToHTTPS(t *testing.T) {
	type args struct {
		tlsPort string
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

			redirectToHTTPS(tArgs.tlsPort)

		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()

		})
	}
}

func Test_internalIP(t *testing.T) {
	tests := []struct {
		name string

		want1      string
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := internalIP()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("internalIP got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("internalIP error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestExternalIP(t *testing.T) {
	tests := []struct {
		name string

		want1      string
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ExternalIP()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ExternalIP got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ExternalIP error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
