from functools import reduce
tt=int(input())
for t in range(tt):
    time=100/reduce(lambda x,y: x*y,[float(i)for i in input().split()])
    time=round(time,2)
    if time < 9.58:
        print("YES")
    else:
        print("NO")
