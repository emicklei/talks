import "testing"

func TestTos(t *testing.T) {
	for _, each := range []struct {
		input interface{}
		want  string
	}{
		{nil, "<nil>"},
		{"s", "s"},
		{42, "42"},
	} {
		if got := Tos(each.input); got != each.want {
			t.Errorf("got %q want %q", got, each.want)
		}
	}
}
