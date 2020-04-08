package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("matched", func(t *testing.T) {
		dictionary := Dictionary{"test" :"this is just a test", "test1":"this is just test1"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("not matched", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test", "test1": "this is just test1"}
		_, err := dictionary.Search("test2")
		want := "not found"
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertString(t, err.Error(), want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want : %s got : %s", got, want)
	}
}