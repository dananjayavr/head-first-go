package main

import "testing"

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{
			list: []string{},
			want: "",
		},
		{
			list: []string{"apple"},
			want: "apple",
		},
		{
			list: []string{"apple", "orange"},
			want: "apple and orange",
		},
		{
			list: []string{"apple", "orange", "pear"},
			want: "apple, orange, and pear",
		},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\" want \"%s\"", test.list, got, test.want)
		}
	}
}

func TestOneElement(t *testing.T)  {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		errorString(t, list, got, want)
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		errorString(t, list, got, want)
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\" want \"%s\"", list, got, want)
	}
}

func errorString(t *testing.T, list []string, got string, want string) {
	t.Errorf("JoinWithCommas(%#v) = \"%s\" want \"%s\"", list, got, want)
}
