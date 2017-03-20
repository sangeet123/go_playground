package main

import "fmt"
import "reflect"

// represents generic type
type custom_type interface {}

// interface that list implements
type listInterface interface{
	insert_at(data custom_type, loc int)
	insert(data custom_type)
	append(data custom_type)
	delete(data custom_type)
	contains(data custom_type) (*Node, bool)
	iterator()func() *Node
	get(loc int) *Node
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
	size int
}

func(list *List)asset_valid_index(ind int){
	if ind > list.size || ind < 0{
		panic("Cannot access data pass the list size")
	}
}

func(list *List) get(loc int)*Node{
	list.asset_valid_index(loc)	
	next := list.iterator()
	head := next()
	for i:=0; i<loc; i++{
		head = next()
	}
	return head
}

func (list *List) contains(data custom_type)(*Node, bool){
	next := list.iterator()
	head := next()
	for head.has_next() {
		if reflect.DeepEqual(head.data, data){
			return head, true
		}
		head = next()
	}

	return nil, false
}

func (list *List) delete(data custom_type){
	node_to_delete, ok := list.contains(data)
	if ok{
		if node_to_delete == list.root{
			next := list.root.next
			next.prev = nil
			list.root.next = nil
			list.root = next
		}else if node_to_delete == list.end{
			list.end = node_to_delete.prev
			node_to_delete.prev = nil
			list.end.next = nil
		}else{
			prev := node_to_delete.prev
			next := node_to_delete.next
			prev.next = next
			next.prev = prev
			node_to_delete.prev = nil
			node_to_delete.next = nil
		}
		list.size--
	}
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
		list.size++
		return
	}
	list.root.prev = node
	list.root = node
	list.size++
}

//append operation to list
func (list *List)append(data custom_type){
	node := get_node(data, list.end , nil)
	if list.root == nil{
		list.root = node 
		list.end = node
		list.size++
		return
	}
	list.end.next = node
	list.end = list.end.next
	list.size++
}

//insert at location to list
func(list *List)insert_at(data custom_type, loc int){
	list.asset_valid_index(loc)
	
	if loc == 0{
		list.insert(data)
		return
	}
	
	if loc == list.size{
		list.append(data)
		return
	}

	next := list.iterator()
	head := next()
	for i := 0 ; i < loc - 1 ; i++{
		head = next()
	}
	node := get_node(data, head, head.next)
	head.next = node
	list.size++
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
	root.insert_at(15,0)
	root.insert_at(25,7)
	root.insert_at([]int{1,2,3,4,5,6,7},1)
	root.delete(12)
	root.delete([]int{1,2,3,4,5,6,7})

	next := root.iterator()
	head := next()
	for head.has_next(){
		fmt.Println(head.data)
		head = next()
	}

	fmt.Println(*(root.get(100)))
}