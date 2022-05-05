package system

import (
	"context"
	"reflect"
	"testing"
)

func Test_initApi_InitializeData(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		i       *initApi
		args    args
		want    context.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &initApi{}
			got, err := i.InitializeData(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("initApi.InitializeData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initApi.InitializeData() = %v, want %v", got, tt.want)
			}
		})
	}
}
