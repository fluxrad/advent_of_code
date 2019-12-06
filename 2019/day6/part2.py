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

    def path(self):
        if self.orbiting is None:
            return []
        return self.orbiting.path() + [self.orbiting.name]


def distance(a, b):
    patha = a.path()
    pathb = b.path()
    return len(set(patha) ^ set(pathb))


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

print(distance(orbit_map['SAN'], orbit_map['YOU']))
