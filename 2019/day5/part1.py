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
    debug(f'Instruction is add: {memory[loc]}, {memory[loc+1]}, {memory[loc+2]}')
    x = memory[memory[loc + 1]] if mode1 == 0 else memory[loc + 1]
    y = memory[memory[loc + 2]] if mode2 == 0 else memory[loc + 2]
    memory[memory[loc + 3]] = x + y
    return loc + 4


def multiply(memory, loc, mode1, mode2, mode3):
    debug(f'Instruction is multiply: {memory[loc]}, {memory[loc+1]}, {memory[loc+2]}')
    x = memory[memory[loc + 1]] if mode1 == 0 else memory[loc + 1]
    y = memory[memory[loc + 2]] if mode2 == 0 else memory[loc + 2]
    memory[memory[loc + 3]] = x * y
    return loc + 4


def save(memory, loc, le_input):
    debug(f'Instruction is save input: {memory[loc]}, {memory[loc+1]}')
    memory[memory[loc + 1]] = le_input
    return loc + 2


def output(memory, loc):
    debug(f'Instruction is print output: {memory[loc]}, {memory[loc+1]}')
    print(f'Test results: {memory[memory[loc + 1]]}')
    return loc + 2


with open('input', 'r') as f:
    memory = list(map(int, f.read().split(',')))

system_id = int(input("Please enter the system ID: "))

halt = False
loc = 0
while not halt:
    opcode = memory[loc]
    mode3, mode2, mode1, op = parse_opcode(opcode)

    debug(f'Operating at location {loc}. opcode is {opcode}. op = {op}, mode1 = {mode1}, mode2 = {mode2}, mode3 = {mode3}')

    if op == 1:
        loc = add(memory, loc, mode1, mode2, mode3)
    elif op == 2:
        loc = multiply(memory, loc, mode1, mode2, mode3)
    elif op == 3:
        loc = save(memory, loc, system_id)
    elif op == 4:
        debug("Checking output:")
        loc = output(memory, loc)
        debug("\n")
    elif op == 99:
        halt = True
    else:
        print("Uah? Invalid opcode")
        halt = True

    debug(f'loc is now {loc}')


# # Initial instructions to replace values
# commands[1] = 12
# commands[2] = 2
#
# for x in range(0, len(commands), 4):
#     cmd = commands[x]
#     if cmd == 99:
#         break
#
#     loc = commands[x+3]
#     if cmd == 1:
#         commands[loc] = commands[commands[x+1]] + commands[commands[x+2]]
#         continue
#     elif cmd == 2:
#         commands[loc] = commands[commands[x+1]] * commands[commands[x+2]]
#         continue
#     elif cmd == 3:
#         commands[commands[x+1]] = 1111
#     elif cmd == 4:
#         print(commands[commands[x+1]])
#
# print('The value of position 0 is', commands[0])

