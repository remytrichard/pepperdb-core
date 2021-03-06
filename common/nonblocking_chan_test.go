package common

import (
	"reflect"
	"testing"
	"time"
)

func TestNonBlockingChan_Send(t *testing.T) {
	type fields struct {
		innerChan chan interface{}
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"1",
			fields{make(chan interface{}, 1)},
			args{0},
			true,
		},
		{
			"2",
			fields{make(chan interface{}, 0)},
			args{0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nbCh := &NonBlockingChan{
				innerChan: tt.fields.innerChan,
			}
			if got := nbCh.Send(tt.args.value); got != tt.want {
				t.Errorf("NonBlockingChan.Send() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNonBlockingChan_Recv(t *testing.T) {
	type fields struct {
		innerChan chan interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
		want1  bool
	}{
		{
			"1",
			fields{newChanWithValues(123)},
			123,
			true,
		},
		{
			"1",
			fields{make(chan interface{}, 1)},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nbCh := &NonBlockingChan{
				innerChan: tt.fields.innerChan,
			}
			got, got1 := nbCh.Recv()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonBlockingChan.Recv() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NonBlockingChan.Recv() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNonBlockingChan_SendWithDeadline(t *testing.T) {
	type fields struct {
		innerChan chan interface{}
	}
	type args struct {
		value    interface{}
		deadline time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"1",
			fields{make(chan interface{}, 1)},
			args{0, 1 * time.Second},
			true,
		},
		{
			"2",
			fields{make(chan interface{}, 0)},
			args{0, 1 * time.Second},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nbCh := &NonBlockingChan{
				innerChan: tt.fields.innerChan,
			}
			if got := nbCh.SendWithDeadline(tt.args.value, tt.args.deadline); got != tt.want {
				t.Errorf("NonBlockingChan.SendWithDeadline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNonBlockingChan_RecvWithDeadline(t *testing.T) {
	type fields struct {
		innerChan chan interface{}
	}
	type args struct {
		deadline time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		{
			"1",
			fields{newChanWithValues(123)},
			args{1 * time.Second},
			123,
			true,
		},
		{
			"1",
			fields{make(chan interface{}, 1)},
			args{1 * time.Second},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nbCh := &NonBlockingChan{
				innerChan: tt.fields.innerChan,
			}
			got, got1 := nbCh.RecvWithDeadline(tt.args.deadline)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonBlockingChan.RecvWithDeadline() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NonBlockingChan.RecvWithDeadline() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func newChanWithValues(value interface{}) chan interface{} {
	ch := make(chan interface{}, 1)
	ch <- value
	return ch
}
