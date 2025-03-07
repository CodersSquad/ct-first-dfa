package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DFA struct {
	InitialState string   `yaml:"S"`
	States       []string `yaml:"K"`
	Alphabet     []string `yaml:"E"`
	Transitions  []string `yaml:"T"`
	FinalStates  []string `yaml:"F"`
}

type Result struct {
	Input  string   `yaml:"input"`
	States []string `yaml:"states"`
	Valid  bool     `yaml:"valid"`
}

func openInputFiles(dfaFile, inputFile string) (dfa DFA, strings []string) {

	yamlFile, err := os.ReadFile(dfaFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &dfa)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	stringsFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer stringsFile.Close()

	scanner := bufio.NewScanner(stringsFile)
	for scanner.Scan() {
		line := scanner.Text()
		strings = append(strings, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func processData(dfa DFA, strings []string) (results []Result) {

	log.Printf("Using the following DFA \n %#v\n", dfa)

	log.Printf("Let's start processing %v strings", len(strings))
	// Your code goes here, below some sample results addition
	results = append(results, Result{"aba", []string{"q1", "q2"}, false})
	results = append(results, Result{"abba", []string{"q1", "q1", "q3", "q2"}, true})

	return
}

func main() {
	var dfa = flag.String("dfa", "dfa.yaml", "The DFA machine definition file")
	var input = flag.String("input", "strings.txt", "Input list of strings file")
	var output = flag.String("output", "results.yaml", "Yaml-formatted results file")

	flag.Parse()

	dfaDef, strings := openInputFiles(*dfa, *input)

	results := processData(dfaDef, strings)

	yamlData, err := yaml.Marshal(&results)
	if err != nil {
		log.Fatalf("Error marshaling YAML: %v", err)
	}

	err = os.WriteFile(*output, yamlData, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}
