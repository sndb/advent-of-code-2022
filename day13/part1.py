def compare(x, y):
    if type(x) == int and type(y) == int:
        return x - y
    if type(x) == list and type(y) == list:
        if len(x) == 0 or len(y) == 0:
            return len(x) - len(y)
        return compare(x[0], y[0]) or compare(x[1:], y[1:])
    return compare(*([z] if type(z) == int else z for z in (x, y)))

pairs = [[eval(y) for y in x.splitlines()] for x in open(0).read().split("\n\n")]

t = 0
for i, p in enumerate(pairs):
    if compare(p[0], p[1]) < 0:
        t += 1 + i
print(t)
