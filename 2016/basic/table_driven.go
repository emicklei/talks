import "testing"

func TestTos(t *testing.T) {

	for i, each := range []struct {
		input interface{}
		want  string
	}{
		{nil, "<nil>"},
		{"s", "s"},
		{42, "42"},
	} {
		if got, want := Tos(each.input), each.want; got != want {
			t.Errorf("[%d] got %v want %v", i, got, want)
		}
	}
}
