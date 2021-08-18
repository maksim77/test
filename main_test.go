package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestSQLConverter(t *testing.T) {
	type args struct {
		sql  string
		args []interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []interface{}
		err   error
	}{
		{"case1",
			args{
				"SELECT id, name FROM table WHERE deleted = ? AND id IN(?) AND count < ?",
				[]interface{}{
					false,
					[]int{1, 6, 234},
					555},
			},
			"SELECT id, name FROM table WHERE deleted = ? AND id IN(?,?,?) AND count < ?",
			[]interface{}{false, 1, 6, 234, 555}, nil},
		{"case2",
			args{
				"SELECT id, name FROM table WHERE deleted = ? AND id IN(?) AND count < ?",
				[]interface{}{
					false,
					3,
					555},
			},
			"SELECT id, name FROM table WHERE deleted = ? AND id IN(?) AND count < ?",
			[]interface{}{false, 3, 555},
			nil},
		{"case3",
			args{
				"SELECT id, name FROM table",
				[]interface{}{},
			},
			"SELECT id, name FROM table",
			[]interface{}{},
			nil},
		{"case4",
			args{
				"SELECT id, name FROM table WHERE deleted = ? AND id IN(?) AND count < ?",
				[]interface{}{
					false},
			},
			"",
			[]interface{}{},
			errors.New("Length of sql string and args should be the same")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := SQLConverter(tt.args.sql, tt.args.args...)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("SQLConverter() err = %v, want %v", err, tt.err)
				}
			}
			if got != tt.want {
				t.Errorf("SQLConverter() got = %v, want %v", got, tt.want)
			}

			if len(got1) > 0 && !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SQLConverter() got1 = %v, want %v", got1, tt.want1)
			}

		})
	}
}
