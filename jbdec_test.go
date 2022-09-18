package jbdec_test

import (
	"testing"

	"github.com/shu-go/gotwant"
	"github.com/shu-go/jbdec"
)

func TestComponent(t *testing.T) {
	d := jbdec.New([]byte(`{}[]:,null true false -12.34e5 "hoge{\a\"\b"`))
	tok := d.Next()
	gotwant.Test(t, tok.Type, jbdec.BeginObject, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.EndObject, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.BeginArray, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.EndArray, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.ValueSeparator, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.Null, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.True, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.False, gotwant.Format("%x"))
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.Number, gotwant.Format("%x"))
	gotwant.Test(t, tok.String(), "-12.34e5")
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
	gotwant.Test(t, tok.String(), `"hoge{\a\"\b"`)
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.EOF, gotwant.Format("%x"))
}

func TestError(t *testing.T) {
	d := jbdec.New([]byte(`abc`))
	tok := d.Next()
	gotwant.Test(t, tok.Type, jbdec.Error, gotwant.Format("%x"))
	gotwant.TestError(t, tok.Error(), `invalid`)
	tok = d.Next()
	gotwant.Test(t, tok.Type, jbdec.Error, gotwant.Format("%x"))

}

func TestObject(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		d := jbdec.New([]byte(`{
  "b": {
    "v": "b-v1",
    "w": "b-v2"
  },
  "1": {
    "v": "a-v1",
    "w": "a-v2"
  }
}`))
		tok := d.Next()
		gotwant.Test(t, tok.Type, jbdec.BeginObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"b"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.BeginObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"v"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"b-v1"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.ValueSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"w"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"b-v2"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.EndObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.ValueSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"1"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.BeginObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"v"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"a-v1"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.ValueSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"w"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"a-v2"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.EndObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.EndObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.EOF, gotwant.Format("%x"))
	})

	t.Run("2", func(t *testing.T) {
		d := jbdec.New([]byte(`{
  "2": {
    "sub3": "san"
  },
  "1": {
    "sub1": "ichi",
    "sub2": "ni"
  }
}`))
		tok := d.Next()
		gotwant.Test(t, tok.Type, jbdec.BeginObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"2"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.BeginObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"sub3"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"san"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.EndObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.ValueSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"1"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.BeginObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"sub1"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"ichi"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.ValueSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"sub2"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.NameSeparator, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.String, gotwant.Format("%x"))
		gotwant.Test(t, tok.String(), `"ni"`, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.EndObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.EndObject, gotwant.Format("%x"))
		tok = d.Next()
		gotwant.Test(t, tok.Type, jbdec.EOF, gotwant.Format("%x"))
	})
}

func BenchmarkDecode(b *testing.B) {
	data := `{
  "2": {
    "sub3": "san"
  },
  "1": {
    "sub1": "ichi",
    "sub2": "ni"
  }
}`

	input := []byte(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := jbdec.New(input)
		_ = d.Next()
		_ = d.Next()
		_ = d.Next()
		_ = d.Next()
		_ = d.Next()
	}
}
