def s(n):
    if n == 0: return 2
    elif n %2 != 0: return (pow(2,2*n, 10000003)-1%1000003)%1000003
    else: return (pow(2,2*n, 10000003)+1%1000003)%1000003
 
print(s(1000000))