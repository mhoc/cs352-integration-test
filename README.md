
# MSP Integration Tests

These tests are designed to test the MSP binary by comparing
output. They are highly accurate tests which look for exact output from
the compiler. It provides actionable feedback when an error occurs.
The selection of test cases is meant to be comprehensive, and as such
multiple different sources are used to create them, including Piazza, the
lab handouts, and other students' public test cases.

# See An Error?

File an issue on the right, open a pull request, or contact me otherwise.

# Easy Setup

These commands download, run, and clean the test files in one swoop...

```
git clone http://github.com/mhoc/cs352-test.git
cd cs352-integration-test && go run main.go ../parser
rm -rf cs352-test
```

Or you can set it up to not have to download each time...

```
git clone http://github.com/mhoc/cs352-test.git test
```

And add to your makefile...

```
test:
  cd test && git pull && go run main.go ../parser
```

Up to you!

# Run Options

`go run main.go path/to/binary`

`go run main.go -exit-on-fail path/to/binary`
This will cause the test suite to exit when it encounters its first failure.
You can use this to decrease the output you need to read.

# Test Files

The testfiles are located in multiple subfolders under the folder `testfiles/`.
Each subfolder is a type of feature of the language and contains both pass
and failure tests which test for proper output from the compiler.

Each test file attempts to test one single feature. That being said, there are
some core features which are impossible to avoid in most tests. These features
are tested in the `testfiles/core/` module and are tested first. They include
things like basic document structure (header/ending tags) and
very basic document.write().

If a testfile has output, it is contained in a similarly named file with a
`.outp` extension.

# Output

In addition to pass/fail, you will receive the time it took to complete the
test in microseconds. If you fail a test, you will receive a printout of
the expected output, your output, and the test case itself. Super handy.
