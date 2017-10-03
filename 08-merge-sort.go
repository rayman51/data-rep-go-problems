package main
import ("fmt")
func main() {
  var n int
  fmt.Println("please enter size of the array:")
  fmt.Scanf("%d\n", &n)
  set(n)
}

func set(n int) {

  a := make([]int, n)
  for i := 0; i < n; i++ {
      fmt.Scan(&a[i])
    }
  fmt.Println(a)
 }
