package main

import "fmt"

func main() {
  for{
  flag:=true
  fmt.Println("Which operation you want to perform on the chaincode?(enter the number corresponding to the operation)")
  fmt.Println("1.Create Asset\n2.Query Asset\n3.Update Asset\n4.DeleteAsset\n5.Transfer Asset")
  var option int
  _,err:=fmt.Scanln(&option)
  if err !=nil {
    error:=fmt.Errorf("Error: %v",err,",but found a different type value")
    fmt.Println(error)
    return 
  }
  switch option {
    case 1:
      fmt.Println("create asset");
    case 2:
      fmt.Println("Query asset");
    case 3:
      fmt.Println("Update asset");
    case 4:
      fmt.Println("Delete asset");
    case 5:
      fmt.Println("Transfer asset");
    default:
      fmt.Println("\n------------Plaese select correct options..........");
      flag=false
  }
  if flag {
    temp:="\n"
    fmt.Println("\ndo you want to continue with other operations?  Type yes(default) or no")
    _,err:=fmt.Scanf("%s\n",&temp)
    if err!=nil{
      // error:=fmt.Errorf("Error: %v",err)
      // fmt.Println(error)
     return
    }
    if temp=="yes" || temp=="\n"{
      flag=true
    } else {
      return
    }
  }
  }
}
