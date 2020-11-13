def TranslateZahlzuText(Zahl):
    a="ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwqxyz0123456789!?#@$%^&*_+/-=[]{}().,<>"
    i=0
    b=""
    while i<len(str(Zahl)):
        m=str(Zahl)[i]+str(Zahl)[i+1]
        n=int(m)-10
        b=b+a[n]
        i=i+2
    return b


import gmpy2
from sympy.ntheory import factorint

# encrypted message
a = 511959376208947061283695640365354741498
# victim public key
n = 514329716272439139523844134754318536519
# victim public d (or knowingly called E)
d = 401

i,j = factorint(n)
m = (i-1)*(j-1)

p=int(gmpy2.invert(d,m))

x = pow(a,p,n)

msg = TranslateZahlzuText(x)
print(msg)