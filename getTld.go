package trans

import (
	"net/url"
	"strings"
)

func GetTLD(s string) string {

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	a := strings.Split(u.Host, ".")

	a = removeFirstIndex(a)

	if strings.Contains(u.Host, "www") {
		a = removeFirstIndex(a)
	}

	return strings.Join(a, ".")
}

func removeFirstIndex(a []string) []string {

	return a[1:len(a)] // Truncate slice.
}
