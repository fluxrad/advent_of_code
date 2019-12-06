#!/usr/bin/env python3

class CelestialObject:
    def __init__(self, name):
        self.name = name
        self.orbiting = None

    def orbits(self, target):
        self.orbiting = target

    def checksum(self):
        if self.orbiting is None:
            return 0
        return self.orbiting.checksum() + 1

with open('input', 'r') as f:
    lines = f.read().splitlines()

orbit_map = dict()
for o in lines:
    a, b = o.split(')')
    if a not in orbit_map:
        orbit_map[a] = CelestialObject(a)
    if b not in orbit_map:
        orbit_map[b] = CelestialObject(b)
    orbit_map[b].orbits(orbit_map[a])

total_orbits = 0
for n, o in orbit_map.items():
    total_orbits += o.checksum()

print(f'The answer is: {total_orbits}')
