#!/usr/bin/env python3

def has_double(num):
    prev = -1
    count = 1
    for i in num:
        if int(i) == prev:
            count += 1
        else:
            if count == 2:
                return True
            count = 1
        prev = int(i)

    if count == 2:
        return True
    return False


def decreases(num):
    prev = -1
    for i in num:
        if int(i) < prev:
            return True
        prev = int(i)
    return False


with open('input', 'r') as f:
    lower, upper = map(int, f.read().strip().split('-'))

count = 0
for i in range(upper - lower):
    candidate = str(lower + i)
    if has_double(candidate) and not decreases(candidate):
        print(f'{candidate} is okay')
        count += 1
    else:
        if not decreases(candidate):
            print(f'{candidate} does not decrease but is not okay')

print(f'The answer is: {count}')
