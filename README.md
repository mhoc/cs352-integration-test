# cs352-test-cases

A completely comprehensive set of test cases for the 252 compiler.
You're looking at other sets of test cases and thinking "these are comprehensive".
No. They aren't. This is the only comprehensive set, guaranteed.

And look at that, it also comes with a python program to run them. Wow.

## Using maker.py

maker.py is a single python script that replaces your entire make process. What? 
How much are you charging for this? $50? $100? No, its free. You're welcome.

Look at the `build_graph` dictionary at the top of maker.py. This defines a graph
which is your build process. Theres a start point `__start`, then each level has
dependencies which are met by another level in the graph. Neat, huh?

Make sure the files in the commands in that graph are named correctly. And that
the commands are what you want.

After that, look at the `clean_command` and make sure that is good.

After that...

`python maker.py build` or `python maker.py clean`

## Using the testing library 

Copy the `test` folder and its content into the same directory as `maker.py` (which
should be the same directory as your project, of course).

Then, just run `python maker.py test`

You'll get colorized output of the 31+ tests currently in this suite. If a test passes,
you get a green check. If it fails, a red X and the output of the compile command after
it. 

Some tests are marked as optional. This would probably be because myself and a few
friends believe that this should be the way miniscript works, but we haven't
confirmed it with the professor. These tests will be marked in yellow if you
fail them.

For convenience, running `python maker.py` with no arguments will do all three steps.
Build, test, and then clean. 

Enjoy!
