package main

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_bindataRead(t *testing.T) {
	type args struct {
		data []byte
		name string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := bindataRead(tArgs.data, tArgs.name)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bindataRead got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("bindataRead error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestbindataFileInfo_Name(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) bindataFileInfo
		inspect func(r bindataFileInfo, t *testing.T) //inspects receiver after test run

		want1 string
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)
			got1 := receiver.Name()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bindataFileInfo.Name got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestbindataFileInfo_Size(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) bindataFileInfo
		inspect func(r bindataFileInfo, t *testing.T) //inspects receiver after test run

		want1 int64
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)
			got1 := receiver.Size()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bindataFileInfo.Size got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestbindataFileInfo_Mode(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) bindataFileInfo
		inspect func(r bindataFileInfo, t *testing.T) //inspects receiver after test run

		want1 os.FileMode
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)
			got1 := receiver.Mode()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bindataFileInfo.Mode got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestbindataFileInfo_ModTime(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) bindataFileInfo
		inspect func(r bindataFileInfo, t *testing.T) //inspects receiver after test run

		want1 time.Time
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)
			got1 := receiver.ModTime()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bindataFileInfo.ModTime got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestbindataFileInfo_IsDir(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) bindataFileInfo
		inspect func(r bindataFileInfo, t *testing.T) //inspects receiver after test run

		want1 bool
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)
			got1 := receiver.IsDir()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bindataFileInfo.IsDir got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestbindataFileInfo_Sys(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) bindataFileInfo
		inspect func(r bindataFileInfo, t *testing.T) //inspects receiver after test run

		want1 interface{}
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)
			got1 := receiver.Sys()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bindataFileInfo.Sys got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func Test_ftwaGitignoreBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaGitignoreBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaGitignoreBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaGitignoreBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaGitignore(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaGitignore()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaGitignore got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaGitignore error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaMainGoSwpBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaMainGoSwpBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaMainGoSwpBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaMainGoSwpBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaMainGoSwp(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaMainGoSwp()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaMainGoSwp got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaMainGoSwp error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaLicenseBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaLicenseBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaLicenseBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaLicenseBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaLicense(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaLicense()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaLicense got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaLicense error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaProcfileBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaProcfileBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaProcfileBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaProcfileBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaProcfile(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaProcfile()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaProcfile got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaProcfile error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaReadmeMdBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaReadmeMdBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaReadmeMdBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaReadmeMdBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaReadmeMd(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaReadmeMd()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaReadmeMd got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaReadmeMd error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaCertShBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaCertShBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaCertShBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaCertShBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaCertSh(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaCertSh()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaCertSh got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaCertSh error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaDownloadsShBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaDownloadsShBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaDownloadsShBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaDownloadsShBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaDownloadsSh(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaDownloadsSh()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaDownloadsSh got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaDownloadsSh error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaInstallerShBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaInstallerShBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaInstallerShBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaInstallerShBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaInstallerSh(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaInstallerSh()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaInstallerSh got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaInstallerSh error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaMainGoBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaMainGoBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaMainGoBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaMainGoBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaMainGo(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaMainGo()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaMainGo got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaMainGo error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaMain_testGoBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaMain_testGoBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaMain_testGoBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaMain_testGoBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaMain_testGo(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaMain_testGo()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaMain_testGo got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaMain_testGo error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaOutTxtBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaOutTxtBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaOutTxtBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaOutTxtBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaOutTxt(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaOutTxt()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaOutTxt got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaOutTxt error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaQrShBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaQrShBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaQrShBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaQrShBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaQrSh(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaQrSh()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaQrSh got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaQrSh error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaServerCrtBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaServerCrtBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaServerCrtBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaServerCrtBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaServerCrt(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaServerCrt()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaServerCrt got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaServerCrt error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaServerKeyBytes(t *testing.T) {
	tests := []struct {
		name string

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaServerKeyBytes()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaServerKeyBytes got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaServerKeyBytes error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test_ftwaServerKey(t *testing.T) {
	tests := []struct {
		name string

		want1      *asset
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := ftwaServerKey()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ftwaServerKey got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ftwaServerKey error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestAsset(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1      []byte
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := Asset(tArgs.name)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Asset got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("Asset error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestMustAsset(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 []byte
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := MustAsset(tArgs.name)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MustAsset got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestAssetInfo(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1      os.FileInfo
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := AssetInfo(tArgs.name)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AssetInfo got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("AssetInfo error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestAssetNames(t *testing.T) {
	tests := []struct {
		name string

		want1 []string
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := AssetNames()

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AssetNames got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestAssetDir(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1      []string
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := AssetDir(tArgs.name)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AssetDir got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("AssetDir error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestRestoreAsset(t *testing.T) {
	type args struct {
		dir  string
		name string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := RestoreAsset(tArgs.dir, tArgs.name)

			if (err != nil) != tt.wantErr {
				t.Fatalf("RestoreAsset error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func TestRestoreAssets(t *testing.T) {
	type args struct {
		dir  string
		name string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			err := RestoreAssets(tArgs.dir, tArgs.name)

			if (err != nil) != tt.wantErr {
				t.Fatalf("RestoreAssets error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}

func Test__filePath(t *testing.T) {
	type args struct {
		dir  string
		name string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 string
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := _filePath(tArgs.dir, tArgs.name)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("_filePath got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
