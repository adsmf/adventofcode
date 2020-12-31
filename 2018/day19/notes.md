```
#ip 3

 0:   addi 3 16 3   # jmp 17
 1:   seti 1 6 1    # r1 = 1
 2:   seti 1 9 5    # r5 = 1
 3:   mulr 1 5 2    # r2 = r1 * r5
 4:   eqrr 2 4 2    # r2 = (r2 == r4)
 5:   addr 2 3 3    # jmp r2+1
 6:   addi 3 1 3    # jmp 2
 7:   addr 1 0 0    # r0 += r1
 8:   addi 5 1 5    # r5++
 9:   gtrr 5 4 2    # r2 = (r5 > r4)
10:   addr 3 2 3    # jmp r2+1
11:   seti 2 4 3    # jmp #3
12:   addi 1 1 1    # r1++
13:   gtrr 1 4 2    # r2 = (r1 > r4)
14:   addr 2 3 3    # jmp r2+1
15:   seti 1 0 3    # jmp 2

16:   mulr 3 3 3    # ip = ip*ip + 1  =>  HALT!

# P1 init
17:   addi 4 2 4    # r4 += 2
18:   mulr 4 4 4    # r4 *= r4
19:   mulr 3 4 4    # r4 *= r3
20:   muli 4 11 4   # r4 *= 11
21:   addi 2 5 2    # r2 += 5
22:   mulr 2 3 2    # r2 *= r3
23:   addi 2 1 2    # r2++
24:   addr 4 2 4    # r4 += r2
25:   addr 3 0 3    # jmp r0+1
26:   seti 0 3 3    # jmp #1   - Only if part 1

# P2 init
27:   setr 3 6 2    # r2 = r3
28:   mulr 2 3 2    # r2 *= r3
29:   addr 3 2 2    # r2 += r3
30:   mulr 3 2 2    # r3 *= r2
31:   muli 2 14 2   # r2 *= 14
32:   mulr 2 3 2    # r2 *= r3
33:   addr 4 2 4    # r4 += r2
34:   seti 0 8 0    # r0 = 0
35:   seti 0 8 3    # jmp #1
```
