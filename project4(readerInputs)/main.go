package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	fmt.Println("Welcome to the project4 , We run a pizza factory here")

	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza")
	rating, err := reader.ReadString('\n');
	if err != nil {
		fmt.Println("Error reading rating")
		return
	}else {
		fmt.Println("Rating is ", rating)
		fmt.Printf("Type of Rating is %T", rating)
	}

	//convert string to float 
	rating = strings.TrimSpace(rating)
	fmt.Println(" ")
	ratingInt , err:=strconv.ParseFloat(rating, 64)
	if(err != nil){
		fmt.Println("Error converting rating to float")
		// return
	}else {
		fmt.Println("Rating in float is ", ratingInt)
	}

	// var rate float64;
	// rate=4;
	// fmt.Println("Rating is ", rate)
	 friends := []string{"Rahul", "Rohit", "Rajesh", "Ramesh"}
	fmt.Println("My friends are ", friends)

	// var friends []string;
	// friends = make([]string, 4)
	// friends[0] = "Rahul"
	// friends[1] = "Rohit"
	// friends[2] = "Rajesh"
	// friends[3] = "Ramesh"
	// fmt.Println("My friends are ", friends)

	rollno := []int {1,2,3,4,5};
	fmt.Println("Roll no are ", rollno)

	var ratings []int;
	ratings = append(ratings, 10, 20, 30, 40, 50);
	fmt.Println("Ratings are ", ratings)



}