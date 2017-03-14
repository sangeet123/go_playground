package main

import "fmt"

// represents generic type
type custom_type interface {}

// interface that list implements
type listInterface interface{
	insert_at(data *List, loc int)
	insert(data *List)
	append(data *List)
	delete(data *List)
	iterator()func() *Node
}

// This type represent the single node of list
type Node struct{
	data custom_type
	next *Node
	prev *Node
}

//List wrapper type
type List struct{
	root *Node
	end  *Node
}

//get the new node 
func get_node(data custom_type, prev *Node, next *Node)*Node{
	return &Node{data:data,prev:prev,next:next}
}

//insert operation to list
func (list *List)insert(data custom_type){
	node := get_node(data, nil , list.root)
	if list.root == nil{
		list.end  = node 
		list.root = node
		return
	}
	list.root.prev = node
	list.root = node
}

//append operation to list
func (list *List)append(data custom_type){
	node := get_node(data, list.end , nil)
	if list.root == nil{
		list.root = node 
		list.end = node
		return
	}
	list.end.next = node
	list.end = list.end.next
}


func (list List) iterator()func() *Node{
	head := list.root
	return func()*Node{
		to_return := head
		if head != nil{
			head = head.next
		}
		return to_return
	}
}

func(node *Node) has_next()bool{
	return node != nil
}

func main(){
	root := new(List)
	root.insert(4)
	root.insert(5)
	root.insert(6)
	root.insert(7)
	root.append(3)
	root.append(1)

	next := root.iterator()

	head := next()

	for head.has_next(){
		fmt.Println(head.data)
		head = next()
	}
}