#!/usr/bin/env python3

with open('input', 'r') as f:
    commands = f.read().split(',')

commands = [ int(x) for x in commands ]

## Initial instructions to replace values
commands[1] = 12
commands[2] = 2

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

print('The value of position 0 is', commands[0])
