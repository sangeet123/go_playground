package tree


func WaysToConstructBst(arr []int)int{
  if len(arr) == 0 {
    return 1
  }
  l,r := split(arr[1:],arr[0])
  waysOfSelection := combination(len(arr)-1, len(r))
  return  waysOfSelection * WaysToConstructBst(l) * WaysToConstructBst(r)
}

func combination(n,r int) int{
  result := 1
  for i := 0 ; i < r ; i++{
    result *= (n - i)
  }

  for i := 2 ; i <= r ; i++{
    result /= i
  }
  return result
}

func split(arr []int, s int)([]int,[]int){
  l := []int{}
  r := []int{}

  for _, v := range arr{
    if v > s{
      r = append(r, v)
    }else{
      l = append(l, v)
    }
  }

  return l,r
}
