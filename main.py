
# TURN BACK ALL YE WHO ENTER HERE
# DONT TRY TO UNDERSTAND HOW IT WORKS, MY GOAL WAS TO GET THE FEWEST LINES
# POSSIBLE IN A SINGLE FILE, NOT FOR IT TO BE READABLE

# =======
# Imports
# =======

import os
from os import listdir
from os.path import isdir, isfile, join
import time
import threading
from subprocess import Popen, PIPE
import sys

# =======
# Globals
# =======

testNo = -1
totalPassed = 0
timeTaken = 0
binaryLocation = ""
tests = []

# ==============
# Util Functions
# ==============

def stripEndNl(st):
    if len(st) > 0 and st[len(st)-1] == '\n':
        return stripEndNl(st[:len(st)-1])
    return st

def getTerminalSize():
    import os
    env = os.environ
    def ioctl_GWINSZ(fd):
        try:
            import fcntl, termios, struct, os
            cr = struct.unpack('hh', fcntl.ioctl(fd, termios.TIOCGWINSZ,
        '1234'))
        except:
            return
        return cr
    cr = ioctl_GWINSZ(0) or ioctl_GWINSZ(1) or ioctl_GWINSZ(2)
    if not cr:
        try:
            fd = os.open(os.ctermid(), os.O_RDONLY)
            cr = ioctl_GWINSZ(fd)
            os.close(fd)
        except:
            pass
    if not cr:
        cr = (env.get('LINES', 25), env.get('COLUMNS', 80))
    return int(cr[1]), int(cr[0])

# ==================
# Printing Functions
# ==================

def green(st):
    sys.stdout.write('\033[92m' + st + '\033[0m')
    sys.stdout.flush()

def red(st):
    sys.stdout.write('\033[91m' + st + '\033[0m')
    sys.stdout.flush()

def pink(st):
    sys.stdout.write('\033[95m' + st + '\033[0m')
    sys.stdout.flush()

def blue(st):
    sys.stdout.write('\033[94m' + st + '\033[0m')
    sys.stdout.flush()

def purple(st):
    sys.stdout.write('\033[1;35m' + st + '\033[0m')
    sys.stdout.flush()

def clearLine():
    (width, height) = getTerminalSize()
    sys.stdout.write('\r')
    for i in xrange(0, width):
        sys.stdout.write(' ')
    sys.stdout.write('\r')
    sys.stdout.flush()

def printHeader(header):
    pink("\033[4m\n=== " + header + " \033[0m")
    for i in xrange(len(header), 75):
        pink("\033[4m=\033[0m")
    print ""

def printTestCase(test):
    f = open(test, "r")
    ln = 1
    for line in f:
        pink("{:0>2d}| ".format(ln))
        sys.stdout.write(line)
        ln += 1
    sys.stdout.flush()

def printSplit(expected, got, splitPt):
    expsp = expected.split("\n")
    gotsp = got.split("\n")
    longestLine = splitPt
    iExp, iGot = 0, 0
    while iExp < len(expsp) or iGot < len(gotsp):
        if iExp < len(expsp):
            sys.stdout.write(expsp[iExp])
            for i in xrange(len(expsp[iExp]), longestLine):
                sys.stdout.write(" ")
            pink("  |")
        else:
            for i in xrange(0, longestLine):
                sys.stdout.write(" ")
            pink("--|")
        if iGot < len(gotsp):
            print " ", gotsp[iGot]
        else:
            pink("--\n")
        iExp += 1
        iGot += 1
    sys.stdout.flush()

def printOutErr(expOut, gotOut, expErr, gotErr):
    m = 23
    for line in expOut.split("\n"):
        if len(line) > m:
            m = len(line)
    for line in expErr.split("\n"):
        if len(line) > m:
            m = len(line)
    printHeader("Stdout === [Expected | Actual]")
    printSplit(expOut, gotOut, m)
    printHeader("Stderr === [Expected | Actual]")
    printSplit(expErr, gotErr, m)

# ==================
# Core Functionality
# ==================

def infiniteLoopTestKiller(process):
    time.sleep(2)
    if process.poll() is None:
        outFile = open("out.temp", "w")
        outFile.write("The test runner detected an infinite loop in your code\n")
        process.kill()

