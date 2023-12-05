def s(n):
    if n == 0: return 1
    elif n == 1: return -4
    elif n == 2: return -2*s(2-1)-s(2-2)
    else: return -2*s(n-1)-s(n-2)
 
print(s(100))