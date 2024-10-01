package timeparse

import "testing"

func TestParseDateCorrect9h3m4s(t *testing.T) {
	time, err := ParseDate("09:03:04")
	if err != nil {
		t.Errorf("test threw error when if should have passed")
	}
	hours := time.hours
	minutes := time.minutes
	seconds := time.seconds
	if hours != 9 {
		t.Errorf("expected 9 hours, got %d", hours)
	}
	if minutes != 3 {
		t.Errorf("expected 3 minutes, got %d", minutes)
	}
	if seconds != 4 {
		t.Errorf("expected 4 seconds, got %d", seconds)
	}
}

func TestParseDateCorrect23h59m59s(t *testing.T) {
	time, err := ParseDate("23:59:59")
	if err != nil {
		t.Errorf("test threw error when if should have passed")
	}
	hours := time.hours
	minutes := time.minutes
	seconds := time.seconds
	if hours != 23 {
		t.Errorf("expected 23 hours, got %d", hours)
	}
	if minutes != 59 {
		t.Errorf("expected 59 minutes, got %d", minutes)
	}
	if seconds != 59 {
		t.Errorf("expected 59 seconds, got %d", seconds)
	}
}
func TestParseDateCorrect0h0m0s(t *testing.T) {
	time, err := ParseDate("0:0:0")
	if err != nil {
		t.Errorf("test threw error when if should have passed")
	}
	hours := time.hours
	minutes := time.minutes
	seconds := time.seconds
	if hours != 0 {
		t.Errorf("expected 0 hours, got %d", hours)
	}
	if minutes != 0 {
		t.Errorf("expected 0 minutes, got %d", minutes)
	}
	if seconds != 0 {
		t.Errorf("expected 0 seconds, got %d", seconds)
	}
}

func TestParseDateErrorInvalidFormat(t *testing.T) {
	_, err := ParseDate("asdfiasdfa")
	if err == nil {
		t.Error("expected error, got nil")
	} else if err.Error() != "unable to parse date string" {
		t.Errorf("expected error string \"unable to parse date string\", got \"%v\"", err)
	}
}
func TestParseDateErrorIncorrectLength(t *testing.T) {
	_, err := ParseDate("9:9")
	if err == nil {
		t.Error("expected error, got nil")
	} else if err.Error() != "unable to parse date string" {
		t.Errorf("expected error string \"unable to parse date string\", got \"%v\"", err)
	}

	_, err = ParseDate("9:9:9:9")
	if err == nil {
		t.Error("expected error, got nil")
	} else if err.Error() != "unable to parse date string" {
		t.Errorf("expected error string \"unable to parse date string\", got \"%v\"", err)
	}
}

func TestParseDateErrorPartsSingular(t *testing.T) {
	_, err := ParseDate("9:abc:89")
	if err == nil {
		t.Error("expected error, got nil")
	} else if err.Error() != "error parsing following time fields: minutes\ninvalid amount following time fields: seconds" {
		t.Errorf("expected error string \"error parsing following time fields: minutes\ninvalid amount following time fields: seconds\", got \"%v\"", err)
	}
}

func TestParseDateErrorPartsPlural(t *testing.T) {
	_, err := ParseDate("abc:abc:89")
	if err == nil {
		t.Error("expected error, got nil")
	} else if err.Error() != "error parsing following time fields: hours, minutes\ninvalid amount following time fields: seconds" {
		t.Errorf("expected error string \"error parsing following time fields: hours, minutes\ninvalid amount following time fields: seconds\", got \"%v\"", err)
	}

	_, err = ParseDate("abc:90:89")
	if err == nil {
		t.Error("expected error, got nil")
	} else if err.Error() != "error parsing following time fields: hours\ninvalid amount following time fields: minutes, seconds" {
		t.Errorf("expected error string \"error parsing following time fields: minutes\ninvalid amount following time fields: minutes, seconds\", got \"%v\"", err)
	}
}

// test table
func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"09:03:04", true},
		{"23:59:59", true},
		{"0:0:0", true},
		{"asdfiasdfa", false},
		{"9:9", false},
		{"9:9:9:9", false},
		{"9:abc:89", false},
		{"abc:abc:89", false},
		{"abc:90:89", false},
	}

	for _, data := range table {
		_, err := ParseDate(data.time)
		if data.ok && err != nil {
			t.Errorf("expected no error, got %v", err)

		}
	}
}
