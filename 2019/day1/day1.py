#!/usr/bin/env python3

with open('input', 'r') as f:
    masses = f.read().splitlines()
 
# Determine if our fuel needs any additional fuel
def fuel(x):
    f = x // 3 - 2
    return 0 if f <= 0 else (f + fuel(f))

total = sum(list(map(lambda x : fuel(int(x)), masses)))
print(total)
