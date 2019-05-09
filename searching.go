package trans

func SearchNearest(n float32, nmap []float32) float32 {

	var lo int

	if n < nmap[0] {
		return nmap[0]
	}

	if n > nmap[len(nmap)-1] {
		return nmap[len(nmap)-1]
	}

	hi := len(nmap) - 1

	for lo <= hi {
		mid := int((hi + lo) / 2)

		if n < nmap[mid] {
			hi = mid - 1
		} else if n > nmap[mid] {
			lo = mid + 1
		} else {
			return nmap[mid]
		}
	}

	x := nmap[lo] - n
	y := n - nmap[hi]

	if x <= y {
		return nmap[lo]
	} else {
		return nmap[hi]
	}

}
