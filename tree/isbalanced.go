package tree

// IsBalanced returns if tree is balance or not
func (this *intTree) IsBalanced()bool{
  ok,_:=isBalanced(this.root)
  return ok
}

func isBalanced(n *node)(bool, int){
  if n == nil{
    return true, 0
  }

  okl, hl := isBalanced(n.l)
  if !okl{
    return false, -1
  }

  okr, hr := isBalanced(n.r)
  if !okr{
    return false, -1
  }

  if abs(hl-hr) > 1{
    return false, -1
  }

  if hl >= hr {
    return true, 1 + hl
  }
  return true, 1 + hr
}

func abs(no int)int{
  if no >= 0{
    return no
  }
  return -no
}
