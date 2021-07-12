package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	complexpb "github.com/prot-buff-proj/src/complex"
	enumpb "github.com/prot-buff-proj/src/enum"
	simplepb "github.com/prot-buff-proj/src/simple"
	// "google.golang.org/protobuf/proto"
	// "google.golang.org/protobuf/proto"
	//example_simple "src/simple"
)

func main() {
	// sm := doSimple()

	// readAndwrite(sm)
	//jsonDemo(sm)

	//doEnum()
	doComplex()

}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third Message",
			},
			&complexpb.DummyMessage{
				Id:   4,
				Name: "Fourth Message",
			},
			&complexpb.DummyMessage{
				Id:   5,
				Name: "Fifth Message",
			},
			&complexpb.DummyMessage{
				Id:   6,
				Name: "Sixth Message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayofTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayofTheWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	msg := toJson(sm)
	fmt.Println(msg)

	sm2 := &simplepb.SimpleMessage{}
	fromJson(msg, sm2)
	fmt.Println("Successfully created proto struct:\n", sm2)
}

func toJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJson(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb) // from JSON format to proto struct SimpleMessage type, 'in' is json format
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into pb struct ", err)
	}
}

func readAndwrite(sm proto.Message) {

	writeToFile("simple.bin", sm) // passing filename and msg
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb) //encode msg to wire format and return byte data

	if err != nil {
		log.Fatalln("Can't serialise to bytes ", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		//ioutil.WriteFile(file name, data []byte, permission)
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")

	return nil

}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't read from file ", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb) //decodes file from wire byte format
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into protocol buffers struct ", err2)
		return err2
	}

	return nil

}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Shivesh",
		SampleList: []int32{1, 4, 7, 8},
	}
	// fmt.Println(sm)

	// sm.Name = "Ojha"
	// fmt.Println(sm)
	// fmt.Println("The ID is : ", sm.Id)

	return &sm

}
