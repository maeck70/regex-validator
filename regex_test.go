package regexvalidator

import (
	"fmt"
	"testing"
)

type test_t struct {
	v string // validator from yaml file
	t string // test value
	r bool   // expected result
}

func TestRegex(t *testing.T) {

	Load("validator_regex_test.yaml")

	ips := []test_t{
		// IP v4 tests
		{"ipv4Regex", "192.168.1.1", true},
		{"ipv4Regex", "10.0.0.255", true},
		{"ipv4Regex", "172.16.0.0", true},
		{"ipv4Regex", "255.255.255.255", true},
		{"ipv4Regex", "0.0.0.0", true},
		{"ipv4Regex", "192.168.1", false},     // Invalid - missing octet
		{"ipv4Regex", "192.168.1.256", false}, // Invalid - octet out of range
		{"ipv4Regex", "192.168..1", false},    // Invalid - double dot
		{"ipv4Regex", "a.b.c.d", false},       // Invalid - non-numeric
		{"ipv4Regex", "192.168.1.1.1", false}, // Invalid - too many octets
		{"ipv4Regex", "::1", false},           // Invalid - IPv6

		// IP v6 tests
		{"ipv6Regex", "::1", true},
		{"ipv6Regex", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"ipv6Regex", "2001:db8:85a3:0:0:8a2e:370:7334", true},
		{"ipv6Regex", "2001:db8:85a3::8a2e:370:7334", true},
		{"ipv6Regex", "fe80:0:0:0:0:0:0:0", true},
		{"ipv6Regex", "fe80::", true},
		{"ipv6Regex", "0::0", true},

		{"ipv6Regex", "2001:0db8:85a3:0000:0000:8a2e:0370:7334:1", false}, // Invalid - too many groups
		{"ipv6Regex", "2001:0gb8:85a3:0000:0000:8a2e:0370:7334", false},   // Invalid - non-hex
		// {"ipv6Regex", "2001:db8:85a3::8a2e:370", false},                   // Invalid - too few groups
		{"ipv6Regex", "192.168.1.1", false}, // Invalid - IPv4

		// Email Tests
		{"emailRegex", "foo@google.com", true},

		{"emailRegex", "foo@google", false},
	}

	rv := "is a valid"
	ri := "is NOT a valid"

	for _, ip := range ips {

		// Execute each test
		res := ValidateRegex(ip.v, ip.t)

		// Validate results
		r := rv
		if !res {
			r = ri
		}
		if res != ip.r {
			t.Fatalf("Expected %s in regex test %s to be %v, but got %v", ip.t, ip.v, ip.r, res)
		}
		fmt.Printf("%s %s %s address [PASS]\n", ip.t, r, ip.v)
	}
}
