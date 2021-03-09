package bible

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ParseKJVBooks() {
	f, err := os.Open("kjv.csv")
	if err != nil {
		log.Fatalf("Unable to read file %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '	'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	kjv := "KJVBooks Books = Books{\n"

	book := ""
	chapter := ""

	for i, row := range records {
		if len(row) != 4 {
			log.Fatalf("wtf: %v", row)
		}

		// book
		newBook := row[0] != book

		if newBook {
			book = row[0]
			chapter = ""

			if i > 0 {
				// close chapter
				kjv += `		},` + "\n"

				// close book
				kjv += `	},` + "\n"
			}

			kjv += fmt.Sprintf(`	"%s": {`+"\n", book)
		}

		// chapter
		if newBook || row[1] != chapter {
			chapter = row[1]

			if !newBook && i > 0 {
				// close chapter
				kjv += `		},` + "\n"
			}

			kjv += fmt.Sprintf(`		%s: {`+"\n", chapter)
		}

		// verse
		kjv += fmt.Sprintf(`			%s: "%s",`+"\n", row[2], row[3])
	}

	// close chapter
	kjv += `		},` + "\n"

	// close book
	kjv += `	},` + "\n"

	// close kjv

	kjv += "}"

	fmt.Println(kjv)
}

func ParseKJVBooksIndex() {
	f, err := os.Open("kjv.csv")
	if err != nil {
		log.Fatalf("Unable to read file %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '	'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	kjv := "KJVBooksIndex BooksIndex = BooksIndex\n"

	book := ""

	for _, row := range records {
		if len(row) != 4 {
			log.Fatalf("wtf: %v", row)
		}

		// book
		if row[0] != book {
			book = row[0]

			kjv += fmt.Sprintf(`	"%s",`+"\n", book)
		}
	}

	// close kjv

	kjv += "}"

	fmt.Println(kjv)
}
