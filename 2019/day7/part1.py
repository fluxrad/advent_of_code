#!/usr/bin/env python3

from intcode import Intcode

with open('input', 'r') as f:
    memory = list(map(int, f.read().split(',')))

# system_id = int(input("Please enter the system ID: "))

m = memory[:]
ampa = Intcode(m)
ampa.run(3, 0)
