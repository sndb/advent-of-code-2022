import itertools

cubes = set([tuple(int(x) for x in s.split(",")) for s in open(0).readlines()])

def points(cube):
    x, y, z = cube
    return [
        (x + u, y + v, z + w)
        for u, v, w in itertools.product(*((0, 1) for _ in range(3)))
    ]

def sides(cube):
    a, e, b, f, c, g, d, h = points(cube)
    front, back = (a, b, c, d), (e, f, g, h)
    top, bottom = (c, d, g, h), (a, b, e, f)
    left, right = (a, c, e, g), (b, d, f, h)
    return sorted((front, back, top, bottom, left, right))

def near(cube):
    x, y, z = cube
    return [
        (x + 1, y, z),
        (x - 1, y, z),
        (x, y + 1, z),
        (x, y - 1, z),
        (x, y, z + 1),
        (x, y, z - 1),
    ]

def inside(c, a, b):
    return (
        c[0] >= a[0]
        and c[0] <= b[0]
        and c[1] >= a[1]
        and c[1] <= b[1]
        and c[2] >= a[2]
        and c[2] <= b[2]
    )

min_a = max_b = 0
for i in range(3):
    min_a = min(min_a, min(cubes, key=lambda c: c[i])[i])
    max_b = max(max_b, max(cubes, key=lambda c: c[i])[i])

a = (min_a - 1,) * 3
b = (max_b + 1,) * 3
fill = set([a, b])
while True:
    fill_new = set()
    for cube in fill:
        for c in near(cube):
            if inside(c, a, b) and c not in cubes:
                fill_new.add(c)
    fill_prev_len = len(fill)
    fill |= fill_new
    if len(fill) == fill_prev_len:
        break

all_sides = set()
covered = set()
for c in fill:
    s = sides(c)
    for s in sides(c):
        if s in all_sides:
            covered.add(s)
        all_sides.add(s)

area = 6 * (b[0] - a[0] + 1) ** 2
print(len(all_sides) - len(covered) - area)
