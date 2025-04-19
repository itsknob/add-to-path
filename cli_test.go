package main

import "testing"


func Test_addToPathFront(t *testing.T) {
    t.Setenv("PATH", "/usr/bin:/usr/local/bin")
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dir     string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
        { 
            name: "Succeeds", 
            dir: "/Users/stephen.reilly/.local/bin",
            want: "/usr/bin:/usr/local/bin:/Users/stephen.reilly/.local/bin",
            wantErr: false,
        },
        { 
            name: "Fails when path does not exist", 
            dir: "/Users/stephen.reilly/.local/doesnotexist",
            want: "/usr/bin:/usr/local/bin",
            wantErr: true,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := addToPathFront(tt.dir)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("addToPathFront() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("addToPathFront() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
            if got != tt.want {
				t.Errorf("addToPathFront() = %v, want %v", got, tt.want)
            }
		})
	}
}

