#!/usr/bin/env python3


def parse_opcode(opcode):
    ops = list(map(int, f'{opcode:05}'))
    op = ops[3] * 10 + ops[4] * 1
    mode_map = ops[:3]
    mode_map.reverse()  # reverse it for easier parsing. first mode is [0], second [1], etc.

    return op, mode_map


# Determine if we're using immediate mode (1) or position mode (0)
# and return the resulting target locations
def deref(memory, loc, *mode_map):
    targets = []
    for i, m in enumerate(mode_map, start=1):
        t = memory[memory[loc + i]] if m == 0 else memory[loc + i]
        targets.append(t)
    return targets


def add(memory, loc, *mode_map):
    x, y = deref(memory, loc, mode_map[0], mode_map[1])
    memory[memory[loc + 3]] = x + y
    return loc + 4


def multiply(memory, loc, *mode_map):
    x, y = deref(memory, loc, mode_map[0], mode_map[1])
    memory[memory[loc + 3]] = x * y
    return loc + 4


def save(memory, loc, le_input):
    print(f'memory is {memory[loc:loc+1]}')
    memory[memory[loc + 1]] = le_input
    return loc + 2


def output(memory, loc, *args):
    print(f'Test results: {memory[memory[loc + 1]]}')
    return loc + 2


def jump_if_true(memory, loc, *mode_map):
    x, y = deref(memory, loc, mode_map[0], mode_map[1])
    return y if x != 0 else loc + 3


def jump_if_false(memory, loc, *mode_map):
    x, y = deref(memory, loc, mode_map[0], mode_map[1])
    return y if x == 0 else loc + 3


def less_than(memory, loc, *mode_map):
    x, y = deref(memory, loc, mode_map[0], mode_map[1])
    memory[memory[loc + 3]] = 1 if x < y else 0
    return loc + 4


def equals(memory, loc, *mode_map):
    x, y = deref(memory, loc, mode_map[0], mode_map[1])
    memory[memory[loc + 3]] = 1 if x == y else 0
    return loc + 4


### BEGIN MAIN ###
opcode_map = {
    1: add,
    2: multiply,
    3: save,
    4: output,
    5: jump_if_true,
    6: jump_if_false,
    7: less_than,
    8: equals
}


with open('input', 'r') as f:
    memory = list(map(int, f.read().split(',')))

system_id = int(input("Please enter the system ID: "))

halt = False
loc = 0
while not halt:
    opcode = memory[loc]
    op, mode_map = parse_opcode(opcode)

    if op == 99:
        halt = True
    elif op == 3:  # save
        loc = opcode_map[3](memory, loc, system_id)
    else:
        loc = opcode_map[op](memory, loc, *mode_map)
