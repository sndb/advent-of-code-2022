i = 0

for line in open(0):
    x, y = [[int(i) for i in r.split('-')] for r in line.split(',')]
    # # more concise but slower in 1.5
    # if set(range(x[0], x[1] + 1)) & set(range(y[0], y[1] + 1)):
    #     i += 1
    if x[0] >= y[0]:
        x, y = y, x
    if y[0] - x[0] <= x[1] - x[0]:
        i += 1

print(i)
