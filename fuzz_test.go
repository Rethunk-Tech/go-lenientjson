package lenientjson

import (
	"encoding/json"
	"reflect"
	"testing"
)

// FuzzClean checks that Clean never panics and that whatever it turns into valid JSON stays valid when cleaned again.
func FuzzClean(f *testing.F) {
	for _, s := range []string{
		"\xef\xbb\xbf{'a': 'b', c: [1,], /* x */ // y\n}",
		"\xff\xfe{\x00}\x00",
		"{a:\"\n\"}", "{'a':'\\", "/*", "{a:-Infinity}", "[NaN,undefined,]", "{\"a\\", "\xfe\xff\x00",
	} {
		f.Add([]byte(s))
	}
	f.Fuzz(func(t *testing.T, b []byte) {
		once := Clean(b)
		if json.Valid(once) && !json.Valid(Clean(once)) {
			t.Fatalf("cleaning valid output %q broke it", once)
		}
	})
}

// FuzzStrictUnchanged checks that strict JSON passes through Clean to the same value.
func FuzzStrictUnchanged(f *testing.F) {
	for _, s := range []string{`{"a":[1,2.5,"x\n",null,true]}`, `"//"`, `[1e5,-0]`, `{"a b":{"c":[]}}`} {
		f.Add([]byte(s))
	}
	f.Fuzz(func(t *testing.T, b []byte) {
		var want any
		if json.Unmarshal(b, &want) != nil {
			return
		}
		var got any
		if err := json.Unmarshal(Clean(b), &got); err != nil || !reflect.DeepEqual(got, want) {
			t.Fatalf("Clean(%q) = %q (%v), want value %v", b, Clean(b), err, want)
		}
	})
}
