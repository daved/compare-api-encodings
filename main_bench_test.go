package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/tgreiser/compare-api-encodings/bindata"
	"github.com/tgreiser/compare-api-encodings/lib"
	"labix.org/v2/mgo/bson"
)

var (
	data, bErr = bindata.Asset("photo.jpg")

	item = &lib.PbFile{
		Name: "photo.jpg",
		Data: data,
	}
)

func init() {
	m := flag.Int("size", 0, "size of data in MB (0 = sample data, max = 1024)")
	flag.Parse()

	if *m == 0 {
		return
	}
	if *m > 1024 {
		*m = 1024
	}

	s := *m * 1000000

	rns := rand.New(rand.NewSource(time.Now().Unix()))
	rb := make([]byte, s)
	if _, err := rns.Read(rb); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	item.Data = rb
}

func TestAsset(t *testing.T) {
	if bErr != nil {
		t.Fatal(bErr)
	}
}

func BenchmarkMarshalBSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := bson.Marshal(item); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := json.Marshal(item); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalPB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := proto.Marshal(item); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalXML(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := xml.Marshal(item); err != nil {
			b.Fatal(err)
		}
	}
}
