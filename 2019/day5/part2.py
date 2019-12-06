#!/usr/bin/env python3

DEBUG = False

def debug(s):
    if DEBUG:
        print(s)


def parse_opcode(opcode):
    ops = list(map(int, f'{opcode:05}'))
    op = ops[3] * 10 + ops[4] * 1
    modes = ops[:3]
    return [*modes, op]


def add(memory, loc, mode1, mode2, mode3):
    debug(f'Instruction is add: [{memory[loc]}, {memory[loc+1]}, {memory[loc+2]}]')
    x = memory[memory[loc + 1]] if mode1 == 0 else memory[loc + 1]
    y = memory[memory[loc + 2]] if mode2 == 0 else memory[loc + 2]
    memory[memory[loc + 3]] = x + y
    return loc + 4


def multiply(memory, loc, mode1, mode2, mode3):
    debug(f'Instruction is multiply: [{memory[loc]}, {memory[loc+1]}, {memory[loc+2]}]')
    x = memory[memory[loc + 1]] if mode1 == 0 else memory[loc + 1]
    y = memory[memory[loc + 2]] if mode2 == 0 else memory[loc + 2]
    memory[memory[loc + 3]] = x * y
    return loc + 4


def save(memory, loc, le_input):
    debug(f'Instruction is save input: [{memory[loc]}, {memory[loc+1]}]')
    memory[memory[loc + 1]] = le_input
    return loc + 2


def output(memory, loc):
    debug(f'Instruction is print output: [{memory[loc]}, {memory[loc+1]}]')
    print(f'Test results: {memory[memory[loc + 1]]}')
    return loc + 2


def jump_if_true(memory, loc, mode1, mode2):
    debug(f'Instruction is jump if true: [{memory[loc]}, {memory[loc+1]}, {memory[loc+2]}]')
    x = memory[memory[loc + 1]] if mode1 == 0 else memory[loc + 1]
    y = memory[memory[loc + 2]] if mode2 == 0 else memory[loc + 2]
    return y if x != 0 else loc + 3


def jump_if_false(memory, loc, mode1, mode2):
    debug(f'Instruction is jump if false: [{memory[loc]}, {memory[loc+1]}, {memory[loc+2]}]')
    x = memory[memory[loc + 1]] if mode1 == 0 else memory[loc + 1]
    y = memory[memory[loc + 2]] if mode2 == 0 else memory[loc + 2]
    return y if x == 0 else loc + 3


def less_than(memory, loc, mode1, mode2):
    debug(f'Instruction is less than: [{memory[loc]}, {memory[loc+1]}, {memory[loc+2]}]')
    x = memory[memory[loc + 1]] if mode1 == 0 else memory[loc + 1]
    y = memory[memory[loc + 2]] if mode2 == 0 else memory[loc + 2]
    memory[memory[loc + 3]] = 1 if x < y else 0
    return loc + 4


def equals(memory, loc, mode1, mode2):
    debug(f'Instruction is equals: [{memory[loc]}, {memory[loc+1]}, {memory[loc+2]}]')
    x = memory[memory[loc + 1]] if mode1 == 0 else memory[loc + 1]
    y = memory[memory[loc + 2]] if mode2 == 0 else memory[loc + 2]
    memory[memory[loc + 3]] = 1 if x == y else 0
    return loc + 4


with open('input', 'r') as f:
    memory = list(map(int, f.read().split(',')))
system_id = int(input("Please enter the system ID: "))

halt = False
loc = 0
while not halt:
    opcode = memory[loc]
    mode3, mode2, mode1, op = parse_opcode(opcode)

    if op == 1:
        loc = add(memory, loc, mode1, mode2, mode3)
    elif op == 2:
        loc = multiply(memory, loc, mode1, mode2, mode3)
    elif op == 3:
        loc = save(memory, loc, system_id)
    elif op == 4:
        loc = output(memory, loc)
    elif op == 5:
        loc = jump_if_true(memory, loc, mode1, mode2)
    elif op == 6:
        loc = jump_if_false(memory, loc, mode1, mode2)
    elif op == 7:
        loc = less_than(memory, loc, mode1, mode2)
    elif op == 8:
        loc = equals(memory, loc, mode1, mode2)
    elif op == 99:
        halt = True
    else:
        print(f'Uah? Invalid op: {op}.')
        halt = True

    debug(f'loc is now {loc}')
