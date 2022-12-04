i = 0

for line in open(0):
    x, y = [[int(i) for i in r.split('-')] for r in line.split(',')]
    if x[0] >= y[0] and x[1] <= y[1] or y[0] >= x[0] and y[1] <= x[1]:
        i += 1

print(i)
