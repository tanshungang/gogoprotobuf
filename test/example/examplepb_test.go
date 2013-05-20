// Code generated by protoc-gen-gogo.
// source: example.proto
// DO NOT EDIT!

package test

import testing "testing"
import math_rand "math/rand"
import time "time"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"
import testing1 "testing"
import math_rand1 "math/rand"
import time1 "time"
import encoding_json "encoding/json"
import testing2 "testing"
import math_rand2 "math/rand"
import time2 "time"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"
import fmt "fmt"
import math_rand3 "math/rand"
import time3 "time"
import testing3 "testing"
import math_rand4 "math/rand"
import time4 "time"
import testing4 "testing"
import fmt1 "fmt"
import math_rand5 "math/rand"
import time5 "time"
import testing5 "testing"
import math_rand6 "math/rand"
import time6 "time"
import testing6 "testing"
import fmt2 "fmt"
import go_parser "go/parser"
import math_rand7 "math/rand"
import time7 "time"
import testing7 "testing"
import math_rand8 "math/rand"
import time8 "time"
import testing8 "testing"
import code_google_com_p_gogoprotobuf_proto2 "code.google.com/p/gogoprotobuf/proto"
import math_rand9 "math/rand"
import time9 "time"
import testing9 "testing"

func TestAProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkAProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := code_google_com_p_gogoprotobuf_proto.Marshal(p); err != nil {
			panic(err)
		}
	}
}

func BenchmarkAProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
			panic(err)
		}
	}
}

func TestBProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkBProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := code_google_com_p_gogoprotobuf_proto.Marshal(p); err != nil {
			panic(err)
		}
	}
}

func BenchmarkBProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
			panic(err)
		}
	}
}

func TestUProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkUProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := code_google_com_p_gogoprotobuf_proto.Marshal(p); err != nil {
			panic(err)
		}
	}
}

func BenchmarkUProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
			panic(err)
		}
	}
}

func TestAJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedA(popr)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func BenchmarkAJSONMarshal(b *testing1.B) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := encoding_json.Marshal(p); err != nil {
			panic(err)
		}
	}
}

func BenchmarkAJSONUnmarshal(b *testing1.B) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := encoding_json.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if err := encoding_json.Unmarshal(data, msg); err != nil {
			panic(err)
		}
	}
}

func TestBJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedB(popr)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func BenchmarkBJSONMarshal(b *testing1.B) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := encoding_json.Marshal(p); err != nil {
			panic(err)
		}
	}
}

func BenchmarkBJSONUnmarshal(b *testing1.B) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := encoding_json.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if err := encoding_json.Unmarshal(data, msg); err != nil {
			panic(err)
		}
	}
}

func TestUJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedU(popr)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func BenchmarkUJSONMarshal(b *testing1.B) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := encoding_json.Marshal(p); err != nil {
			panic(err)
		}
	}
}

func BenchmarkUJSONUnmarshal(b *testing1.B) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := encoding_json.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if err := encoding_json.Unmarshal(data, msg); err != nil {
			panic(err)
		}
	}
}

func TestAProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestAProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkAProtoTextMarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	}
}

func BenchmarkAProtoCompactTextMarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	}
}

func BenchmarkAProtoTextUnmarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		msg.Reset()
		if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
			panic(fmt.Sprintf("%v given %v", err, data))
		}
	}
}

func TestBProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestBProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkBProtoTextMarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	}
}

func BenchmarkBProtoCompactTextMarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	}
}

func BenchmarkBProtoTextUnmarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		msg.Reset()
		if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
			panic(fmt.Sprintf("%v given %v", err, data))
		}
	}
}

func TestUProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestUProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkUProtoTextMarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	}
}

func BenchmarkUProtoCompactTextMarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	}
}

func BenchmarkUProtoTextUnmarshal(b *testing2.B) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		msg.Reset()
		if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
			panic(fmt.Sprintf("%v given %v", err, data))
		}
	}
}

func TestUUnion(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedU(popr)
	v := p.GetValue()
	msg := &U{}
	if !msg.SetValue(v) {
		t.Fatalf("Union: Could not set Value")
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Union Equal %#v", msg, p)
	}
}
func TestAStringGen(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedA(popr)
	s1 := p.String()
	s2 := fmt1.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func BenchmarkAStringGen(b *testing4.B) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data := p.String()
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.String()
	}
}

func TestBStringGen(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedB(popr)
	s1 := p.String()
	s2 := fmt1.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func BenchmarkBStringGen(b *testing4.B) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data := p.String()
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.String()
	}
}

func TestUStringGen(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedU(popr)
	s1 := p.String()
	s2 := fmt1.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func BenchmarkUStringGen(b *testing4.B) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data := p.String()
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.String()
	}
}

func BenchmarkAPopulate(b *testing5.B) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewPopulatedA(popr)
	}
}

func BenchmarkBPopulate(b *testing5.B) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewPopulatedB(popr)
	}
}

func BenchmarkUPopulate(b *testing5.B) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewPopulatedU(popr)
	}
}

func TestAGoString(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedA(popr)
	s1 := p.GoString()
	s2 := fmt2.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func BenchmarkAGoString(b *testing6.B) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data := p.GoString()
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.GoString()
	}
}

func TestBGoString(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedB(popr)
	s1 := p.GoString()
	s2 := fmt2.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func BenchmarkBGoString(b *testing6.B) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data := p.GoString()
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.GoString()
	}
}

func TestUGoString(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedU(popr)
	s1 := p.GoString()
	s2 := fmt2.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func BenchmarkUGoString(b *testing6.B) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data := p.GoString()
	b.SetBytes(int64(len(data)))
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.GoString()
	}
}

func TestAFace(t *testing7.T) {
	popr := math_rand7.New(math_rand7.NewSource(time7.Now().UnixNano()))
	p := NewPopulatedA(popr)
	msg := p.TestProto()
	if !p.Equal(msg) {
		t.Fatalf("%#v !Face Equal %#v", msg, p)
	}
}
func BenchmarkAFace(b *testing7.B) {
	popr := math_rand7.New(math_rand7.NewSource(time7.Now().UnixNano()))
	p := NewPopulatedA(popr)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.TestProto()
	}
}

func TestAVerboseEqual(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedA(popr)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto2.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func BenchmarkAEqual(b *testing8.B) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedA(popr)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.Equal(p)
	}
}

func TestBVerboseEqual(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedB(popr)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &B{}
	if err := code_google_com_p_gogoprotobuf_proto2.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func BenchmarkBEqual(b *testing8.B) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedB(popr)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.Equal(p)
	}
}

func TestUVerboseEqual(t *testing8.T) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedU(popr)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &U{}
	if err := code_google_com_p_gogoprotobuf_proto2.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func BenchmarkUEqual(b *testing8.B) {
	popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
	p := NewPopulatedU(popr)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.Equal(p)
	}
}

func TestBDescription(t *testing9.T) {
	popr := math_rand9.New(math_rand9.NewSource(time9.Now().UnixNano()))
	p := NewPopulatedB(popr)
	p.Description()
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
