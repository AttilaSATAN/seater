package plane

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		pp PlanePrototype
	}
	tests := []struct {
		name    string
		args    args
		want    Plane
		wantErr bool
	}{
		{name: "Success", args: args{PlanePrototype{"THY800", 60, 10, []int{3, 7}}}, want: Plane{"THY800", map[int]string{1: "A", 2: "B", 10: "J", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H", 9: "I"}, 6, 60, []int{3, 7}, []int{3, 4, 7, 8, 1, 10, 2, 5, 6, 9}}, wantErr: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.pp)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
