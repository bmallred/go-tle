package tle

import "testing"

// Test a valid payload can be parsed
func TestParse(t *testing.T) {
	payload := make([]string, 3)
	payload[0] = "ISS (ZARYA)"
	payload[1] = "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927"
	payload[2] = "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537"

	tle := Tle{}
	if err := tle.Parse(payload); err != nil {
		t.Fail()
	}
}

// Test an invalid payload will return an error
func TestParsePayloadFail(t *testing.T) {
	payload := make([]string, 3)
	payload[0] = "ISS (ZARYA)"
	payload[1] = "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927"

	tle := Tle{}
	if err := tle.Parse(payload); err == nil {
		t.Fail()
	}
}

// Test an invalid title will return an error
func TestParseTitleFail(t *testing.T) {
	payload := make([]string, 3)
	payload[0] = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	payload[1] = "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927"
	payload[2] = "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537"

	tle := Tle{}
	if err := tle.Parse(payload); err == nil {
		t.Fail()
	}
}

// Test an invalid line 1 will return an error
func TestParseLineOneFail(t *testing.T) {
	payload := make([]string, 3)
	payload[0] = "ISS (ZARYA)"
	payload[1] = ""
	payload[2] = "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537"

	tle := Tle{}
	if err := tle.Parse(payload); err == nil {
		t.Fail()
	}
}

// Test an invalid line 2 will return an error
func TestParseLineTwoFail(t *testing.T) {
	payload := make([]string, 3)
	payload[0] = "ISS (ZARYA)"
	payload[1] = "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927"
	payload[2] = ""

	tle := Tle{}
	if err := tle.Parse(payload); err == nil {
		t.Fail()
	}
}
