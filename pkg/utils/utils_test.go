package utils

import (
	"net/http"
	"reflect"
	"testing"
)

func TestWriteResponse(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		statusCode int
		message    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteResponse(tt.args.w, tt.args.statusCode, tt.args.message)
		})
	}
}

func TestHandleServiceFunc(t *testing.T) {
	type args struct {
		w           http.ResponseWriter
		r           *http.Request
		serviceFunc func() (interface{}, error)
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleServiceFunc(tt.args.w, tt.args.r, tt.args.serviceFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleServiceFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleServiceFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
