def a(n):
    if n == 0: return 0
    elif n == 1: return 10
    elif n == 2: return 20
    elif n == 3: return 30
    elif n > 1000: return n*10
    else: return a(n-1)+10
 
print(a(1000))