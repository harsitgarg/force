package main

import (
	"fmt"
	"io/ioutil"
	"os"
	
)



var cmdApex = &Command{
	Run:   generateTest,
	Usage: "generateTest",
	Short: "Create Test Class",
	Long: ``,
}

func init() {
}

func generateTest(cmd *Command, args []string) {
	var jsonData []byte
	var generatorClass string
	var sourceClass string
	var err error
	
    if len(args) == 3 {
		generatorClass = args[0]
		sourceClass = args[1]
		jsonData, err = ioutil.ReadFile(args[2])
	}else{
		fmt.Println("Please enter the parameters in the given format ---- generatorClassName sourceClass jsonData")
		return
	}
	
	if err != nil {
		ErrorAndExit(err.Error())
	}
	force, _ := ActiveForce()
	if len(args) == 3 {
		
		err := force.generateTestClass(generatorClass,sourceClass,string(jsonData))
		if err != nil {
			ErrorAndExit(err.Error())
		}
		
	}
		
}