//ipsorter is a small package to sort a slice of IP addresses as strings
package ipsorter

import (
	"regexp"
	"sort"
	"strconv"
)

var res string = `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})(:\d*)?$`
var re *regexp.Regexp

type byIP []string

/* Canned functions */
func (a byIP) Len() int      { return len(a) }
func (a byIP) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byIP) Less(i, j int) bool {
	/* Break into parts */
	o := re.FindStringSubmatch(a[i])
	t := re.FindStringSubmatch(a[j])
	/* Go down the line to work out which is less */
	for k := 1; k < 5; k++ {
		/* Make ints */
		m, _ := strconv.Atoi(o[k])
		n, _ := strconv.Atoi(t[k])
		/* Try the next if they're equal */
		if m == n {
                        continue
		} else {
                        return m < n
                }
	}
	return false
}

/* Sort4 returns a sorted slice of IPv4 addresses with optional port numbers.
Strings which are not IPv4 addressses are discarded */
func Sort4(a []string) []string {
	/* Make the regex the first time it's used */
	if nil == re {
		re = regexp.MustCompile(res)
	}
	/* Slice to hold results */
        r := []string{}
	/* Iterate over input slice to make output slice */
CheckLoop:
	for _, s := range a {
		/* Split into octets */
		o := re.FindStringSubmatch(s)
		/* Give up if it didn't match at all */
		if len(o) < 5 {
			continue
		}
		/* Make sure each match is in the right range */
		for i := 1; i < 5; i++ {
			n, err := strconv.Atoi(o[i])
			if err != nil || n < 0 || n > 255 {
				continue CheckLoop
			}
		}
		/* Add to the slice to be sorted */
		r = append(r, s)
	}
	/* Sort the output slice */
	sort.Sort(byIP(r))
	return r
}
