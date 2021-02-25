package service

import (
	"go-multiply/domain/multiply/application/command"
	"go-multiply/domain/multiply/infrastructure/service"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want command.MultiplyCommand
	}{
		{
			name: "Constructor Ok",
			want: service.NewService(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplyService_MultiplyNumbers(t *testing.T) {
	type args struct {
		numbers []float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ok",
			args: args{
				numbers: []float64{float64(5), float64(2)},
			},
			want: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &service.MultiplyService{}
			if got := ms.MultiplyNumbers(tt.args.numbers...); got != tt.want {
				t.Errorf("MultiplyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
