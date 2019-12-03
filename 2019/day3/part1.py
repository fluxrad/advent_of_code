#!/usr/bin/env python3

import numpy as np

def place_wire1(grid, wire):
    posX, posY = 0, 0

    for i in wire:
        direction = i[0]
        distance = int(i[1:])

        if direction == 'R':
            for _ in range(distance):
                posX += 1
                grid[posX, posY] = 1
        elif direction == 'L':
            for _ in range(distance):
                posX -= 1
                grid[posX, posY] = 1
        elif direction == 'U':
            for _ in range(distance):
                posY += 1
                grid[posX, posY] = 1
        elif direction == 'D':
            for _ in range(distance):
                posY -= 1
                grid[posX, posY] = 1

    return grid


def place_wire2(grid, wire, cross_coords):
    posX, posY = 0, 0

    for i in wire:
        direction = i[0]
        distance = int(i[1:])

        if direction == 'R':
            for _ in range(distance):
                posX += 1
                if grid[posX, posY] == 1:
                    cross_coords.append([posX, posY])
        elif direction == 'L':
            for _ in range(distance):
                posX -= 1
                if grid[posX, posY] == 1:
                    cross_coords.append([posX, posY])
        elif direction == 'U':
            for _ in range(distance):
                posY += 1
                if grid[posX, posY] == 1:
                    cross_coords.append([posX, posY])
        elif direction == 'D':
            for _ in range(distance):
                posY -= 1
                if grid[posX, posY] == 1:
                    cross_coords.append([posX, posY])

    return grid, cross_coords
# Main


with open('input.example', 'r') as f:
    wires = f.read().splitlines()

wire1 = wires[0].split(',')
wire2 = wires[1].split(',')

# TODO: figure out how to know what the appropriate size grid is based on the
# wires.
grid = np.zeros((1000000, 1000000), int)

cross_coords = []
grid = place_wire1(grid, wire1)
grid, cross_coords = place_wire2(grid, wire2, cross_coords)

print(cross_coords)

min_distance = abs(cross_coords[0][0]) + abs(cross_coords[0][1])
for c in cross_coords:
    distance = abs(c[0]) + abs(c[1])
    min_distance = distance if distance < distance else min_distance

print(f'The answer is {min_distance}')
