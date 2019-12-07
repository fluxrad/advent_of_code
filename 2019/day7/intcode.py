class Intcode:
    def __init__(self, memory):
        self.memory = memory

    def parse_opcode(self, opcode):
        ops = list(map(int, f'{opcode:05}'))
        op = ops[3] * 10 + ops[4] * 1
        mode_map = ops[:3]
        mode_map.reverse()  # reverse it for easier parsing. first mode is [0], second [1], etc.
        return op, mode_map


    # Determine if we're using immediate mode (1) or position mode (0)
    # and return the resulting target locations
    def deref(self, loc, *mode_map):
        mem = self.memory
        targets = []
        for i, m in enumerate(mode_map, start=1):
            t = mem[mem[loc + i]] if m == 0 else mem[loc + i]
            targets.append(t)
        return targets


    def add(self, loc, *mode_map):
        x, y = self.deref(loc, mode_map[0], mode_map[1])
        self.memory[self.memory[loc + 3]] = x + y
        return loc + 4


    def multiply(self, loc, *mode_map):
        x, y = self.deref(loc, mode_map[0], mode_map[1])
        self.memory[self.memory[loc + 3]] = x * y
        return loc + 4


    def save(self, loc, le_input):
        self.memory[self.memory[loc + 1]] = le_input
        return loc + 2


    def output(self, loc, *args):
        self.output_signal = self.memory[self.memory[loc + 1 ]]
        print(f'Test results: {self.output_signal}')
        return loc + 2


    def jump_if_true(self, loc, *mode_map):
        x, y = self.deref(loc, mode_map[0], mode_map[1])
        return y if x != 0 else loc + 3


    def jump_if_false(self, loc, *mode_map):
        x, y = self.deref(loc, mode_map[0], mode_map[1])
        return y if x == 0 else loc + 3


    def less_than(self, loc, *mode_map):
        x, y = self.deref(loc, mode_map[0], mode_map[1])
        self.memory[self.memory[loc + 3]] = 1 if x < y else 0
        return loc + 4


    def equals(self, loc, *mode_map):
        x, y = self.deref(loc, mode_map[0], mode_map[1])
        self.memory[self.memory[loc + 3]] = 1 if x == y else 0
        return loc + 4


    def run(self, *inputs):
        halt = False
        loc = 0
        input_id = 0
        while not halt:
            opcode = self.memory[loc]
            op, mode_map = self.parse_opcode(opcode)

            if op == 99:
                halt = True
            elif op == 3:  # save
                loc = self.OPCODES[3](self, loc, inputs[input_id])
                input_id += 1
            else:
                print('doing stuff')
                loc = self.OPCODES[op](self, loc, *mode_map)
        print(f'output signal is: {self.output_signal}')
        return self.output_signal


    # For whatever reason you have to define these late?
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
