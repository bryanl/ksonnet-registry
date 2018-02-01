package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ksonnet/ksonnet/metadata/parts"

	"github.com/sirupsen/logrus"
)

func main() {
	var in string
	flag.StringVar(&in, "in", "", "parts.yaml to upgrade")
	flag.Parse()

	if in == "" {
		flag.Usage()
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(in)
	if err != nil {
		logrus.WithError(err).Fatal("read file")
	}

	spec, err := parts.Unmarshal(b)
	if err != nil {
		logrus.WithError(err).Fatal("unmarshal part")
	}

	if spec.APIVersion != "0.0.1" {
		logrus.Info("parts.yaml doesn't need to be upgraded")
	}

	spec.APIVersion = "0.2.0"

	b, err = spec.Marshal()
	if err != nil {
		logrus.WithError(err).Fatal("marshal part")
	}

	fmt.Println(string(b))
}
