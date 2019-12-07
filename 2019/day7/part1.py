#!/usr/bin/env python3

from .intcode import Intcod

with open('input', 'r') as f:
    memory = list(map(int, f.read().split(',')))

system_id = int(input("Please enter the system ID: "))


