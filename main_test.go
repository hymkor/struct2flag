package struct2flag_test

import (
	"flag"
	"os"
	"testing"

	"github.com/hymkor/struct2flag"
)

type ts struct {
	S1 string `flag:"string option(1)"`
	S2 string `flag:"S2,string option(2)"`
	S3 string `flag:"S3,string option(3)"`

	I1 int `flag:"integer option(1)"`
	I2 int `flag:"I2,integer option(2)"`
	I3 int `flag:"I3,integer option(3)"`

	B1 bool `flag:"boolean option(1)"`
	B2 bool `flag:"B2,boolean option(2)"`
	B3 bool `flag:"B3,boolean option(3)"`
}

func TestBind(t *testing.T) {
	ts1 := &ts{}
	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	struct2flag.Bind(flagSet, ts1)
	err := flagSet.Parse(
		[]string{
			"-s1", "foo",
			"-S2", "bar",
			"-i1", "9",
			"-I2", "8",
			"-b1",
			"-B2",
		},
	)

	if err != nil {
		t.Fatal(err.Error())
	}

	if expect := "foo"; ts1.S1 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.S1, expect)
	}
	if expect := "bar"; ts1.S2 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.S2, expect)
	}
	if expect := ""; ts1.S3 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.S3, expect)
	}

	if expect := 9; ts1.I1 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.I1, expect)
	}
	if expect := 8; ts1.I2 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.I2, expect)
	}
	if expect := 0; ts1.I3 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.I3, expect)
	}

	if expect := true; ts1.B1 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.B1, expect)
	}
	if expect := true; ts1.B2 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.B2, expect)
	}
	if expect := false; ts1.B3 != expect {
		t.Fatalf("expect %#v,but %#v", ts1.B3, expect)
	}

	ts1 = &ts{}
	flagSet = flag.NewFlagSet("", flag.ContinueOnError)
	struct2flag.Bind(flagSet, ts1)
	cases := [][]string{
		[]string{"-S1", "should be lower case"},
		[]string{"-s2", "should be upper case"},
		[]string{"-I1", "should be lower case"},
		[]string{"-i2", "should be upper case"},
		[]string{"-B1"},
		[]string{"-b2"},
	}
	stderrSaved := os.Stderr
	devnull, err := os.Create(os.DevNull)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer devnull.Close()
	for _, case1 := range cases {
		os.Stderr = devnull
		err := flagSet.Parse(case1)
		os.Stderr = stderrSaved
		if err == nil {
			t.Fatalf("expect error, but succeeded for %#v", case1)
		}
	}
}
