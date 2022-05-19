package main

import "fmt"

func main() {
	user1 := &user{
		name: "a",
		age:  24,
	}

	user2 := &user{
		name: "b",
		age:  26,
	}

	user3 := &user{
		name: "c",
		age:  28,
	}

	user_collection := &userCollection{
		users: []*user{user1, user2, user3},
	}

	iterator := user_collection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %v\n", user)
	}

}
