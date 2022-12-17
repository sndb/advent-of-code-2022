# x = offset from left border
# y = offset from floor
shapes = [
    [(0, 0), (1, 0), (2, 0), (3, 0)],
    [(1, 0), (0, 1), (1, 1), (2, 1), (1, 2)],
    [(0, 0), (1, 0), (2, 0), (2, 1), (2, 2)],
    [(0, 0), (0, 1), (0, 2), (0, 3)],
    [(0, 0), (1, 0), (0, 1), (1, 1)],
]

s = set()
block = 0

def next_block():
    global block
    b = shapes[block % len(shapes)]
    block += 1
    return b

def down(b):
    return [(x, y - 1) for x, y in b]

def left(b):
    return [(x - 1, y) for x, y in b]

def right(b):
    return [(x + 1, y) for x, y in b]

def equiv(b):
    return [(x, y) for x, y in b]

def collides(b):
    for x, y in b:
        if x < 0:
            return True
        if x > 6:
            return True
        if y < 0:
            return True
        if (x, y) in s:
            return True
    return False

def initial(b):
    x = 2
    y = 3
    if len(s) > 0:
        y += max(s, key = lambda cord: cord[1])[1] + 1
    return [(u + x, v + y) for u, v in b]

jet = 0
jets = open(0).read()
jets = jets[:-1]

def next_jet():
    global jet
    j = jets[jet % len(jets)]
    jet += 1
    return j

b = next_block()
b = initial(b)
stopped = 0
while True:
    if stopped == 2022:
        break
    if next_jet() == "<":
        l = left(b)
        if not collides(l):
            b = l
    else:
        r = right(b)
        if not collides(r):
            b = r
    d = down(b)
    if not collides(d):
        b = d
    else:
        for x, y in b:
            s.add((x, y))
        b = next_block()
        b = initial(b)
        stopped += 1

print(1 + max(s, key = lambda c: c[1])[1])
