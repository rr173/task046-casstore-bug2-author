package cas

import "testing"

func TestProbeEmptyGetReturnsNonNilSlice(t *testing.T) {

	s := New()
	h, _, err := s.Put([]byte{})
	if err != nil {
		t.Fatal(err)
	}
	got, err := s.Get(h)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("Get of an empty block must return a non-nil empty slice")
	}
	if len(got) != 0 {
		t.Fatalf("empty Get len=%d, want 0", len(got))
	}
}
