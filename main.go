
package main

import (
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "os/exec"
  "strings"
  "sync"
  "time"
)

type TestSuite struct {
  DirPath string
  ExpectOutput bool
}

/* Returns the expected output of a given file inside the test suite */
func (t *TestSuite) ExpectedOutput(file string) string {
  if t.ExpectOutput {
    bd, err := ioutil.ReadFile(file + ".outp")
    check(err)
    return string(bd)
  } else {
    return ""
  }
}

type TestCase struct {
  Path string
  Id string
  Time int64
  Content string
  Output string
  Expected string
}

func (t *TestCase) PrettyName() string {
  sp := strings.Split(t.Path, "/")
  nm := sp[len(sp)-1]
  return strings.Title(strings.Replace(nm, "-", " ", -1))
}

var wg sync.WaitGroup
var results chan *TestCase
var parserLoc string

/** Checks and panics an error if it is not null */
func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  // Get the location of the binary
  if len(os.Args) != 2 {
    panic("Must provide location of binary as argument")
  }
  parserLoc = os.Args[1]

  // Open the binary file to ensure it exists
  _, err := os.Open(parserLoc)
  check(err)

  // Create the channel which will return the results of each test
  results = make(chan *TestCase)

  // Run each of the test suites
  initTests(&TestSuite{DirPath: "testfiles/syntax-good/", ExpectOutput: false})
  initTests(&TestSuite{DirPath: "testfiles/syntax-bad/", ExpectOutput: true})

  // Print the results
  go printResults()

  // Block the main thread until we are finished
  wg.Wait()
}

/** Initializes and starts all of the tests in a given directory */
func initTests(t *TestSuite) {

  // Open the directory
  files, err := ioutil.ReadDir(t.DirPath)
  check(err)

  // For each file in the directory
  for _, file := range files {

    name := t.DirPath + file.Name()

    // Ignore the output files if they exist
    if strings.Contains(name, ".outp") {
      continue
    }

    // Create the test param object
    tp := TestCase{Path: name, Expected: t.ExpectedOutput(name)}

    // Add one to the waitgroup and spin off a test goroutine
    wg.Add(1)
    go runTest(&tp)

  }
}

/** Concurrently runs a single test and returns the results through the results channel */
func runTest(tc *TestCase) {

  // Get the content of the test case
  testCaseBody, err := ioutil.ReadFile(tc.Path)
  check(err)
  tc.Content = string(testCaseBody)

  // Generate an id
  h := md5.New()
  io.WriteString(h, tc.Content)
  tc.Id = hex.EncodeToString(h.Sum(nil))[:4]

  // Execute the parser with the given test parameter
  before := time.Now()
  result, err := exec.Command(parserLoc, tc.Path).Output()
  tc.Time = time.Since(before).Nanoseconds()
  check(err)

  // Store the result
  tc.Output = string(result)

  results <- tc

}

func printResults() {
  for tc := range results {

    // Check the result against what we expect
    if tc.Output == tc.Expected {
      printSuccess(tc)
    } else {
      printFailure(tc)
    }

    wg.Done()

  }
}

func printSuccess(tc *TestCase) {
  fmt.Print(" ")
  printGreen(tc.Id)
  fmt.Print("\t" + tc.PrettyName() + "\n")
}

func printFailure(tc *TestCase) {
  fmt.Print(" ")
  printRed(tc.Id)
  fmt.Print("\t" + tc.PrettyName() + "\n")
}

func printRed(s string) {
  fmt.Print("\033[0;31m" + s + "\033[0;00m")
}

func printGreen(s string) {
  fmt.Print("\033[0;32m" + s + "\033[0;00m")
}
