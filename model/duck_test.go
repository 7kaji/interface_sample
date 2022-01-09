package model

import "testing"

func TestDuck_Cry(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Duck Cry",
			fields: fields{
				Name: "Donald",
			},
			want: "Quack",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDuck(tt.fields.Name)
			if got := d.Cry(); got != tt.want {
				t.Errorf("Duck.Cry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuck_GetName(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Duck GetName",
			fields: fields{
				Name: "Donald",
			},
			want: "Donald",
		},
		{
			name: "Test Dog GetName (NoName)",
			fields: fields{
				Name: "",
			},
			want: "NoName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDuck(tt.fields.Name)
			if got := d.GetName(); got != tt.want {
				t.Errorf("Dog.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
