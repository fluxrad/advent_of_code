#!/usr/bin/env python3
import sys

def calculate_output(commands, noun, verb):
    commands[1] = noun
    commands[2] = verb

    for x in range(0, len(commands), 4):
        cmd = commands[x]
        if cmd == 99:
            break

        loc = commands[x+3]
        if cmd == 1:
            commands[loc] = commands[commands[x+1]] + commands[commands[x+2]]
            continue
        elif cmd == 2:
            commands[loc] = commands[commands[x+1]] * commands[commands[x+2]]
            continue

    return commands[0]

TARGET_OUTPUT = 19690720

with open('input', 'r') as f:
    command_list = [int(x) for x in f.read().split(',')]

print(f'Looking for target output: {TARGET_OUTPUT}')

for noun in range(100):
    for verb in range(100):
        output = calculate_output(list(command_list), noun, verb)
        if output == TARGET_OUTPUT:
            print(f'Found noun:{noun} and verb:{verb}')
            print('The answer is:', 100 * noun + verb)
            sys.exit()

