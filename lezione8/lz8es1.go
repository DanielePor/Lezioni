package main
import "fmt" 
import "os"
import "strconv"

func main(){
	b:=os.Args[1:]
	a:=Conversione(b)
	fmt.Println(Calcola(a))
}

func Conversione(sl []string) []int{
	f := make([]int, len(sl))
	for i, s:= range sl{
		f[i], _ = strconv.Atoi(s)
	}
	return f
}

func Calcola(sl []int) int{
	var s, p int
	for i:= range sl{
		for j:= i+1; j<len(sl); j++{
			p = sl[i]*sl[j]
			if p%2 == 0{
				s=s+p
			}
		}
	}
	return s;
}
