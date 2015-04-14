
# Miniscript Test Suite

This is a test suite for running a series of tests designed for the miniscript
parser. These tests are compiled from a variety of sources, including the
lab handouts, instructor piazza notes, other students' test cases (thanks
Logan!) and elsewhere.

I am convinced that this is the single most advanced test runner and most
comprehensive test suite available for the parser, and it is updated every day?

![Running All The Tests](http://zippy.gfycat.com/OddballNaughtyHydra.gif)

Yes, you're reading that right. At the time of writing, this test runner has over 400 test cases which cover all functionality up to and not including Part 4. Part 4 functionality will be added within the week. 

![Seeing Output For A Test That Failed](http://i.imgur.com/QTT9qtQ.png)

Expected output is on the left and the output from your parser is on the right. Output is correctly separated by stdout and stderr. 

# Add a test case? See an error?

File an issue on the right, open a pull request, or contact me otherwise.

# Easy Setup

You can add the following commands to your makefile as you wish to download
the latest files, run them, and clean up in one swoop.

```
git clone http://github.com/mhoc/cs352-test.git
cd cs352-test && python main.py ../parser
rm -rf cs352-test
```

Or you can do a more complicated setup so you dont have to download it each
time. Up to you.

# Run Options

`python main.py [path-to-binary]`

Runs all of the tests and presents concise error output for which ones you've
failed.

`python main.py [path-to-binary] [test-number]`

Runs a single test case and presents highly detailed output about what the
test case looks like, your output, and the expected output.

# Test Files

The testfiles are located in multiple subfolders under the folder `cases/`.

Each subfolder is viewed as a module and attempts to test one feature or
contains tests from one source. For example, one module might be to just
test array features, whereas another might be the test cases provided
by the class TAs for grading after they are done grading. I've got them
all.

There *will* be a lot of testing overlap between modules and cases. Its
inevitable.

Also note that for the official test cases, I updated them to comply with
the changing project specifications such that you should get 100% with your
final build in Project 4, and (most likely) the binary you turned in for,
say, part 1 would not pass the part-1-official tests in this repository.

Completely unmodified copies of the official test runner and cases is provided
for your convenience under `official/`. These is not used in the test runner
in any way.

# Test Case Comments

* Logans Test 125 from project 3 is confirmed to be invalid due to the usage
of `!(x - y)` on line 21. It is included as a test here but I have modified
the expected output to be empty and expected error to be a `syntax error`

* Part 2 Official: I removed test labeled "19.js19.js" as it appears to
be a repeat of 18.js (or, p2-ta-18 on my runner)

* Part 1 Official: Provided test cases had no expected output
so I filled it in myself.

* Part 1 Official: Test10 is provided in the targz and under `official/` but
is withdrawn from the testrunner as per piazza post 99

# Final Note

Getting 100% on these test cases does not mean you will get 100% on the
project. Failing some of these test cases does not mean you wont get
100% on the project.
