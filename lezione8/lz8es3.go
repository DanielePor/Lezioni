package main
import  (
		"fmt"
		"os"
 		"strconv"
		"math/rand"
		"time"
)
func main(){
	s:=LeggiSoglia()
	fmt.Println(s)
	b:=Genera(s)
	fmt.Println(b)
}

func LeggiSoglia() (f int){
	b:=os.Args[1:]
	f, _ = strconv.Atoi(b[0])
	return
}
func Genera(soglia int) []int{
	rand.Seed(int64(time.Now().Nanosecond())) 
	var a []int
	i := rand.Intn(100)
	for i>=soglia{
		a = append (a , i)
		i = rand.Intn(100)
	}
	return a
}