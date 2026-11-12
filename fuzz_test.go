package lenientjson

import (
	"encoding/binary"
	"testing"
)

func FuzzClean(f *testing.F) {
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
		{},
		[]byte(`"`), []byte(`/*`), []byte(`/`), []byte(`{`), []byte(`[`),
	} {
		f.Add(in)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		out := Clean(in)
		if n, limit := len(out), 16*(len(in)+8); n > limit {
			t.Fatalf("unbounded rewrite: in=%d out=%d max=%d", len(in), n, limit)
		}
	})
}
