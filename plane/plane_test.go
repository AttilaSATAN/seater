package plane

import "testing"

func TestPlane_SeatByIndex(t *testing.T) {
	type fields struct {
		ColNames     map[int]string
		RowNumber    int
		TotalSeats   int
		AislesAfter  []int
		SeatOrderMap []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"by index", fields{map[int]string{3: "C", 6: "F", 7: "G", 9: "I", 10: "J", 1: "A", 2: "B", 4: "D", 5: "E", 8: "H"}, 6, 60, []int{3, 7}, []int{3, 4, 7, 8, 1, 10, 2, 5, 6, 9}}, args{26}, "2F"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plane{
				ColNames:     tt.fields.ColNames,
				RowNumber:    tt.fields.RowNumber,
				TotalSeats:   tt.fields.TotalSeats,
				AislesAfter:  tt.fields.AislesAfter,
				SeatOrderMap: tt.fields.SeatOrderMap,
			}
			if got := p.SeatByIndex(tt.args.i); got != tt.want {
				t.Errorf("Plane.SeatByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
