tt=int(input())
for t in range(tt):
    n,k=[int(i)for i in input().split()]
    s=input()
    if "*"*k in s:
        print("YES")
    else:
        print("NO")
