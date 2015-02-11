# cs352-test-cases

Test cases and a test case runner for the 352 compiler project.

## Warning

I won't be liable if this accidentally deletes your stuff. Specfiically, running this without 
arguments invokes `clean_cmd` in the python script, which deletes files. It shouldn't be important
files, but I can't guarantee that. Also, running `python maker.py makefile` generates a new makefile
for you, overwriting your old one. Be aware. Use git.

## Complaints About Test Cases?

Open an issue or submit a pull request. I seriously appreciate any contributions anyone wants to make.

## Setup

1. Copy maker.py and the test/ folder into the root directory of your project.

2. Open up maker.py

3. Make sure the commands specified in `build_graph` are correct. Namely that the files specified are of the correct name. 

4. Make sure the clean command is good. Should be the same as a `make clean`

5. Make sure the `command` variable in `test_parameters` is the same name as your binary (aka the -o in gcc in `build_graph` above)

## Using

To build your project: `python maker.py build`

To test your project: `python maker.py test`

To clean your project `python maker.py clean`

To do all three, in that order: `python maker.py`

To generate a makefile based upon the buildgraph and clean command specified in maker.py: `python maker.py makefile`
