package main

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/tgreiser/compare-api-encodings/bindata"
	"github.com/tgreiser/compare-api-encodings/lib"
	"labix.org/v2/mgo/bson"
)

var (
	data, _ = bindata.Asset("photo.jpg")

	item = &lib.PbFile{
		Name: "photo.jpg",
		Data: data,
	}
)

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
