package user_handles

import (
	"fmt"
	"sort"
	"userservice-go/types"
)

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(u1, u2 *types.UserOut) bool

// Closures that order the User structure.
var email = func(u1 *types.UserOut, u2 *types.UserOut) bool {
	return u1.Email < u2.Email
}

var username = func(u1, u2 *types.UserOut) bool {
	return u1.Username < u2.Username
}

var createdAt = func(u1, u2 *types.UserOut) bool {
	// TBD: Comparison of createdAt
	return true
}

var modifiedAt = func(u1, u2 *types.UserOut) bool {
	// TBD: Comparison of modifiedAt
	return true
}

var byDecreasingEmail = func(u1, u2 *types.UserOut) bool {
	return u2.Email < u1.Email
}

var byDecreasingUsername = func(u1, u2 *types.UserOut) bool {
	return u2.Username < u1.Username
}

var byDecreasingCreatedAt = func(u1, u2 *types.UserOut) bool {
	// TBD: Comparison of createdAt
	return true
}

var byDecreasingModifiedAt = func(u1, u2 *types.UserOut) bool {
	// TBD: Comparison of modifiedAt
	return true
}

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(users []types.UserOut) {
	us := &userSorter{
		users: users,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(us)
}

// userSorter joins a By function and a slice of Planets to be sorted.
type userSorter struct {
	users []types.UserOut
	by    func(p1, p2 *types.UserOut) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *userSorter) Len() int {
	return len(s.users)
}

// Swap is part of sort.Interface.
func (s *userSorter) Swap(i, j int) {
	s.users[i], s.users[j] = s.users[j], s.users[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *userSorter) Less(i, j int) bool {
	return s.by(&s.users[i], &s.users[j])
}

func SortByEmail(userList []types.UserOut, asc bool) []types.UserOut {
	if asc {
		By(email).Sort(userList)
	} else {
		By(byDecreasingEmail).Sort(userList)
	}
	fmt.Println("Sorted Users by email:", userList)
	return userList
}

func SortByUserName(userList []types.UserOut, asc bool) []types.UserOut {
	if asc {
		By(username).Sort(userList)
	} else {
		By(byDecreasingUsername).Sort(userList)
	}
	return userList
}

func SortByCreatedAt(userList []types.UserOut, asc bool) []types.UserOut {
	if asc {
		By(createdAt).Sort(userList)
	} else {
		By(byDecreasingCreatedAt).Sort(userList)
	}
	return userList
}

func SortByModifiedAt(userList []types.UserOut, asc bool) []types.UserOut {
	if asc {
		By(modifiedAt).Sort(userList)
	} else {
		By(byDecreasingModifiedAt).Sort(userList)
	}
	return userList
}
