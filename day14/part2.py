s = set()
floor = 0

for line in [
    [tuple(int(x) for x in d.split(",")) for d in line.split(" -> ")]
    for line in open(0)
]:
    for (x0, y0), (x1, y1) in zip(line, line[1:]):
        x0, x1 = sorted([x0, x1])
        y0, y1 = sorted([y0, y1])
        for x in range(x0, x1 + 1):
            for y in range(y0, y1 + 1):
                floor = max(floor, y + 2)
                s.add((x, y))

def step(x, y):
    if y + 1 == floor:
        return (x, y)
    if (x, y + 1) not in s:
        return (x, y + 1)
    if (x - 1, y + 1) not in s:
        return (x - 1, y + 1)
    if (x + 1, y + 1) not in s:
        return (x + 1, y + 1)
    return (x, y)

j = 0
x, y = x0, y0 = 500, 0
while (x0, y0) not in s:
    px, py = x, y
    x, y = step(x, y)
    if (px, py) == (x, y):
        s.add((x, y))
        j += 1
        x, y = x0, y0

print(j)
