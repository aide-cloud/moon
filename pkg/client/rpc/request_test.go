package rpc

import (
	"testing"
)

func TestRequest_RPCMethodName(t *testing.T) {
	type fields struct {
		c *RESTClient
	}
	type args struct {
		api     string
		path    string
		method  string
		service string
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				c: nil,
			},
			args: args{
				api:     "api",
				path:    "path",
				method:  MethodCreate,
				service: "Test",
			},
			want: "/api.path.Test/Create",
		},
		{
			name: "test",
			fields: fields{
				c: nil,
			},
			args: args{
				api:     "api",
				path:    "path",
				method:  MethodDelete,
				service: "Test",
			},
			want: "/api.path.Test/Delete",
		},
		{
			name: "test",
			fields: fields{
				c: nil,
			},
			args: args{
				api:     "api",
				path:    "path",
				method:  MethodGet,
				service: "Test",
			},
			want: "/api.path.Test/Get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Request{
				c: tt.fields.c,
			}
			x.API(tt.args.api).
				Path(tt.args.path).
				Service(tt.args.service).
				Method(tt.args.method)
			if got := x.RPCMethodName(); got != tt.want {
				t.Errorf("RPCMethodName() = %v, want %v", got, tt.want)
			}
		})
	}
}
