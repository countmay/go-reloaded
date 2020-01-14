package student

func Atoi(s string) int {
a:=[]rune(s)
l:=0
c:=1
p:=0
p2:=0
for range a{
	l++
}
for i:=0; i<l; i++{
if a[i]>='0'&&a[i]<='9'|| a[i]=='+'&&i==0||a[i]=='-'&&i==0{
if a[i]=='+'{
   continue
}
if a[i]=='-'{
	c=-1
	continue
}
p=int(a[i])-48+p*10
if p<p2{
	return 0
}
p2=p

} else{
	return 0
}

}
return p*c
}
