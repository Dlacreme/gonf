package gonf

import (
	"fmt"
	"testing"
)

func TestFetchFromFileNoFilters(t *testing.T) {
	expected := expectedSets()
	sets := fetchFromFile("tests/basic.gonf", []string{})
	if _, err := isSetsMatch(expected, sets); err != nil {
		setsMismatch(t, expected, sets)
		t.Errorf("Test failed.")
	}
}

func TestFetchFromFileWithFilters(t *testing.T) {
	expected := []Set{expectedSets()[0]}
	sets := fetchFromFile("tests/basic.gonf", []string{"hostname"})
	if _, err := isSetsMatch(expected, sets); err != nil {
		setsMismatch(t, expected, sets)
		t.Errorf("Test failed. Error: [%s]\n", err)
	}
}

func isSetsMatch(expected []Set, got []Set) (bool, error) {
	if len(expected) != len(got) {
		return false, fmt.Errorf("Number of item do not match. Expected %d. Got %d\n", len(expected), len(got))
	}
	for i := range expected {
		if expected[i].Key != got[i].Key {
			return false, fmt.Errorf("Index %d. %s != %s", i, expected[i].Key, got[i].Key)
		}

		if expected[i].RawValue != got[i].RawValue {
			return false, fmt.Errorf("Index %d. %s != %s", i, got[i].RawValue, got[i].RawValue)
		}
	}
	return true, nil
}

func printSets(t *testing.T, s []Set) {
	for _, s := range s {
		t.Logf("[%s]={%s}\n", s.Key, s.RawValue)
	}
}

func setsMismatch(t *testing.T, expected []Set, got []Set) {
	t.Logf("With filters failed. Expected:\n")
	printSets(t, expected)
	t.Logf("Got\n")
	printSets(t, got)
}

func expectedSets() []Set {
	sets := []Set{}
	sets = append(sets, Set{
		Key:      "hostname",
		RawValue: "localhost",
	})

	sets = append(sets, Set{
		Key:      "cert",
		RawValue: "file!(cert.crt)",
	})

	sets = append(sets, Set{
		Key:      "smtp_host",
		RawValue: "env?(SMTP_HOST, xxxxxxxxx)",
	})

	sets = append(sets, Set{
		Key:      "smtp_username",
		RawValue: "_",
	})

	sets = append(sets, Set{
		Key:      "smtp_password",
		RawValue: "!",
	})

	return sets
}
