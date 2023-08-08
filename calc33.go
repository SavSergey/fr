/*
# All rights reserved!
# Author: Sergey Savostin S.
# Email: work11111111111@gmail.com
# Email: work1111111111@yandex.ru
*/

package main 
  
  
 import ( 
   "fmt" 
   "strconv" 
   "strings" 
   "bufio"   // модуль для ввода данных stdin 
   "os"      
   ) 
  
 
 //********************
 //********************
 // Ред. 07/08/2023
 // Ред. 08/08/2023
 func get_keys_a(s string) (z string, a int, b int) { 
   
   f := strconv.Atoi 
   char := ""
   i, l := 0, len(s)
   for i < l {
       char = string(s[i])
       if (char == "+") || (char == "-") || (char == "*") || (char == "/") {
           z = char
           a, _ = f(s[:i])
           b, _ = f(s[(i+1):])
           i = l
       }
       i++
   }
   
   if (a < 0) || (10 < a) || (b < 0) || (10 < b) {
        panic("\n___ERROR___\n The value is out of range [0..10]!\n")
   }

   return  
 } 
  
  
  
 //********************
 //********************
 // Ред. 08/08/2023
 func get_keys_r(s string) (z string, a int, b int) { 
 
   char := ""
   i, j, l := 0, 0, len(s)
   for i < l {
       char = string(s[i])
       if (char == "+") || (char == "-") || (char == "*") || (char == "/") {
           z = char
           //...
           j = i
           i = l
       }
       i++
   }

   roma_array := [10]string{"I", "II", "III", "IV","V", "VI", "VII", "VIII", "IX", "X"} 
   
   i = 0
   for i < 10 { 
     if roma_array[i] == s[:j] { 
       a = i+1
       i = 32
     }
     i++
   }
   
   //...
   j++
   
   i = 0
   for i < 10 { 
     if roma_array[i] == s[j:] { 
       b = i+1
       i = 32
     }
     i++
   }
   
   
   if (a == 0) || (b == 0) {
       panic("\n___ERROR___\n The value is out of range [1..10]!\n")
   }
   
   //...
   return 
 }
  
  
 
 //********************
 //********************
 func processing(z string, a, b int) (res int) { 
   switch z { 
     case "+": 
       res = a+b 
     case "-": 
       res = a-b 
     case "*": 
       res = a*b 
     case "/": 
       if b == 0 {
           panic("\n___ERROR___\n Division by zero!")
       }
       res = a/b 
     default: 
       panic("\n___ERROR___\n Unknown operation!\n")
   } 
   return  
 } 
  
  
 //********************
 //********************
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
   return  
 } 


 //********************
 //********************
 // Доб. 29/07/2023
 // Ввод выражения пользователем.
 func input_data_from_user() (data string) { 
     fmt.Println("Input expression:") 
     reader := bufio.NewReader(os.Stdin)  
     data, _ = reader.ReadString('\n') 
     data = strings.TrimSpace(data) 
     data = strings.ToUpper(data) 
     return  
 } 


 //********************
 //********************
 // Доб. 07/08/2023
 // Ред. 08/08/2023
 // Проверка данных. 
func check_data_from_user(data string) (flag bool) { 

    fmt.Println(data)
    fmt.Println()
    
    roman := [3]string{"I", "V", "X"}
    arabic := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
    
    // arabic
    flag = true 
    char := ""
    i, j, k, m, l := 0, 0, 0, 0, len(data)
    for i < l {
        char = string(data[i])
        j = 0
        for j < 10 {
            if char == arabic[j] {
                j = 32
            } else if (char == ("+") || (char == ("-")) || (char == ("*")) || (char == ("/"))) {
                j = 32
                m = i
                k++
            } 
            j++
        }
        
        if j != 33 {
            flag = false
            i = l
                  
    //fmt.Println("\n___INFO___\n The expression is not valid as an «arabic expression»!\n")
        }
        
        i++
    }

    if flag {
        msg := ""
        switch {
            case l < 3:
              msg = "\n___ERROR___\n Слишком короткое выражение!"
            case i != l:
              msg = "\n___ERROR___\n Обнаружен недопустимый символ!"
            case k == 0:
              msg = "\n___ERROR___\n Отсутствует допустимый знак операции!"
            case k > 1:
              msg = "\n___ERROR___\n Превышено максимальное количество допустимых операций!"
            case m+1 == l:
              msg = "\n___ERROR___\n Отсутствует второй аргумент!"
            case m == 0:
              msg = "\n___ERROR___\n Отсутствует первый аргумент!"
        }
        
        if msg != "" {
          panic(msg)
        }
    
        return
    }
    
    
    // roman 
    char = ""
    i, j, k, m, l = 0, 0, 0, 0, len(data)
    for i < l {
        char = string(data[i])
        j = 0
        for j < 3 {
            if char == roman[j] {
                j = 32
            } else if (char == ("+") || (char == ("-")) || (char == ("*")) || (char == ("/"))) {
                j = 32
                m = i
                k++
            } 
            j++
        }
        
        if j != 33 {
            flag = false
            i = l
          
    //fmt.Println("\n___INFO___\n The expression is not valid as an «roman expression»!\n")
        }
        
        i++
    }
    
    
    msg := ""
    switch {
        case l < 3:
          msg = "\n___ERROR___\n Слишком короткое выражение!"
        case i != l:
          msg = "\n___ERROR___\n Обнаружен недопустимый символ!"
        case k == 0:
          msg = "\n___ERROR___\n Отсутствует допустимый знак операции!"
        case k > 1:
          msg = "\n___ERROR___\n Превышено максимальное количество допустимых операций!"
        case m+1 == l:
          msg = "\n___ERROR___\n Отсутствует второй аргумент!"
        case m == 0:
          msg = "\n___ERROR___\n Отсутствует первый аргумент!"
    }
    
    
    if msg != "" {
        panic(msg)
    }
    
    return  
 } 




 //********************
 //********************
 func main() { 
 
   var ( 
     expr, z string 
     ans string
     a, b int 
     res int
     flag bool
     ) 
  
   expr = input_data_from_user() 
   flag = check_data_from_user(expr)
   
   if flag { 
     // arabic 
     z, a, b = get_keys_a(expr) 
     res = processing(z, a, b) 
     fmt.Println("\nOutput:")
     fmt.Println(expr,"=", res) 
   } else { 
     // roman
     z, a, b = get_keys_r(expr) 
     res = processing(z, a, b) 
     if res < 1 {
         panic("\n___ERROR___\n This result is less than one and cannot be written in Roman numerals!\n")
     }
     ans = convert_a_r(res)
     fmt.Println("\nOutput:")
     fmt.Println(expr,"=", ans) 
   } 
  
 }






