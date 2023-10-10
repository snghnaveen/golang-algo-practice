package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnqiueEmail(t *testing.T) {
	t.Log(`
	Every valid email consists of a local name and a domain name, 
	separated by the '@' sign. Besides lowercase letters, 
	the email may contain one or more '.' or '+'.

	For example, in "alice@leetcode.com", "alice" is the local name, and "leetcode.com" is the domain name.
	If you add periods '.' between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name. Note that this rule does not apply to domain names.

	For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.
	If you add a plus '+' in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered. Note that this rule does not apply to domain names.

	For example, "m.y+name@email.com" will be forwarded to "my@email.com".
	It is possible to use both of these rules at the same time.

	Given an array of strings emails where we send one email to each emails[i], 
	return the number of different addresses that actually receive mails.

	Example 1:
	Input: emails = ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
	Output: 2
	Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails.
	
	Example 2:
	Input: emails = ["a@leetcode.com","b@leetcode.com","c@leetcode.com"]
	Output: 3
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []string{
			"test.email+alex@leetcode.com",
			"test.e.mail+bob.cathy@leetcode.com",
			"testemail+david@lee.tcode.com",
		}
		exp := 2
		assert.Equal(t, exp, RunUniqueEmail(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
		exp := 3
		assert.Equal(t, exp, RunUniqueEmail(in))
	})
}

func RunUniqueEmail(in []string) int {
	unqEmail := make(map[string]struct{})

	for _, email := range in {
		unqEmail[getEmail(email)] = struct{}{}
	}

	return len(unqEmail)
}

func getEmail(email string) string {
	i := 0
	local := ""
	for email[i] != '@' && email[i] != '+' {
		if email[i] != '.' {
			local = local + string(email[i])
		}
		i++
	}
	for email[i] != '@' {
		i++
	}

	domain := string(email[i+1:])

	parsedEmail := local + "@" + domain

	return parsedEmail
}
