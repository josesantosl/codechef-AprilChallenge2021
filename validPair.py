socks=input().split()
if len(socks)==len(set(socks)):
    print("NO")
else:
    print("YES")
