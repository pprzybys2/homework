package slices

import "sort"

func SortStringsInPlace(s []string) {
	sort.Strings(s)
}

func SortStringsPure(s []string) []string {
	var s2 = make([]string, len(s))
	copy(s2, s)
	sort.Strings(s2)
	return s2
}

type User struct {
	FirstName string
	LastName  string
}

func SortUsersPure(s []User) []User {
	u := make([]User, len(s))
	copy(u, s)
	sort.Slice(u, func(i, j int) bool {
		if u[i].FirstName != u[j].FirstName {
			return u[i].FirstName < u[j].FirstName
		}
		return u[i].LastName < u[j].LastName
	})
	return u
}

func SortUsersPureDesc(s []User) []User {
	u := make([]User, len(s))
	copy(u, s)
	sort.Slice(u, func(i, j int) bool {
		if u[i].FirstName != u[j].FirstName {
			return u[i].FirstName > u[j].FirstName
		}
		return u[i].LastName > u[j].LastName
	})
	return u
}
