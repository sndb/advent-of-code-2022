s = []
f = open(0)

for line in f:
    if line[0] == " ":
        break
    s.append(line[1::4])

s = [[x for x in t if x != " "][::-1] for t in zip(*s)]

f.readline()

for line in f:
    q, src, dst = [int(x) for x in line.split(" ")[1::2]]
    s[dst - 1].extend(s[src - 1][-q:])
    s[src - 1] = s[src - 1][:-q]

print("".join([t[-1] for t in s]))
