
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
  "text/tabwriter"
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
    s := string(bd)
    if len(s) > 0 && s[len(s)-1] == '\n' {
      return s[:len(s)-1]
    } else {
      return s
    }
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
var tabWriter *tabwriter.Writer

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

  // We eventually print out the results with a tabwriter
  tabWriter = new(tabwriter.Writer)
  tabWriter.Init(os.Stdout, 0, 8, 0, '\t', 0)

  // Create the channel which will return the results of each test
  results = make(chan *TestCase)

  // Run each of the test suites
  initTests(&TestSuite{DirPath: "testfiles/syntax-good/", ExpectOutput: false})
  initTests(&TestSuite{DirPath: "testfiles/syntax-bad/", ExpectOutput: true})

  // Print the results
  go printResults()

  // Block the main thread until we are finished
  wg.Wait()

  // Print out the results
  tabWriter.Flush()
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
  if len(tc.Content) > 0 && tc.Content[len(tc.Content)-1] == '\n' {
    tc.Content = tc.Content[:len(tc.Content)-1]
  }

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
  if len(tc.Output) > 0 && tc.Output[len(tc.Output)-1] == '\n' {
    tc.Output = tc.Output[:len(tc.Output)-1]
  }

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
  fmt.Fprintf(tabWriter, " ")
  printGreen(tc.Id)
  fmt.Fprintf(tabWriter, "\t" + tc.PrettyName() + "\t%d us\n", tc.Time / 1000)
}

func printFailure(tc *TestCase) {
  fmt.Fprintf(tabWriter, " ")
  printRed(tc.Id)
  fmt.Fprintf(tabWriter, "\t" + tc.PrettyName() + "\n")
  printFailureHeader("Expected")
  fmt.Fprintf(tabWriter, tc.Expected + "\n")
  printFailureHeader("Output")
  fmt.Fprintf(tabWriter, tc.Output + "\n")
  printFailureHeader("Test Case")
  fmt.Fprintf(tabWriter, tc.Content + "\n")
  printFailureFooter()
  fmt.Println("\n")
}

func printRed(s string) {
  fmt.Fprintf(tabWriter, "\033[0;31m" + s + "\033[0;00m")
}

func printGreen(s string) {
  fmt.Fprintf(tabWriter, "\033[0;32m" + s + "\033[0;00m")
}

func printCyan(s string) {
  fmt.Fprintf(tabWriter, "\033[1;36m" + s + "\033[0;00m")
}

func printFailureHeader(s string) {
  printCyan("==== " + s + " ")
  for i := 0; i < 30-len(s); i++ {
    printCyan("=")
  }
  fmt.Fprintf(tabWriter, "\n")
}

func printFailureFooter() {
  printCyan("====================================\n") // 36
}
