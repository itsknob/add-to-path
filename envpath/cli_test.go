package envpath_test

import (
	"os"
	"testing"

	"github.com/itsknob/hawk-tui/envpath"
)

func TestPath_AddToPathFront(t *testing.T) {
    os.Setenv("PATH", "/usr/bin:/usr/local/bin")
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dir     string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
        struct{name string; dir string; want string; wantErr bool}{
            name: "Success",
            dir: "/Users/stephen.reilly/personal/test/.local/bin",
            want: "/Users/stephen.reilly/personal/test/.local/bin:/usr/bin:/usr/local/bin",
            wantErr: false,
        },
        {
            name: "Dir does not exist",
            dir: "/asdf",
            want: "/usr/bin:/usr/local/bin",
            wantErr: true,
        },

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var p envpath.Path
			got, gotErr := p.AddToPathFront(tt.dir)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("AddToPathFront() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("AddToPathFront() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
            if os.Getenv("PATH") != tt.want {
				t.Errorf("AddToPathFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

