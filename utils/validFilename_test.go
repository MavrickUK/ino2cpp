package utils

import "testing"

func TestRemoveInvalidFilenameChars(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{fn: "file?name*with<>invalid|chars"},
			want: "file_name_with_invalid_chars",
		},
		{
			name: "Test case 2",
			args: args{fn: "filename_with_no_invalid_chars"},
			want: "filename_with_no_invalid_chars",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveInvalidFilenameChars(tt.args.fn); got != tt.want {
				t.Errorf("RemoveInvalidFilenameChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
