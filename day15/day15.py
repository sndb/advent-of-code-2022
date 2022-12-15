def parse(f):
    data = []
    for line in f:
        fields = line.split()
        sx = int(fields[2][2:-1])
        sy = int(fields[3][2:-1])
        bx = int(fields[8][2:-1])
        by = int(fields[9][2:])
        data.append(((sx, sy), (bx, by)))
    return data


def distance(p, q):
    return abs(p[0] - q[0]) + abs(p[1] - q[1])


def interval(p, q, y):
    r = (p[0], y)
    d = distance(p, q)
    e = distance(p, r)
    return (p[0] - (d - e), p[0] + (d - e))


def perimeter(p, d):
    e = set()
    for i in range(d + 1):
        j = d - i
        e.add((p[0] + i, p[1] + j))
        e.add((p[0] + i, p[1] - j))
        e.add((p[0] - i, p[1] + j))
        e.add((p[0] - i, p[1] - j))
    return e


def intersect(sb, p):
    return distance(sb[0], sb[1]) >= distance(sb[0], p)


def candidate(p, limit):
    if p[0] < 0 or p[0] > limit or p[1] < 0 or p[1] > limit:
        return False
    for sb in sb_pairs:
        if intersect(sb, p):
            return False
    return True


def part1():
    y = 2000000
    s = set()
    for sb in sb_pairs:
        i = interval(sb[0], sb[1], y)
        for j in range(i[0], i[1] + 1):
            s.add(j)
    return len(s) - 1


def part2():
    limit = 4000000
    s = set()
    for sb in sb_pairs:
        d = distance(sb[0], sb[1])
        for p in perimeter(sb[0], d + 1):
            if candidate(p, limit):
                s.add(p)
    p = s.pop()
    return p[0] * limit + p[1]


sb_pairs = parse(open(0))
print(part1())
print(part2())
