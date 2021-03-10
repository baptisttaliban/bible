package bible

/*
@Ostracon@BaptistTaliban.com
https://BaptistTaliban.com/Ostracon

"Do not I hate them, O Lord, that hate thee? and am not I grieved with those that rise up against thee?" Psalm 139:21 📖👆

"For by grace are ye saved through faith; and that not of yourselves: it is the gift of God:
Not of works, lest any man should boast." Ephesians 2:8-9
*/

const (
	Genesis             = `Genesis`
	Exodus              = `Exodus`
	Leviticus           = `Leviticus`
	Numbers             = `Numbers`
	Deuteronomy         = `Deuteronomy`
	Joshua              = `Joshua`
	Judges              = `Judges`
	Ruth                = `Ruth`
	FirstSamuel         = `1 Samuel`
	SecondSamuel        = `2 Samuel`
	FirstKings          = `1 Kings`
	SecondKings         = `2 Kings`
	FirstChronicles     = `1 Chronicles`
	SecondChronicles    = `2 Chronicles`
	Ezra                = `Ezra`
	Nehemiah            = `Nehemiah`
	Esther              = `Esther`
	Job                 = `Job`
	Psalms              = `Psalms`
	Proverbs            = `Proverbs`
	Ecclesiastes        = `Ecclesiastes`
	SongOfSolomon       = `Song of Solomon`
	Isaiah              = `Isaiah`
	Jeremiah            = `Jeremiah`
	Lamentations        = `Lamentations`
	Ezekiel             = `Ezekiel`
	Daniel              = `Daniel`
	Hosea               = `Hosea`
	Joel                = `Joel`
	Amos                = `Amos`
	Obadiah             = `Obadiah`
	Jonah               = `Jonah`
	Micah               = `Micah`
	Nahum               = `Nahum`
	Habakkuk            = `Habakkuk`
	Zephaniah           = `Zephaniah`
	Haggai              = `Haggai`
	Zechariah           = `Zechariah`
	Malachi             = `Malachi`
	Matthew             = `Matthew`
	Mark                = `Mark`
	Luke                = `Luke`
	John                = `John`
	Acts                = `Acts`
	Romans              = `Romans`
	FirstCorinthians    = `1 Corinthians`
	SecondCorinthians   = `2 Corinthians`
	Galatians           = `Galatians`
	Ephesians           = `Ephesians`
	Philippians         = `Philippians`
	Colossians          = `Colossians`
	FirstThessalonians  = `1 Thessalonians`
	SecondThessalonians = `2 Thessalonians`
	FirstTimothy        = `1 Timothy`
	SecondTimothy       = `2 Timothy`
	Titus               = `Titus`
	Philemon            = `Philemon`
	Hebrews             = `Hebrews`
	James               = `James`
	FirstPeter          = `1 Peter`
	SecondPeter         = `2 Peter`
	FirstJohn           = `1 John`
	SecondJohn          = `2 John`
	ThirdJohn           = `3 John`
	Jude                = `Jude`
	Revelation          = `Revelation`
)

type (
	Bible struct {
		Books      Books
		BooksIndex BooksIndex
	}
	Book         string
	Books        map[Book]Chapters
	BooksIndex   []Book
	BookPosition int
	Chapters     map[Chapter]Verses
	Chapter      int
	Scripture    string
	Verses       map[Verse]Scripture
	Verse        int
)

func GetScripture(book Book, chapter Chapter, verse Verse) Scripture {
	return KJV.GetScripture(book, chapter, verse)
}

func (self Bible) GetScripture(book Book, chapter Chapter, verse Verse) Scripture {
	if _, ok := self.Books[book]; !ok {
		// book not found
		return ""
	}

	if _, ok := self.Books[book][chapter]; !ok {
		// chapter not found
		return ""
	}

	if _, ok := self.Books[book][chapter][verse]; !ok {
		// verse not found
		return ""
	}

	return self.Books[book][chapter][verse]
}

func GetNextScripture(book Book, chapter Chapter, verse Verse) (Book, Chapter, Verse, Scripture) {
	return KJV.GetNextScripture(book, chapter, verse)
}

func (self Bible) GetNextScripture(book Book, chapter Chapter, verse Verse) (Book, Chapter, Verse, Scripture) {
	if _, ok := self.Books[book]; !ok {
		// book not found
		return "", -1, -1, ""
	}

	nextVerse := self.Books.GetNextVerse(book, chapter, verse)
	if nextVerse == -1 {
		// verse not found
		return "", -1, -1, ""
	}

	if nextVerse < verse {
		// next chapter
		nextChapter := self.Books.GetNextChapter(book, chapter)
		if nextChapter < chapter {
			// next book
			nextBook := self.BooksIndex.GetNextBook(book)
			return nextBook, 1, 1, self.Books[nextBook][1][1]
		}

		return book, nextChapter, nextVerse, self.Books[book][nextChapter][nextVerse]
	}

	return book, chapter, nextVerse, self.Books[book][chapter][nextVerse]
}

func GetPreviousScripture(book Book, chapter Chapter, verse Verse) (Book, Chapter, Verse, Scripture) {
	return KJV.GetPreviousScripture(book, chapter, verse)
}

