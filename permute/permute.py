#!/usr/bin/env python3
# Generate a random permutation of a list of n integers and
# output them as a list and string

import sys
import random

nargs = len(sys.argv)
if  nargs != 2:
    print(f"Usage: {sys.argv[0]} n")
    exit(1)

n = int(sys.argv[1])

# Generate function
def generatePermutation(n):
    intList = [x + 1 for x in range(n)]
    random.shuffle(intList)
    intStr = ' '.join(str(x) for x in intList)
    print(intStr)
    
if __name__ == "__main__":
    generatePermutation(n)