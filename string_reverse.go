package main
import "fmt"

func main(){
s:="vishwanath"
l:=len(s)
r:=[]rune(s)
for i,j:=0,l-1;i<j;i,j=i+1,j-1{
r[i],r[j]=r[j],r[i]
}
fmt.Println(string(r))
}
