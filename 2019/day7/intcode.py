#!/usr/bin/env python3

class Intcode(in):
    OPCODES = {
        1: add,
        2: multiply,
        3: save,
        4: output,
        5: jump_if_true,
        6: jump_if_false,
        7: less_than,
        8: equals
    }

    def __init__(self, memory, phase):
        self.phase = phase
        self.memory = memory

    def parse_opcode(opcode):
        ops = list(map(int, f'{opcode:05}'))
        op = ops[3] * 10 + ops[4] * 1
        mode_map = ops[:3]
        mode_map.reverse()  # reverse it for easier parsing. first mode is [0], second [1], etc.
        return op, mode_map


    # Determine if we're using immediate mode (1) or position mode (0)
    # and return the resulting target locations
    def deref(self, loc, *mode_map):
        targets = []
        for i, m in enumerate(mode_map, start=1):
            t = self.memory[self.memory[loc + i]] if m == 0 else self.memory[loc + i]
            targets.append(t)
        return targets


    def add(self, loc, *mode_map):
        x, y = self.deref(self, loc, mode_map[0], mode_map[1])
        self.memory[self.memory[loc + 3]] = x + y
        return loc + 4


    def multiply(self, loc, *mode_map):
        x, y = self.deref(self, loc, mode_map[0], mode_map[1])
        self.memory[self.memory[loc + 3]] = x * y
        return loc + 4


    def save(self, loc, le_input):
        print(f'memory is {self.memory[loc:loc+1]}')
        self.memory[self.memory[loc + 1]] = le_input
        return loc + 2


    def output(self, loc, *args):
        print(f'Test results: {self.memory[self.memory[loc + 1]]}')
        return loc + 2


    def jump_if_true(self, loc, *mode_map):
        x, y = self.deref(self, loc, mode_map[0], mode_map[1])
        return y if x != 0 else loc + 3


    def jump_if_false(self, loc, *mode_map):
        x, y = self.deref(self, loc, mode_map[0], mode_map[1])
        return y if x == 0 else loc + 3


    def less_than(self, loc, *mode_map):
        x, y = self.deref(self, loc, mode_map[0], mode_map[1])
        self.memory[self.memory[loc + 3]] = 1 if x < y else 0
        return loc + 4


    def equals(self, loc, *mode_map):
        x, y = self.deref(self, loc, mode_map[0], mode_map[1])
        self.memory[self.memory[loc + 3]] = 1 if x == y else 0
        return loc + 4

    def run(phase):
        halt = False
        loc = 0
        while not halt:
            opcode = self.memory[loc]
            op, mode_map = self.parse_opcode(opcode)

            if op == 99:
                halt = True
            elif op == 3:  # save
                loc = self.OPCODES[3](self, loc, system_id)
            else:
                loc = self.OPCODES[op](self, loc, *mode_map)
        return 0
