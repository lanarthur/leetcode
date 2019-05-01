/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */
func countPrimes(n int) int {
	if n == 0 {
		return 0
	}
	prime := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		prime[i] = true
	}
	p := 2
	for p*p <= n {
		if prime[p] {
			for i := p * p; i < n+1; i += p {
				prime[i] = false
			}
		}
		p += 1
	}
	count := 0
	for i := 2; i < n; i++ {
		if prime[i] {
			count++
		}
	}
	return count
}

