package auxfunctions

func GetPhi(n int) []int {
	sieve := make([]bool, n+1)
	primes := []int{}

	for p := 2; p*p <= n; p++ {
		if !sieve[p] {
			for i := p * p; i <= n; i += p {
				sieve[i] = true
			}
		}
	}

	for p := 2; p <= n; p++ {
		if !sieve[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func ContainsInt(n int, slice []int) (bool, int) {
	for index, element := range slice {
		if element == n {
			return true, index
		}
	}
	return false, -1
}
