package arrayprob

func MatchesPattern(word string, pattern string)bool{
  transistionMap := make(map[int]int)
  lastIndexOfAstisk := -1

  for i := 0 ; i < len(pattern); i++{
    if lastIndexOfAstisk != -1 && transistionMap[i] != '*'{
      transistionMap[i]=lastIndexOfAstisk + 1
    }
    if pattern[i] == '*' && i-1 != lastIndexOfAstisk{
      lastIndexOfAstisk = i
    }
  }
  return matches(word,pattern,transistionMap,0,0,len(pattern))
}

func matches(word,pattern string,transistionMap map[int]int,wind,pind, end int)bool{

  if pind == end && wind == len(word){
    return true
  }

  if wind >= len(word) || pind >= len(pattern){
    return false
  }


  if word[wind] == pattern[pind] || pattern[pind] == '?'{
    _,ok := transistionMap[pind]
    return matches(word,pattern,transistionMap,wind + 1,pind + 1,end) ||
      (ok && matches(word,pattern,transistionMap,wind + 1,transistionMap[pind],end))
  }

  if pattern[pind] == '*'{
    return matches(word,pattern,transistionMap,wind,pind + 1,end)||
     matches(word,pattern,transistionMap,wind + 1,pind + 1,end)
  }

  return false
}
