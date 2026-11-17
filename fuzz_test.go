package lenientjson

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"testing"
)

func addCleanSeeds(f *testing.F) {
	for _, in := range [][]byte{
		[]byte("\xef\xbb\xbf{\n  // line\n  \"a\": 1, /* block */\n  \"b\": [2,],\n}\n"),
		[]byte(`{"a":"x y","b":[1,null]}`),
		[]byte(`{'a':'x y','b':[1,null]}`),
		[]byte(`{a:"x y", b : [1,null]}`),
		[]byte(`{ a: 'x y', 'b': [1, null,], // note` + "\n}"),
		[]byte("\xef\xbb\xbf" + `{"a":"x y","b":[1,null]}`),
		utf16Of(`{"a":"x y","b":[1,null]}`, binary.LittleEndian, []byte{0xff, 0xfe}),
		utf16Of(`{"a":"x y","b":[1,null]}`, binary.BigEndian, []byte{0xfe, 0xff}),
		utf32Of(`{"k":[1,2,],}`, binary.LittleEndian, []byte{0xff, 0xfe, 0, 0}),
		utf32Of(`{"a":"x y","b":[1,null]}`, binary.BigEndian, []byte{0, 0, 0xfe, 0xff}),
		[]byte(`{a:"x y",b:[1,undefined]}`),
		[]byte(`{a:"x y",b:[1,-Infinity]}`),
		[]byte(`{/* c */a:"x y",b:[1,NaN]}`),
		[]byte("{\"a\":\"l1\nl2\tx\x01\"}"),
		[]byte(`{'a':'it\'s "q"'}`),
		[]byte(`{"a":"// not /* a */ comment, ]"}`),
		[]byte(`{"a":"b:c, d: e"}`),
		[]byte(`{"a":"\u00e9\n\"\\"}`),
		[]byte(`{"a":"don't"}`),
		[]byte("{/*\n\n*/ \"a\": 1}"),
		[]byte(`{"o":1} /* x`),
		[]byte(`[,]`), []byte(`[1,,2]`), []byte(`[,1]`), []byte(`[1,,]`), []byte(`[1,]`),
		[]byte(`{"o":,"s":"x"}`), []byte(`{"o":,}`),
		[]byte(`{,}`), []byte(`{"o":1,,"s":"x"}`), []byte(`{"o":1,,}`),
		[]byte("{\"a\":\"x\ny\nz\",\n\"b\": }"),
		[]byte(`[.5,-.5,5.,1.5,5.e2]`), []byte(`[+1]`),
		[]byte("{\"a\":\"x\\\ny\"}"), []byte("{\"a\":\"x\\\ty\"}"),
		[]byte(`[0x1F]`), []byte(`[012]`),
		[]byte("{/*\n\n*/\"a\":1,\n\"b\": }"),
		[]byte("{// c\n\"a\":1,\n\"b\": }"),
		utf16Of("{\n\"a\":1,\n\"b\": }", binary.LittleEndian, []byte{0xff, 0xfe}),
		[]byte(`{"a":-NaN}`), []byte(`[-undefined]`), []byte(`[1,-NaN]`),
		[]byte(`{1:2}`),
		[]byte("{a:\u00a01}"),
		[]byte("{//c\r\"a\":1}"),
		[]byte("{a /*c*/ : 1}"),
		[]byte("{a // c\n : 1}"),
		[]byte(`{"a"/*c*/:1}`),
		[]byte("{\"a\":1}\x00\x00"),
		[]byte("{\x00a\x00:\x001}"),
		[]byte("[1,\x00]"),
		[]byte("\"a\x00b\""),
		[]byte("{ä:1}"),
		[]byte("{٣:1}"),
		[]byte("{$_a1:1}"),
		[]byte("{k“:1}"),
		{0xff, 0xfe, '1', 0, '2'},
		{0xfe, 0xff, 0, '1', 0},
		{0xff, 0xfe, 0, 0, '1', 0, 0, 0, '2'},
		{0, 0, 0xfe, 0xff, 0, 0, 0, '1', 0, 0},
		{},
		[]byte(`"`), []byte(`/*`), []byte(`/`), []byte(`{`), []byte(`[`),
	} {
		f.Add(in)
	}
}

func rewriteLimit(in []byte) int {
	return 16 * (len(in) + 8)
}

func FuzzClean(f *testing.F) {
	addCleanSeeds(f)
	f.Fuzz(func(t *testing.T, in []byte) {
		out := Clean(in)
		if n, limit := len(out), rewriteLimit(in); n > limit {
			t.Fatalf("unbounded rewrite: in=%d out=%d max=%d", len(in), n, limit)
		}
		if json.Valid(in) && !bytes.Equal(out, in) {
			t.Fatalf("strict JSON rewritten: in=%q out=%q", in, out)
		}
		if json.Valid(out) && !bytes.Equal(Clean(out), out) {
			t.Fatalf("not idempotent: out=%q again=%q", out, Clean(out))
		}
	})
}

func FuzzTryClean(f *testing.F) {
	addCleanSeeds(f)
	f.Fuzz(func(t *testing.T, in []byte) {
		out, err := TryClean(in)
		if err != nil {
			if !errors.Is(err, ErrTruncated) {
				t.Fatalf("unexpected error %v for %q", err, in)
			}
			if !bytes.Equal(Clean(in), in) {
				t.Fatalf("truncated input rewritten: in=%q out=%q", in, Clean(in))
			}
			return
		}
		if n, limit := len(out), rewriteLimit(in); n > limit {
			t.Fatalf("unbounded rewrite: in=%d out=%d max=%d", len(in), n, limit)
		}
		if !bytes.Equal(out, Clean(in)) {
			t.Fatalf("TryClean/Clean mismatch: try=%q clean=%q", out, Clean(in))
		}
		if json.Valid(in) && !bytes.Equal(out, in) {
			t.Fatalf("strict JSON rewritten: in=%q out=%q", in, out)
		}
		if json.Valid(out) {
			again, err := TryClean(out)
			if err != nil || !bytes.Equal(again, out) {
				t.Fatalf("not idempotent: out=%q again=%q err=%v", out, again, err)
			}
		}
	})
}
