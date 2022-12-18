cubes = [tuple(int(x) for x in s.split(",")) for s in open(0).readlines()]

def points(x, y, z):
    return [
        (x + 1, y + 1, z + 1),
        (x + 1, y + 1, z),
        (x + 1, y, z + 1),
        (x + 1, y, z),
        (x, y + 1, z + 1),
        (x, y + 1, z),
        (x, y, z + 1),
        (x, y, z),
    ]

def sides(x, y, z):
    p = points(x, y, z)
    front = (p[7], p[3], p[5], p[1])
    back = (p[6], p[2], p[4], p[0])
    top = (p[5], p[1], p[4], p[0])
    bottom = (p[7], p[3], p[6], p[2])
    left = (p[7], p[6], p[5], p[4])
    right = (p[3], p[2], p[1], p[0])
    return sorted((front, back, top, bottom, left, right))

all_sides = set()
covered = set()
for c in cubes:
    s = sides(*c)
    for s in sides(*c):
        if s in all_sides:
            covered.add(s)
        all_sides.add(s)

print(len(all_sides) - len(covered))
