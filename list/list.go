package main

import "fmt"
import "reflect"

// represents generic type
type custom_type interface {}

//=============================Generic data type ======================================================

// interface that list implements
type listInterface interface{
	InsertAt(data custom_type, loc int)
	Insert (data custom_type)
	Append(data custom_type)
	Delete(data custom_type)
	contains(data custom_type) (*node, bool)
	IsPresent(data custom_type) bool
	iterator()func() *node
	get(loc int) custom_type
}

// This type represent the single node of list
type node struct{
	data custom_type
	next *node
	prev *node
}


// represents list data structure
type List struct{
	root *node
	end  *node
	size int
}
 //=============================================================List data structures completes here ===========================================

//Iterator interface
type Iterator interface{
	Iterator()ListIterator
	Next()custom_type
	Prev()custom_type
	HasNext()bool
}

type ListIterator struct{
	list List
}

func(list List)Iterator()*ListIterator{
	listIterator := new (ListIterator)
	listIterator.list = list
	return listIterator
}

func(this *ListIterator) HasNext()bool{
	return this.list.root != nil
}

func(this *ListIterator)Next()custom_type{
	data := this.list.root.data
	this.list.root = this.list.root.next
	return data
}

func(this *ListIterator)Prev()custom_type{
	data := this.list.root.data
	this.list.root = this.list.root.prev
	return data
}

//====================================================================List Iterator completes here =========================================


func(list *List)asset_valid_index(ind int){
	if ind > list.size || ind < 0{
		panic("Cannot access data pass the list size")
	}
}

func(list *List) get(loc int)custom_type{
	list.asset_valid_index(loc)
	head := list.root	
	for i:=0; i< loc ; i++{
		head = head.next
	}
	return head.data
}

func (list *List) contains(data custom_type)(*node, bool){
	head := list.root
	for head != nil {
		if reflect.DeepEqual(head.data, data){
			return head, true
		}
		head = head.next
	}

	return nil, false
}

func(list *List) IsPresent(data custom_type)bool{
	_,isPresent := list.contains(data)
	return isPresent
}

func (list *List) Delete(data custom_type){
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
func get_node(data custom_type, prev *node, next *node)*node{
	return &node{data:data,prev:prev,next:next}
}

//insert operation to list
func (list *List)Insert(data custom_type){
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

//Append operation to list
func (list *List)Append(data custom_type){
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
func(list *List)InsertAt(data custom_type, loc int){
	list.asset_valid_index(loc)
	
	if loc == 0{
		list.Insert(data)
		return
	}
	
	if loc == list.size{
		list.Append(data)
		return
	}

	head := list.root
	for i := 0 ; i < loc - 1 ; i++{
		head = head.next
	}
	node := get_node(data, head, head.next)
	head.next = node
	list.size++
}

func main(){
	root := new(List)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	root.Insert(7)
	root.Append(3)
	root.Append(1)
	root.InsertAt(15,0)
	root.InsertAt(25,7)
	root.InsertAt([]int{1,2,3,4,5,6,7},1)
	root.Delete(12)
	root.Delete([]int{1,2,3,4,5,6,7})

	listIterator := root.Iterator()

	for listIterator.HasNext() {
		fmt.Println(listIterator.Next())
	}
	fmt.Println(root.get(2))
}