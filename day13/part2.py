from functools import cmp_to_key

def compare(x, y):
    if type(x) == int and type(y) == int:
        return x - y
    if type(x) == list and type(y) == list:
        if len(x) == 0 or len(y) == 0:
            return len(x) - len(y)
        return compare(x[0], y[0]) or compare(x[1:], y[1:])
    return compare(*([z] if type(z) == int else z for z in (x, y)))

dividers = [[[2]], [[6]]]
packets = [eval(x) for x in open(0).read().split()] + dividers
packets.sort(key=cmp_to_key(compare))

t = 1
for i, p in enumerate(packets):
    if p in dividers:
        t *= i + 1
print(t)
