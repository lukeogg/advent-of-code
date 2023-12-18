"""
-- Day 16: The Floor Will Be Lava ---
With the beam of light completely focused somewhere, the reindeer leads you deeper still into the Lava Production Facility. At some point, you realize that the steel facility walls have been replaced with cave, and the doorways are just cave, and the floor is cave, and you're pretty sure this is actually just a giant cave.

Finally, as you approach what must be the heart of the mountain, you see a bright light in a cavern up ahead. There, you discover that the beam of light you so carefully focused is emerging from the cavern wall closest to the facility and pouring all of its energy into a contraption on the opposite side.

Upon closer inspection, the contraption appears to be a flat, two-dimensional square grid containing empty space (.), mirrors (/ and \), and splitters (| and -).

The contraption is aligned so that most of the beam bounces around the grid, but each tile on the grid converts some of the beam's light into heat to melt the rock in the cavern.

You note the layout of the contraption (your puzzle input). For example:
"""

import sys


input = []
with open(sys.argv[1], 'r') as f:
    input = f.read().splitlines()
input = [list(x) for x in input]

R  = len(input)
C = len(input[0])

DR = [-1, 0, 1, 0]
DC = [0, 1, 0, -1]


def step(r, c, d):
    return (r + DR[d], c + DC[d], d)

def eval(er, ec, ed):
    #print (er, ec, ed)
    seen = set()
    seen2 = set()
    Pos = [(er, ec, ed)]
    while True:
        new_pos = []
        if not Pos:
            break
        for r, c, d in Pos:
            if 0 <= r < R and 0 <= c < C :
                seen.add((r, c))

                if (r, c, d) in seen2:
                    continue
                seen2.add((r, c, d))
                
                ch = input[r][c]

                #print(r, c, d, ch)

                if ch == '.':
                    new_pos.append(step(r, c, d))
                elif ch == '/':
                    new_pos.append(step(r, c, {0:1, 1:0, 2:3, 3:2}[d]))
                elif ch == '\\':
                    new_pos.append(step(r, c, {0:3, 1:2, 2:1, 3:0}[d]))
                elif ch == '|':
                    if d in [0, 2]:
                        new_pos.append(step(r, c, d))
                    else:
                        new_pos.append(step(r, c, 0))
                        new_pos.append(step(r, c, 2))
                elif ch == '-':
                    if d in [1, 3]:
                        new_pos.append(step(r, c, d))
                    else:
                        new_pos.append(step(r, c, 1))
                        new_pos.append(step(r, c, 3))
                else:
                    print("Error")
                    assert False
        Pos = new_pos
    return len(seen)

print("Part 1:", eval(0, 0, 1))

ANS = 0 
for r in range(R):
    ANS = max(ANS, eval(r, 0, 1))
    ANS = max(ANS, eval(r, C-1, 3))
for c in range(C):
    ANS = max(ANS, eval(0, c, 2))
    ANS = max(ANS, eval(R-1, c, 0))

print("Part 2:", ANS)

    