func (self Bible) GetPreviousScripture(book Book, chapter Chapter, verse Verse) (Book, Chapter, Verse, Scripture) {
	if _, ok := self.Books[book]; !ok {
		// book not found
		return "", -1, -1, ""
	}

	previousVerse := self.Books.GetPreviousVerse(book, chapter, verse)
	if previousVerse == -1 {
		// verse not found
		return "", -1, -1, ""
	}

	if previousVerse > verse {
		// previous chapter

		previousChapter := self.Books.GetPreviousChapter(book, chapter)

		if previousChapter > chapter {
			// previous book
			previousBook := self.BooksIndex.GetPreviousBook(book)

			lastChapter := Chapter(len(self.Books[previousBook]))
			lastVerse := Verse(len(self.Books[previousBook][lastChapter]))

			return previousBook, lastChapter, lastVerse, self.Books[previousBook][lastChapter][lastVerse]
		}

		lastVerse := Verse(len(self.Books[book][previousChapter]))

		return book, previousChapter, lastVerse, self.Books[book][previousChapter][lastVerse]
	}

	return book, chapter, previousVerse, self.Books[book][chapter][previousVerse]
}

func GetNextChapter(book Book, chapter Chapter) Chapter {
	return KJV.Books.GetNextChapter(book, chapter)
}

func (self Books) GetNextChapter(book Book, chapter Chapter) Chapter {
	if _, ok := self[book]; !ok {
		// book not found
		return -1
	}

	if _, ok := self[book][chapter]; !ok {
		// chapter not found
		return -1
	}

	if _, ok := self[book][chapter+1]; !ok {
		// start at the first chapter again
		return 1
	}

	// next chapter
	return chapter + 1
}

func GetPreviousChapter(book Book, chapter Chapter) Chapter {
	return KJV.Books.GetPreviousChapter(book, chapter)
}

func (self Books) GetPreviousChapter(book Book, chapter Chapter) Chapter {
	if _, ok := self[book]; !ok {
		// book not found
		return -1
	}

	if _, ok := self[book][chapter]; !ok {
		// chapter not found
		return -1
	}

	if _, ok := self[book][chapter-1]; !ok {
		// start at the last chapter again
		return Chapter(len(self[book]))
	}

	// previous chapter
	return chapter - 1
}

func GetNextVerse(book Book, chapter Chapter, verse Verse) Verse {
	return KJV.Books.GetNextVerse(book, chapter, verse)
}

func (self Books) GetNextVerse(book Book, chapter Chapter, verse Verse) Verse {
	if _, ok := self[book]; !ok {
		// book not found
		return -1
	}

	if _, ok := self[book][chapter]; !ok {
		// chapter not found
		return -1
	}

	if _, ok := self[book][chapter][verse]; !ok {
		// verse not found
		return -1
	}

	if _, ok := self[book][chapter][verse+1]; !ok {
		// start at the first verse again
		return 1
	}

	// next verse
	return verse + 1
}

func GetPreviousVerse(book Book, chapter Chapter, verse Verse) Verse {
	return KJV.Books.GetPreviousVerse(book, chapter, verse)
}

func (self Books) GetPreviousVerse(book Book, chapter Chapter, verse Verse) Verse {
	if _, ok := self[book]; !ok {
		// book not found
		return -1
	}

	if _, ok := self[book][chapter]; !ok {
		// chapter not found
		return -1
	}

	if _, ok := self[book][chapter][verse]; !ok {
		// verse not found
		return -1
	}

	if _, ok := self[book][chapter][verse-1]; !ok {
		// start at the last verse again
		return Verse(len(self[book][chapter]))
	}

	// previous verse
	return verse - 1
}

func GetBookPosition(book Book) BookPosition {
	return KJV.BooksIndex.GetBookPosition(book)
}

func (self BooksIndex) GetBookPosition(book Book) BookPosition {
	for i, b := range self {
		if b == book {
			// book found
			return BookPosition(i)
		}
	}

	// book not found
	return -1
}

func GetNextBook(book Book) Book {
	return KJV.BooksIndex.GetNextBook(book)
}

func (self BooksIndex) GetNextBook(book Book) Book {
	i := self.GetNextBookPosition(book)
	if i == -1 {
		// book not found
		return ""
	}

	return self[i]
}

func GetNextBookPosition(book Book) BookPosition {
	return KJV.BooksIndex.GetNextBookPosition(book)
}

func (self BooksIndex) GetNextBookPosition(book Book) BookPosition {
	i := self.GetBookPosition(book)
	if i == -1 {
		// book not found
		return -1
	}

	if i == BookPosition(len(self)-1) {
		// start at the first book again
		return 0
	}

	// next book
	return i + 1
}

func GetPreviousBook(book Book) Book {
	return KJV.BooksIndex.GetPreviousBook(book)
}

func (self BooksIndex) GetPreviousBook(book Book) Book {
	i := self.GetPreviousBookPosition(book)
	if i == -1 {
		// book not found
		return ""
	}

	return self[i]
}

func GetPreviousBookPosition(book Book) BookPosition {
	return KJV.BooksIndex.GetPreviousBookPosition(book)
}

func (self BooksIndex) GetPreviousBookPosition(book Book) BookPosition {
	i := self.GetBookPosition(book)
	if i == -1 {
		// book not found
		return -1
	}

	if i == 0 {
		// start at the last book again
		return BookPosition(len(self) - 1)
	}

	// previous book
	return i - 1
}
