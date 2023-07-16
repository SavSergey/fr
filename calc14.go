package main


import (
  "fmt"
  "strconv"
  "strings"
  )



func get_keys_a(s string) (z string, a int, b int) {
  f := strconv.Atoi
  if s[:2] == "10" {
    z = s[2:3]
    a = int(10)
    b, _ = f(s[3:])
  } else {
    z = s[1:2]
    a, _ = f(s[0:1])
    b, _ = f(s[2:])
  }
  return 
}



func get_keys_r(s string) (z string, a int, b int) {
  roma_array := [10]string{"I", "II", "III", "IV","V", "VI", "VII", "VIII", "IX", "X"}
  lol := 0
  i := 0
  for i < 10 {
    lol = len(roma_array[i])
    if roma_array[i] == s[:lol] {
      z = s[lol:(lol+1)]
      if (z == "+") || (z == "-") || (z == "*") || (z == "/") {
        a = i+1
        lol++
        i = 0
        for i < 10 {
          if roma_array[i] == s[lol:] {
            b = i+1
          }
          i++
        }
      }
    }
    i++
  }
  return
}



func processing(z string, a, b int) (res int) {
  msg := "unknown expression"
  switch z {
    case "+":
      res = a+b
    case "-":
      res = a-b
    case "*":
      res = a*b
    case "/":
      res = a/b
    default:
      fmt.Println(msg)
  }
  return 
}



func convert_a_r(a int) (z string) {
  for 0 < a {
    switch {
      case 999 < a:
        a -= 1000
        z += "M" 
      case 499 < a:
        a -= 500
        z += "D" 
      case 99 < a:
        a -= 100
        z += "C" 
      case 49 < a:
        a -= 50
        z += "L" 
      case 9 < a:
        a -= 10
        z += "X" 
      case 9 == a:
        a -= 9
        z += "IX" 
      case 4 < a:
        a -= 5
        z += "V" 
      case 4 == a:
        a -= 4
        z += "IV" 
      case 0 < a:
        a -= 1
        z += "I" 
    }
  }
  //fmt.Println(z)
  return 
}



func main() {
  var (
    z string
    a, b int
    )
  
  s := "IX*IX"
  //s := "IX*IV"
  //s := "VIII*III"
  //s := "2*2"
  
  fmt.Scan(s)
  s = strings.ToUpper(s)
  fmt.Println(s)
  
  
  flag := false
  if flag {
    z, a, b = get_keys_a(s)
    //fmt.Println(z, a, b)
  } else {
    z, a, b = get_keys_r(s)
    //fmt.Println(z, a, b)
  }
  

  res := processing(z, a, b)
  if flag {
    fmt.Println(s,"=", res)
  } else {
    ans := ""
    ans = convert_a_r(res)
    fmt.Println(s,"=", ans)
  }


}
