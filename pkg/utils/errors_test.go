package utils

import "testing"

func TestWrapWithMessage(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WrapWithMessage(tt.args.err, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("WrapWithMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrap(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_wrap(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := wrap(tt.args.err, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("wrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getFuncInfo(t *testing.T) {
	type args struct {
		pc   uintptr
		file string
		line int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFuncInfo(tt.args.pc, tt.args.file, tt.args.line); got != tt.want {
				t.Errorf("getFuncInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
