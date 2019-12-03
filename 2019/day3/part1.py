#!/usr/bin/env python3


def place_wire(grid, wire):
    index = grid[0][0]
    for i in wire:
        direction = i[0]
        distance = int(i[1:])

## Main
with open('input', 'r') as f:
    wires = f.read().splitlines()

wire1 = wires[0].split(',')
wire2 = wires[1].split(',')

# TODO: figure out how to know what the appropriate size grid is based on the wires.
grid = [[0] * 10000 for i in range(10000)]

grid = place_wire(grid, wire1)