# Runs a single test. Returns true if it passes, otherwise false.
# Printing can be enabled by passing in true for verbosity
def runTest(testfile, verbose=False):
    global testNo, totalPassed, timeTaken
    testNo += 1
    outFile, errFile = open("out.temp", "w+"), open("error.temp", "w+")
    startTime = time.time() * 1000
    process = Popen([binaryLocation, testfile], stdout=outFile, stderr=errFile)
    t = threading.Thread(target=infiniteLoopTestKiller, args=(process, ))
    t.start()
    process.wait()
    timeTaken += time.time()*1000 - startTime
    t._Thread__stop()
    outFile.seek(0)
    errFile.seek(0)
    stdout, stderr = stripEndNl(outFile.read()), stripEndNl(errFile.read())
    os.remove("out.temp")
    os.remove("error.temp")
    expectedOut, expectedError = "", ""
    passed = True
    if verbose:
        printHeader(testfile.split("/")[2].replace("-", " ").title())
        printTestCase(testfile)
    if isfile(testfile + ".outp"):
        f = open(testfile + ".outp")
        expectedOut = stripEndNl(f.read())
    passed = expectedOut == stdout
    if isfile(testfile + ".error"):
        f = open(testfile + ".error")
        expectedError = stripEndNl(f.read())
    passed = passed and expectedError == stderr
    if verbose:
        printOutErr(expectedOut, stdout, expectedError, stderr)
    if verbose and passed:
        green("\n" + u"\u2713" + " Test Passed\n")
    elif verbose and not passed:
        red("\n" + u"\u2717" + " Test Failed\n")
    if passed:
        totalPassed += 1
    return passed

# Runs a single module, with print handling
def runModule(module):
    failed = []
    passedIn, totalIn = 0, 0
    pink(module.split("/")[1].replace("-", " ").title() + "\n|\tPassed 0 of 0 tests")
    cases = [ join(module, f) for f in listdir(module) if isfile(join(module, f)) and not ".outp" in f and not ".error" in f]
    for case in cases:
        totalIn += 1
        if runTest(case):
            passedIn += 1
        else:
            failed.append(testNo)
        clearLine()
        pink("|")
        red("\tPassed {} of {} tests".format(passedIn, totalIn))
        sys.stdout.flush()
    if len(failed) > 0:
        print ""
        pink("|")
        red("\tFailed test cases   ")
        for case in failed:
            red('\033[1m' + str(case) + "\033[0m ")
    else:
        clearLine()
        pink("|")
        green("\tPassed all {} cases".format(totalIn))
    print ""

# Runs every test case in alphabetical order {module}->{test-name}
def runTests():
    modules = [ join("cases", f) for f in listdir("cases") if isdir(join("cases",f)) ]
    for module in modules:
        runModule(module)
    blue("\nPassed:\t")
    purple("{} ({:.0f}%)\n".format(totalPassed, (totalPassed / (testNo+1))*100))
    blue("Failed:\t")
    purple("{}\n".format(testNo-totalPassed+1))
    blue("Total:\t")
    purple("{}\n".format(testNo+1))
    blue("Average time per test: ")
    purple("{:.2f}ms\n".format(timeTaken / (testNo+1)))
    blue("Run ")
    purple("python main.py [binary] [test-no] ")
    blue("to see detailed output about a specific test you failed\n")

# This is used during specific test case running so we know which numbered test case
# the user is trying to access. Otherwise we iterate by module, but it ends up being
# in the same order in both cases.
def loadAllTests():
    modules = [ join("cases", f) for f in listdir("cases") if isdir(join("cases",f)) ]
    for module in modules:
        tests.extend([ join(module, f) for f in listdir(module) if isfile(join(module, f)) and not ".outp" in f and not ".error" in f])

# ===============
# Begin Execution
# ===============

def printHelp():
    print "Usage:"
    print "python main.py [path-to-binary]"
    print "\tRuns all of the test cases against your binary and prints succinct results"
    print "python main.py [path-to-binary] [test-number]"
    print "\tRuns your binary against a single test and prints expanded results"

if len(sys.argv) == 2:
    if sys.argv[1] == '-h' or sys.argv[1] == '--help':
        printHelp()
        sys.exit(0)
    binaryLocation = sys.argv[1]
    print ""
    runTests()
elif len(sys.argv) == 3:
    binaryLocation = sys.argv[1]
    loadAllTests()
    runTest(tests[int(sys.argv[2])], True)
else:
    printHelp()

print ""

if totalPassed != testNo + 1:
    sys.exit(1)
