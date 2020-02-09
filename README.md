# replicating-app
An application that replicates itself.

After reading [Ken Thompson's Turing Award Lecture _Reflections on Trusting Trust_](https://www.archive.ece.cmu.edu/~ganger/712.fall02/papers/p761-thompson.pdf), I was inspired to create my own self-reproducing program.

## Execution Steps
1. Load the contents of the currently running file into memory
2. Write the contents into an executable file in the clones folder
3. Execute the new file

## Managing Clones
The self-production of this application can become difficult to manage, so it is programmed to keep clones in a `clones/` directory.

The filename of the file produced in step 2 is the current unix timestamp. Because that value only changes once per second, the number of files written is artificially restricted. Naming strategies could be adopted to remove this restriction which includes choosing to use a number that changes more rapidly (`time.Now().UnixNano()`), always incrementing a number, or appending characters to the end of the filename.

## Purpose
This program has no purpose.
