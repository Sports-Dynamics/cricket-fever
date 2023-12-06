package handlers

import (
	"net/http"
	"testing"

	"github.com/sports-dynamics/cricket-fever/internal/services"
)

func Test_createTeam_ServeHTTP(t *testing.T) {
	type fields struct {
		service services.Team
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := createTeam{
				service: tt.fields.service,
			}
			tr.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}
