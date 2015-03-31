
package main

import (
  "crypto/md5"
  "encoding/hex"
  "flag"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "os/exec"
  "strings"
  "text/tabwriter"
  "time"
)

var (
  TestCases []*TestCase
  TabWriter *tabwriter.Writer
  BinaryLocation string
  ExitOnFail bool

  TestFiles = []string{
    "testfiles/core/",
    "testfiles/variables/",
    "testfiles/expressions/",
    "testfiles/objects/",
  }
)

type TestCase struct {
  FileName string
  PrettyName string
  Content string
  ExpectedOutput string
  ActualOutput string
  Id string
  Passed bool
  Time int64
}

func main() {

  // Set a flag for exiting on first fail with an error code
  // This is necessary for my CI system
  flag.BoolVar(&ExitOnFail, "exit-on-fail", false, "Exit and return a failure code when the first test fails")
  flag.Parse()

  // Init the tab writer
  TabWriter = new(tabwriter.Writer)
  TabWriter.Init(os.Stdout, 0, 8, 1, '\t', 0)

  // Init the test cases
  TestCases = make([]*TestCase, 0, 0)

  // Get the location of the binary
  if (flag.NArg() != 1) {
    panic("Must provide path to binary")
  } else {
    BinaryLocation = flag.Arg(0)
  }

  // Open the directory of test cases
  for _, directory := range TestFiles {
    files, err := ioutil.ReadDir(directory)
    Check(err)

    // Create a test case for each file
    for _, file := range files {
      if strings.Contains(file.Name(), ".outp") {
        continue
      }
      filename := directory + file.Name()
      TestCases = append(TestCases, CreateTest(filename))
    }
  }

  // Run and print the test results
  for _, test := range TestCases {
    RunTest(test)
    PrintTestResult(test)
  }

  TabWriter.Flush()

}

func CreateTest(filename string) *TestCase {

  // Create a pretty name
  sp := strings.Split(filename, "/")
  name := sp[len(sp)-1]
  pretty := strings.Title(strings.Replace(name, "-", " ", -1))

  // Read the content of the test case
  contentb, err := ioutil.ReadFile(filename)
  Check(err)
  content := string(contentb)

  // Read the expected output
  expectedb, err := ioutil.ReadFile(filename + ".outp")
  expected := ""
  if err == nil {
    expected = string(expectedb)
  }

  // Remove trailing newlines from the expected output
  content = StripTabs(StripEndNewline(content))
  expected = StripTabs(StripEndNewline(expected))

  // Generate an id
  md5h := md5.New()
  io.WriteString(md5h, content)
  id := hex.EncodeToString(md5h.Sum(nil))[:4]

  // Create the test case
  testCase := &TestCase{
    FileName: filename,
    Content: content,
    ExpectedOutput: expected,
    Id: id,
    PrettyName: pretty,
  }

  return testCase

}

func RunTest(test *TestCase) {

  // Record the before time
  before := time.Now()

  // Run the test
  result, err := exec.Command(BinaryLocation, test.FileName).CombinedOutput()
  test.Time = time.Since(before).Nanoseconds()
  Check(err)

  // Check the output against the expected
  test.ActualOutput = StripTabs(StripEndNewline(string(result)))
  test.Passed = test.ActualOutput == test.ExpectedOutput

}

func PrintTestResult(test *TestCase) {
  if (test.Passed) {
    PrintTestPass(test)
  } else {
    PrintTestFail(test)
  }
}

// ==========
// Test Printing
// ==========

func PrintTestPass(test *TestCase) {
  line := fmt.Sprintf("%s\t%s\t%d us\n", FormatGreen(test.Id), test.PrettyName, test.Time / 1000)
  fmt.Fprintf(TabWriter, line)
}

func PrintTestFail(test *TestCase) {
  line := fmt.Sprintf("%s\t%s\t%d us\n", FormatLightRed(test.Id), test.PrettyName, test.Time / 1000)
  fmt.Fprintf(TabWriter, line)
  fmt.Fprintf(TabWriter, FormatLightRed("==== Expected ==========================") + "\n")
  fmt.Fprintf(TabWriter, test.ExpectedOutput + "\n")
  fmt.Fprintf(TabWriter, FormatLightRed("==== Output ============================") + "\n")
  fmt.Fprintf(TabWriter, test.ActualOutput + "\n")
  fmt.Fprintf(TabWriter, FormatLightRed("==== Test Case =========================") + "\n")
  fmt.Fprintf(TabWriter, test.Content + "\n")
  fmt.Fprintf(TabWriter, FormatLightRed("========================================") + "\n\n")
  if (ExitOnFail) {
    TabWriter.Flush()
    fmt.Println("Test failure caught. Exiting and reporting error.")
    os.Exit(1)
  }
}

// ==========
// Utility Functions
// ==========

func Check(er error) {
  if er != nil {
    panic(er)
  }
}

func StripEndNewline(s string) string {
  if len(s) > 0 && s[len(s)-1] == '\n' {
    return s[:len(s)-1]
  } else {
    return s
  }
}

func StripTabs(s string) string {
  return strings.Replace(s, "\t", "", -1)
}

// Formats a string to be colored red
func FormatRed(s string) string {
  return "\033[0;31m" + s + "\033[0;00m"
}

// Formats a string light red
func FormatLightRed(s string) string {
  return "\033[1;31m" + s + "\033[0;00m"
}

// Formats a string to be colored yellow
func FormatYellow(s string) string {
  return "\033[1;33m" + s + "\033[0;00m"
}

// Formats a string to be colored green
func FormatGreen(s string) string {
  return "\033[0;32m" + s + "\033[0;00m"
}

// Formats a string to be colored cyan
func FormatCyan(s string) string {
  return "\033[1;36m" + s + "\033[0;00m"
}

func FormatWhite(s string) string {
  return "\033[0;00m" + s + "\033[0;00m"
}
