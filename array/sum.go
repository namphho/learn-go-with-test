package array

func Sum(numbers []int) int{
	sum := 0
	//simple version
	//for i:=0; i < 5; i++{
	//	sum += number[i]
	//}
	//refactor version
	for _, number := range numbers{
		sum += number
	}
	return sum
}

func SumAll(collections ...[]int) (sum []int){
	//simple version
	//lengthOfCollection := len(collections)
	//sum = make([]int, lengthOfCollection)
	//for i, collection := range collections{
	//	sum[i] = Sum(collection)
	//}
	//refactor version
	for _, collection := range  collections{
		sum = append(sum, Sum(collection))
	}
	return
}

func SumAllTails(collections ...[]int) (sum []int){
	for _, collection := range  collections{
		if len(collection) == 0{
			sum = append(sum, 0)
		} else {
			tails := collection[1:]
			sum = append(sum, Sum(tails))
		}
	}
	return
}