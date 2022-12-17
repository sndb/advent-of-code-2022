blocks = [ # (left border offset, floor offset)
    [(0, 0), (1, 0), (2, 0), (3, 0)],
    [(1, 0), (0, 1), (1, 1), (2, 1), (1, 2)],
    [(0, 0), (1, 0), (2, 0), (2, 1), (2, 2)],
    [(0, 0), (0, 1), (0, 2), (0, 3)],
    [(0, 0), (1, 0), (0, 1), (1, 1)],
]
jets = open(0).read()[:-1]
block = jet = 0

def next_block():
    global block
    b = blocks[block]
    block = (block + 1) % len(blocks)
    return b

def next_jet():
    global jet
    j = jets[jet]
    jet = (jet + 1) % len(jets)
    return j

def down(b):
    return [(x, y - 1) for x, y in b]

def left(b):
    return [(x - 1, y) for x, y in b]

def right(b):
    return [(x + 1, y) for x, y in b]

def collides(b):
    for x, y in b:
        if x < 0 or x > 6 or y < 0 or (x, y) in state:
            return True
    return False

def initial():
    u, v = 2, 3
    if len(state) > 0:
        v += height()
    return [(x + u, y + v) for x, y in next_block()]

def height():
    return 1 + max(state, key=lambda c: c[1])[1]

state = set()
b = initial()

seen = set()
r_want = 1000000000000
i = h = h_delta = h_prev = r = r_delta = r_prev = r_down = 0

while True:
    if (block, jet) in seen:
        seen = set()
        i += 1

        h = height()
        h_delta = h - h_prev
        h_prev = h

        r = r_down
        r_delta = r - r_prev
        r_prev = r

    seen.add((block, jet))

    if i == 2:
        m = (r_want - r) // r_delta
        if r_down == r_want - (m * r_delta):
            print(height() + m * h_delta)
            break

    x = left(b) if next_jet() == "<" else right(b)
    if not collides(x):
        b = x

    x = down(b)
    if not collides(x):
        b = x
    else:
        for x, y in b:
            state.add((x, y))
        b = initial()
        r_down += 1
