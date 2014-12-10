package tle

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
)

// First line of the two-line element set format
type Line1 struct {
	LineNumber               string `json:"lineNumber"`
	SatelliteNumber          string `json:"satelliteNumber"`
	Classication             string `json:"classification"`
	InternationalDesignator1 string `json:"internationalDesignator1"`
	InternationalDesignator2 string `json:"internationalDesignator2"`
	InternationalDesignator3 string `json:"internationalDesignator3"`
	EpochYear                string `json:"epochYear"`
	Epoch                    string `json:"epoch"`
	FirstTimeDerivative      string `json:"firstTimeDerivative"`
	SecondTimeDerivative     string `json:"secondTimeDerivative"`
	Bstar                    string `json:"bstar"`
	Zero                     string `json:"zero"`
	ElementSetNumber         string `json:"elementSetNumber"`
	Checksum                 string `json:"checksum"`
}

// The second line of the two-line element set format
type Line2 struct {
	LineNumber              string `json:"lineNumber"`
	SatelliteNumber         string `json:"satelliteNumber"`
	Inclination             string `json:"inclination"`
	RightAscension          string `json:"rightAscension"`
	Eccentricity            string `json:"eccentricity"`
	ArgumentOfPerigee       string `json:"argumentOfPerigee"`
	MeanAnomaly             string `json:"meanAnomaly"`
	MeanMotion              string `json:"meanMotion"`
	RevolutionNumberAtEpoch string `json:"revolutionNumberAtEpoch"`
	Checksum                string `json:"checksum"`
}

// Two-line element set
type Tle struct {
	Title string `json:"title"`
	Line1 Line1  `json:"line1"`
	Line2 Line2  `json:"line2"`
}

// Reads a stream of JSON and translate it to a two-line element set
func (t *Tle) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t)
}

// Translates the two-line element set to a JSON byte array
func (t *Tle) ToJson() ([]byte, error) {
	return json.Marshal(t)
}

// Parse a 3-line string array to build the two-line element set
func (t *Tle) Parse(input []string) error {
	if len(input) != 3 {
		return errors.New("Expecting three lines")
	}
	if len(input[0]) > 24 {
		return errors.New("Title line cannot exceed 24 columns")
	}
	if len(input[1]) != 69 {
		return errors.New("Line 1 cannot exceed 69 columns")
	}
	if len(input[2]) != 69 {
		return errors.New("Line 2 cannot exceed 69 columns")
	}

	// The first line is the title (24 columns)
	t.Title = input[0]

	// The second line is `line 1` (69 columns)
	t.Line1.LineNumber = input[1][0:1]
	t.Line1.SatelliteNumber = input[1][2:7]
	t.Line1.Classication = input[1][7:8]
	t.Line1.InternationalDesignator1 = input[1][9:11]
	t.Line1.InternationalDesignator2 = input[1][11:14]
	t.Line1.InternationalDesignator3 = input[1][14:17]
	t.Line1.EpochYear = input[1][18:20]
	t.Line1.Epoch = input[1][20:32]
	t.Line1.FirstTimeDerivative = input[1][33:43]
	t.Line1.SecondTimeDerivative = input[1][44:52]
	t.Line1.Bstar = input[1][53:61]
	t.Line1.Zero = input[1][62:63]
	t.Line1.ElementSetNumber = input[1][64:68]
	t.Line1.Checksum = input[1][68:69]

	// The third line is `line 2` (69 columns)
	t.Line2.LineNumber = input[2][0:1]
	t.Line2.SatelliteNumber = input[2][2:7]
	t.Line2.Inclination = input[2][8:16]
	t.Line2.RightAscension = input[2][17:25]
	t.Line2.Eccentricity = input[2][26:33]
	t.Line2.ArgumentOfPerigee = input[2][34:42]
	t.Line2.MeanAnomaly = input[2][43:51]
	t.Line2.MeanMotion = input[2][52:63]
	t.Line2.RevolutionNumberAtEpoch = input[2][63:68]
	t.Line2.Checksum = input[2][68:69]

	return nil
}

// Scan from a reader taking and output payloads to standard output
func Scan(reader io.Reader, writer io.Writer) error {
	// Create a new scanner reading from the standard input
	scanner := bufio.NewScanner(reader)

	// Initialize the payload and counter
	payload := make([]string, 3)
	i := 0

	// Scan each line as it comes in
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		payload[i] = text
		i = i + 1

		if i > 2 {
			tle := Tle{}

			// Create a new two-line element set
			tle.Parse(payload)
			b, err := tle.ToJson()
			if err != nil {
				return err
			}
			writer.Write(b)

			// Re-initiate our variables to receive the next payload
			payload = make([]string, 3)
			i = 0
		}
	}

	// Check for any errors and report them.
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
