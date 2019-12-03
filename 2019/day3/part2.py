#!/usr/bin/env python3

import numpy as np


def place_wire1(grid, wire):
    posX, posY = 0, 0
    steps = 0

    for i in wire:
        direction = i[0]
        distance = int(i[1:])

        if direction == 'R':
            for _ in range(distance):
                steps += 1
                posX += 1
                if grid[posX, posY] == 0:
                    grid[posX, posY] = steps
        elif direction == 'L':
            for _ in range(distance):
                steps += 1
                posX -= 1
                if grid[posX, posY] == 0:
                    grid[posX, posY] = steps
        elif direction == 'U':
            for _ in range(distance):
                steps += 1
                posY += 1
                if grid[posX, posY] == 0:
                    grid[posX, posY] = steps
        elif direction == 'D':
            for _ in range(distance):
                steps += 1
                posY -= 1
                if grid[posX, posY] == 0:
                    grid[posX, posY] = steps

    return grid


def place_wire2(grid, wire):
    steps = 0
    min_steps = 10000000000000  # whatever
    posX, posY = 0, 0

    for i in wire:
        direction = i[0]
        distance = int(i[1:])

        if direction == 'R':
            for _ in range(distance):
                steps +=1
                posX += 1
                if grid[posX, posY] > 0 and steps + grid[posX, posY] < min_steps:
                    min_steps = steps + grid[posX, posY]
        elif direction == 'L':
            for _ in range(distance):
                steps +=1
                posX -= 1
                if grid[posX, posY] > 0 and steps + grid[posX, posY] < min_steps:
                    min_steps = steps + grid[posX, posY]
        elif direction == 'U':
            for _ in range(distance):
                steps +=1
                posY += 1
                if grid[posX, posY] > 0 and steps + grid[posX, posY] < min_steps:
                    min_steps = steps + grid[posX, posY]
        elif direction == 'D':
            for _ in range(distance):
                steps +=1
                posY -= 1
                if grid[posX, posY] > 0 and steps + grid[posX, posY] < min_steps:
                    min_steps = steps + grid[posX, posY]

    return grid, min_steps
# Main


with open('input', 'r') as f:
    wires = f.read().splitlines()

wire1 = wires[0].split(',')
wire2 = wires[1].split(',')

# TODO: figure out how to know what the appropriate size grid is based on the
# wires.
grid = np.zeros((1000000, 1000000), int)

grid = place_wire1(grid, wire1)
grid, min_steps = place_wire2(grid, wire2)

print(f'The answer is {min_steps}')
