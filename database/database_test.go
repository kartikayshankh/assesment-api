package database

import (
	"api/ent"
	"testing"

	_ "github.com/lib/pq"
)

func TestNewEntClient(t *testing.T) {
	type args struct {
		database string
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Client
		wantErr error
	}{
		{
			name: "1st test case",
			args: args{
				database: "employee",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewEntClient(tt.args.database)
			if tt.wantErr != nil {
				t.Errorf("NewEntClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
