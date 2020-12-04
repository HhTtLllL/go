package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewObjPool(t *testing.T) {
	type args struct {
		numOfObj int
	}
	tests := []struct {
		name string
		args args
		want *ObjPool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewObjPool(tt.args.numOfObj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewObjPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjPool_GetObj(t *testing.T) {
	type fields struct {
		bufchan chan *ReusableObj
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ReusableObj
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ObjPool{
				bufchan: tt.fields.bufchan,
			}
			got, err := p.GetObj(tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObj() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjPool_ReleaseObj(t *testing.T) {
	type fields struct {
		bufchan chan *ReusableObj
	}
	type args struct {
		obj *ReusableObj
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ObjPool{
				bufchan: tt.fields.bufchan,
			}
			if err := p.ReleaseObj(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("ReleaseObj() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}