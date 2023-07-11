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
		source  string
		options map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "err",
			args: args{
				id:      1,
				source:  "douban",
				options: map[string]string{},
			},
			wantErr: true,
		},
		{
			name: "spider man",
			args: args{
				id:      569094,
				source:  "tmdb",
				options: map[string]string{},
			},
			wantErr: false,
		},
	}
	di.C.Invoke(func(ms *services.MovieService) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := ms.GetMovieMetaData(tt.args.id, tt.args.source, tt.args.options)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetMovieMetaData() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})
	utils.PostTest()
}
