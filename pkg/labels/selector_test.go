package labels

import "testing"

func TestRequirements_Matches(t *testing.T) {
	type args struct {
		labels Labels
	}
	tests := []struct {
		name string
		x    Requirements
		args args
		want bool
	}{
		{
			name: "empty require",
			x:    Requirements{},
			args: args{
				labels: Set{},
			},
			want: true,
		},
		{
			name: "empty labels",
			x: Requirements{
				{
					key:      "foo",
					operator: In,
					values:   []string{"bar"},
				},
			},
			args: args{
				labels: Set{},
			},
			want: false,
		},
		{
			name: "match with single value",
			x: Requirements{
				{
					key:      "foo",
					operator: In,
					values:   []string{"bar"},
				},
			},
			args: args{
				labels: Set{
					"foo": "bar",
				},
			},
			want: true,
		},
		{
			name: "not match with single value",
			x: Requirements{
				{
					key:      "foo",
					operator: In,
					values:   []string{"foo"},
				},
			},
			args: args{
				labels: Set{
					"foo": "bar",
				},
			},
			want: false,
		},
		{
			name: "match with mutil keys and values",
			x: Requirements{
				{
					key:      "foo",
					operator: In,
					values:   []string{"bar"},
				},
				{
					key:      "a",
					operator: Equals,
					values:   []string{"a"},
				},
			},
			args: args{
				labels: Set{
					"foo": "bar",
					"a":   "a",
				},
			},
			want: true,
		},
		{
			name: "not match with mutil keys and values",
			x: Requirements{
				{
					key:      "foo",
					operator: In,
					values:   []string{"bar"},
				},
				{
					key:      "a",
					operator: Equals,
					values:   []string{"a"},
				},
			},
			args: args{
				labels: Set{
					"foo": "bar",
					"a":   "b",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.Matches(tt.args.labels); got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}
