// internal/routes/routes.go

package routes

import (
	"reflect"
	"skylab-book-chameleon/internal/database"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetupRouter(t *testing.T) {
	type args struct {
		db *database.DB
	}
	tests := []struct {
		name string
		args args
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetupRouter(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetupRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
