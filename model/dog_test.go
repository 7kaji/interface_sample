package model

import (
	"testing"
)

func TestDog_Cry(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Dog Cry",
			fields: fields{
				Name: "Pochi",
			},
			want: "Bow-wow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDog(tt.fields.Name)
			if got := d.Cry(); got != tt.want {
				t.Errorf("Dog.Cry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDog_GetName(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Dog GetName",
			fields: fields{
				Name: "Pochi",
			},
			want: "Pochi",
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
			d := NewDog(tt.fields.Name)
			if got := d.GetName(); got != tt.want {
				t.Errorf("Dog.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
