l=(input())
n=len(l)
sum=0
check=False
for i in(l[n-2::-2]):
  sum=sum+int(i)
x=sum*2
sum2=0
for i in (l[n-1::-2]):
  sum2=sum2+int(i)
y=sum2
if x+y==20:
  check=True
else:
  check=False
print(check)
print(x,y)
if check==True:
    if(l[0:2])==34 or 37:
        print("ae")
    if(l[0:2])in(51,52,53,54,55):
        print("b")
    if(l[0])==4:
        print("c")
