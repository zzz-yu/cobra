package cobra

import "testing"

func TestExecuteCReportsUnknownCommandForNonRunnable(t *testing.T) {
	root := &Command{Use: "app"}
	root.AddCommand(&Command{Use: "foo"})
	root.SetArgs([]string{"foo", "unknown"})

	_, err := root.ExecuteC()
	if err == nil {
		t.Fatal("expected an unknown command error")
	}
	if got, want := err.Error(), `unknown command "unknown" for "app foo"`; got != want {
		t.Fatalf("unexpected error: got %q, want %q", got, want)
	}
}
