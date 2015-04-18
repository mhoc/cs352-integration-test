
# Miniscript Test Suite

This is a test suite for running a series of tests designed for the miniscript
parser. These tests are compiled from a variety of sources, including the
lab handouts, instructor piazza notes, other students' test cases (thanks
Logan!) and elsewhere.

![Running All The Tests](http://zippy.gfycat.com/OddballNaughtyHydra.gif)

At the time of writing, this test runner has over 400 test cases which cover all functionality up to and not including Part 4. Part 4 functionality will be added within the week.

![Seeing Output For A Test That Failed](http://i.imgur.com/QTT9qtQ.png)

Expected output is on the left and the output from your parser is on the right. Output is correctly separated by stdout and stderr.

# Add a test case? See an error?

File an issue on the right, open a pull request, or contact me otherwise.

# Setup

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

# Known "Bugs"

* The runner employes a rudimentary form of infinite loop detection, though which if any test takes longer than 2 seconds, it assumes that its an infinite loop and you fail the test. Normally these tests should never take more than 5 milliseconds, so if you hit one of these bugs on a test you know should be passing, you can raise the timeout in the main.py file and also re-evaluate some core design patterns in your parser because jesus man, 2 seconds?

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

* Made some minor modifications to error output on Logan's (part 2) test cases
to account for the change in the part 3 handout:

> A read of an undeclared variable prior to any write to that variable returns an undefined type, and is considered a value error.

# Final Note

Getting 100% on these test cases does not mean you will get 100% on the
project. Failing some of these test cases does not mean you wont get
100% on the project.
