package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/teajunk/aws-sqs-cli/internal/awssqs"
)

var (
	// BaseVersion is a program version, should be updated on a build stage
	BaseVersion string

	flagAwsProfile  = flag.String("profile", "default", "Aws profile to use from ~/.aws/credentials.")
	flagAwsRegion   = flag.String("region", "us-east-1", "Aws region to use.")
	flagAwsQueueURL = flag.String("queue", "", "Aws SQS queue url")
	flagOutputFile  = flag.String("foutput", "", "File to use for output instead of stdout")
	flagHelp        = flag.Bool("help", false, "Prints help message")
	flagVersion     = flag.Bool("version", false, "Prints version")
	filename        = filepath.Base(os.Args[0])
)

func main() {
	flag.Parse()

	if *flagVersion == true {
		fmt.Println("Version: " + BaseVersion)
		return
	}

	if *flagHelp == true || *flagAwsQueueURL == "" {
		fmt.Println("Usage: " + filename + " [-version] [-help] [-profile] [-region] [-queue] command\n")
		flag.PrintDefaults()
		return
	}

	msg := awssqs.New()
	msg.GetSingleMessage(awssqs.NewSession(flagAwsProfile, flagAwsRegion), flagAwsQueueURL)

	if *flagOutputFile != "" {
		if err := ioutil.WriteFile(*flagOutputFile, msg.Body, 0644); err != nil {
			log.Fatalln("Can't create the file specified: ", err)
		}
	}
}
