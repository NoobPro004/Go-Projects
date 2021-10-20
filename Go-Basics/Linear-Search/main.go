package main

import(
	"fmt"
)

func linearSearch(dataList []int,key int) bool{
	for _, item:= range dataList{
		if item ==key{
			return true;
		}
	}

	return false;
}


func main(){
	arr:= []int {1,2,3,4,5};
	if linearSearch(arr ,3){
		fmt.Println("Found")
	}else{
		fmt.Println("Not Found");
	}


}