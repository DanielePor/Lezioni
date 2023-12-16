package main
import "fmt"
import "os"
import "strconv"

func main() {
	var b []int
	b = LeggiVoti()
	fmt.Println(b)
	x,y := FiltraVoti(b)
	fmt.Println(x)
	fmt.Println(y)
}

func LeggiVoti() []int {
	voti:=os.Args[1:]
	a := Conversione(voti)
	return a
}

func Conversione(sl []string) []int{
	f := make([]int, len(sl))
	for i, s:= range sl{
		f[i], _ = strconv.Atoi(s)
	}
	return f
}

func FiltraVoti(voti []int) (suf, ins []int){
	for i:= range voti {
		if voti[i] >=60 {
			suf = append(suf, voti[i])
		} else{
			ins = append(ins, voti[i])
		}
	}
	return
}