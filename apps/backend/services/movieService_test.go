package services_test

import (
	"testing"
	"travisroad/gotracker/di"
	"travisroad/gotracker/services"
	"travisroad/gotracker/utils"
)

func TestGetMovieMetaData(t *testing.T) {
	utils.PreTest()
	type args struct {
		id      int
		options map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				id:      1,
				options: map[string]string{},
			},
			wantErr: true,
		},
	}
	di.C.Invoke(func(ms *services.MovieService) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ms.GetMovieMetaData(tt.args.id, tt.args.options); (err != nil) != tt.wantErr {
					t.Errorf("GetMovieMetaData() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})
}
