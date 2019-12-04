#!/usr/bin/env python3

def has_double(num):
    prev = -1
    for i in num:
        if int(i) == prev:
            return True
        prev = int(i)
    return False


def decreases(num):
    prev = -1
    for i in num:
        if int(i) < prev:
            return True
        prev = int(i)
    return False


with open('input', 'r') as f:
    lower, upper = list(map(int, f.read().strip().split('-')))

count = 0
for i in range(upper - lower):
    candidate = str(lower + i)
    if has_double(candidate) and not decreases(candidate):
        count += 1

print(f'The answer is: {count}')
