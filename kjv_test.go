package bible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBible_GetScripture(t *testing.T) {
	assert.Equal(t,
		Scripture("In the beginning God created the heaven and the earth."),
		KJV.GetScripture("Genesis", 1, 1),
	)
	assert.Equal(t,
		Scripture(""),
		KJV.GetScripture("genesis", 1, 1),
	)
	assert.Equal(t,
		Scripture(""),
		KJV.GetScripture("Genesis", 0, 1),
	)
	assert.Equal(t,
		Scripture(""),
		KJV.GetScripture("Genesis", 1, 0),
	)

	assert.Equal(t,
		Scripture("The grace of our Lord Jesus Christ be with you all. Amen."),
		KJV.GetScripture("Revelation", 22, 21),
	)
	assert.Equal(t,
		Scripture(""),
		KJV.GetScripture("revelation", 22, 21),
	)
	assert.Equal(t,
		Scripture(""),
		KJV.GetScripture("Revelation", 0, 21),
	)
	assert.Equal(t,
		Scripture(""),
		KJV.GetScripture("Revelation", 22, 0),
	)
}

func TestBible_GetNextScripture(t *testing.T) {
	// In the beginning God created the heaven and the earth.
	book, chapter, verse, scripture := KJV.GetNextScripture("Genesis", 1, 1)
	assert.Equal(t, Book("Genesis"), book)
	assert.Equal(t, Chapter(1), chapter)
	assert.Equal(t, Verse(2), verse)
	assert.Equal(t,
		Scripture("And the earth was without form, and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face of the waters."),
		scripture,
	)
	book, chapter, verse, scripture = KJV.GetNextScripture("genesis", 1, 1)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetNextScripture("Genesis", 0, 1)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetNextScripture("Genesis", 1, 0)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)

	// And God saw every thing that he had made, and, behold, it was very good. And the evening and the morning were the sixth day.
	book, chapter, verse, scripture = KJV.GetNextScripture("Genesis", 1, 31)
	assert.Equal(t, Book("Genesis"), book)
	assert.Equal(t, Chapter(2), chapter)
	assert.Equal(t, Verse(1), verse)
	assert.Equal(t,
		Scripture("Thus the heavens and the earth were finished, and all the host of them."),
		scripture,
	)
	book, chapter, verse, scripture = KJV.GetNextScripture("genesis", 1, 31)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetNextScripture("Genesis", 0, 31)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetNextScripture("Genesis", 1, 0)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)

	// So Joseph died, being an hundred and ten years old: and they embalmed him, and he was put in a coffin in Egypt.
	book, chapter, verse, scripture = KJV.GetNextScripture("Genesis", 50, 26)
	assert.Equal(t, Book("Exodus"), book)
	assert.Equal(t, Chapter(1), chapter)
	assert.Equal(t, Verse(1), verse)
	assert.Equal(t,
		Scripture("Now these are the names of the children of Israel, which came into Egypt; every man and his household came with Jacob."),
		scripture,
	)

	// The grace of our Lord Jesus Christ be with you all. Amen.
	book, chapter, verse, scripture = KJV.GetNextScripture("Revelation", 22, 21)
	assert.Equal(t, Book("Genesis"), book)
	assert.Equal(t, Chapter(1), chapter)
	assert.Equal(t, Verse(1), verse)
	assert.Equal(t,
		Scripture("In the beginning God created the heaven and the earth."),
		scripture,
	)
	book, chapter, verse, scripture = KJV.GetNextScripture("revelation", 22, 21)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetNextScripture("Revelation", 0, 21)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetNextScripture("Revelation", 22, 0)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
}

func TestBible_GetPreviousScripture(t *testing.T) {
	// And the earth was without form, and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face of the waters.
	book, chapter, verse, scripture := KJV.GetPreviousScripture("Genesis", 1, 2)
	assert.Equal(t, Book("Genesis"), book)
	assert.Equal(t, Chapter(1), chapter)
	assert.Equal(t, Verse(1), verse)
	assert.Equal(t,
		Scripture("In the beginning God created the heaven and the earth."),
		scripture,
	)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("genesis", 1, 2)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Genesis", 0, 2)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Genesis", 1, 0)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)

	// Thus the heavens and the earth were finished, and all the host of them.
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Genesis", 2, 1)
	assert.Equal(t, Book("Genesis"), book)
	assert.Equal(t, Chapter(1), chapter)
	assert.Equal(t, Verse(31), verse)
	assert.Equal(t,
		Scripture("And God saw every thing that he had made, and, behold, it was very good. And the evening and the morning were the sixth day."),
		scripture,
	)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("genesis", 2, 1)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Genesis", 0, 1)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Genesis", 2, 0)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)

	// Now these are the names of the children of Israel, which came into Egypt; every man and his household came with Jacob.
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Exodus", 1, 1)
	assert.Equal(t, Book("Genesis"), book)
	assert.Equal(t, Chapter(50), chapter)
	assert.Equal(t, Verse(26), verse)
	assert.Equal(t,
		Scripture("So Joseph died, being an hundred and ten years old: and they embalmed him, and he was put in a coffin in Egypt."),
		scripture,
	)

	// In the beginning God created the heaven and the earth.
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Genesis", 1, 1)
	assert.Equal(t, Book("Revelation"), book)
	assert.Equal(t, Chapter(22), chapter)
	assert.Equal(t, Verse(21), verse)
	assert.Equal(t,
		Scripture("The grace of our Lord Jesus Christ be with you all. Amen."),
		scripture,
	)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("genesis", 1, 1)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Genesis", 0, 1)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
	book, chapter, verse, scripture = KJV.GetPreviousScripture("Genesis", 1, 0)
	assert.Equal(t, Book(""), book)
	assert.Equal(t, Chapter(-1), chapter)
	assert.Equal(t, Verse(-1), verse)
	assert.Equal(t, Scripture(""), scripture)
}

func TestBooksIndex_GetBookPosition(t *testing.T) {
	assert.Equal(t, BookPosition(0), KJVBooksIndex.GetBookPosition("Genesis"))
	assert.Equal(t, BookPosition(65), KJVBooksIndex.GetBookPosition("Revelation"))
	assert.Equal(t, BookPosition(-1), KJVBooksIndex.GetBookPosition("genesis"))
}

func TestBooksIndex_GetNextBookPosition(t *testing.T) {
	assert.Equal(t, BookPosition(1), KJVBooksIndex.GetNextBookPosition("Genesis"))
	assert.Equal(t, BookPosition(2), KJVBooksIndex.GetNextBookPosition("Exodus"))
	assert.Equal(t, BookPosition(65), KJVBooksIndex.GetNextBookPosition("Jude"))
	assert.Equal(t, BookPosition(0), KJVBooksIndex.GetNextBookPosition("Revelation"))
	assert.Equal(t, BookPosition(-1), KJVBooksIndex.GetNextBookPosition("genesis"))
}

func TestBooksIndex_GetPreviousBookPosition(t *testing.T) {
	assert.Equal(t, BookPosition(65), KJVBooksIndex.GetPreviousBookPosition("Genesis"))
	assert.Equal(t, BookPosition(0), KJVBooksIndex.GetPreviousBookPosition("Exodus"))
	assert.Equal(t, BookPosition(63), KJVBooksIndex.GetPreviousBookPosition("Jude"))
	assert.Equal(t, BookPosition(64), KJVBooksIndex.GetPreviousBookPosition("Revelation"))
	assert.Equal(t, BookPosition(-1), KJVBooksIndex.GetPreviousBookPosition("genesis"))
}

func TestBooksIndex_GetNextBook(t *testing.T) {
	assert.Equal(t, Book("Exodus"), KJVBooksIndex.GetNextBook("Genesis"))
	assert.Equal(t, Book("Revelation"), KJVBooksIndex.GetNextBook("Jude"))
	assert.Equal(t, Book("Genesis"), KJVBooksIndex.GetNextBook("Revelation"))
	assert.Equal(t, Book(""), KJVBooksIndex.GetNextBook("genesis"))
}

func TestBooksIndex_GetPreviousBook(t *testing.T) {
	assert.Equal(t, Book("Revelation"), KJVBooksIndex.GetPreviousBook("Genesis"))
	assert.Equal(t, Book("Genesis"), KJVBooksIndex.GetPreviousBook("Exodus"))
	assert.Equal(t, Book("Jude"), KJVBooksIndex.GetPreviousBook("Revelation"))
	assert.Equal(t, Book(""), KJVBooksIndex.GetPreviousBook("genesis"))
}

func TestBooks_GetNextChapter(t *testing.T) {
	assert.Equal(t, Chapter(2), KJVBooks.GetNextChapter("Genesis", 1))
	assert.Equal(t, Chapter(1), KJVBooks.GetNextChapter("Genesis", 50))
	assert.Equal(t, Chapter(-1), KJVBooks.GetNextChapter("Genesis", 0))
	assert.Equal(t, Chapter(-1), KJVBooks.GetNextChapter("genesis", 1))
	assert.Equal(t, Chapter(-1), KJVBooks.GetNextChapter("Genesis", 51))

	assert.Equal(t, Chapter(2), KJVBooks.GetNextChapter("Revelation", 1))
	assert.Equal(t, Chapter(1), KJVBooks.GetNextChapter("Revelation", 22))
	assert.Equal(t, Chapter(-1), KJVBooks.GetNextChapter("Revelation", 0))
	assert.Equal(t, Chapter(-1), KJVBooks.GetNextChapter("revelation", 1))
	assert.Equal(t, Chapter(-1), KJVBooks.GetNextChapter("Revelation", 23))
}

func TestBooks_GetPreviousChapter(t *testing.T) {
	assert.Equal(t, Chapter(50), KJVBooks.GetPreviousChapter("Genesis", 1))
	assert.Equal(t, Chapter(49), KJVBooks.GetPreviousChapter("Genesis", 50))
	assert.Equal(t, Chapter(-1), KJVBooks.GetPreviousChapter("Genesis", 0))
	assert.Equal(t, Chapter(-1), KJVBooks.GetPreviousChapter("genesis", 1))
	assert.Equal(t, Chapter(-1), KJVBooks.GetPreviousChapter("Genesis", 51))

	assert.Equal(t, Chapter(21), KJVBooks.GetPreviousChapter("Revelation", 22))
	assert.Equal(t, Chapter(22), KJVBooks.GetPreviousChapter("Revelation", 1))
	assert.Equal(t, Chapter(-1), KJVBooks.GetPreviousChapter("Revelation", 0))
	assert.Equal(t, Chapter(-1), KJVBooks.GetPreviousChapter("revelation", 1))
	assert.Equal(t, Chapter(-1), KJVBooks.GetPreviousChapter("Revelation", 23))
}

func TestBooks_GetNextVerse(t *testing.T) {
	// In the beginning God created the heaven and the earth.
	assert.Equal(t, Verse(1), KJVBooks.GetNextVerse("Genesis", 1, 31))
	// And the earth was without form, and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face of the waters.
	assert.Equal(t, Verse(2), KJVBooks.GetNextVerse("Genesis", 1, 1))
	// And Joseph fell upon his father's face, and wept upon him, and kissed him.
	assert.Equal(t, Verse(1), KJVBooks.GetNextVerse("Genesis", 50, 26))
	assert.Equal(t, Verse(-1), KJVBooks.GetNextVerse("Genesis", 0, 0))
	assert.Equal(t, Verse(-1), KJVBooks.GetNextVerse("genesis", 1, 1))
	assert.Equal(t, Verse(-1), KJVBooks.GetNextVerse("Genesis", 1, 0))
	assert.Equal(t, Verse(-1), KJVBooks.GetNextVerse("Genesis", 51, 1))

	// The Revelation of Jesus Christ, which God gave unto him, to shew unto his servants things which must shortly come to pass; and he sent and signified it by his angel unto his servant John:
	assert.Equal(t, Verse(1), KJVBooks.GetNextVerse("Revelation", 1, 20))
	// Who bare record of the word of God, and of the testimony of Jesus Christ, and of all things that he saw.
	assert.Equal(t, Verse(2), KJVBooks.GetNextVerse("Revelation", 1, 1))
	// And he shewed me a pure river of water of life, clear as crystal, proceeding out of the throne of God and of the Lamb.
	assert.Equal(t, Verse(1), KJVBooks.GetNextVerse("Revelation", 22, 21))
	assert.Equal(t, Verse(-1), KJVBooks.GetNextVerse("Revelation", 0, 0))
	assert.Equal(t, Verse(-1), KJVBooks.GetNextVerse("revelation", 1, 1))
	assert.Equal(t, Verse(-1), KJVBooks.GetNextVerse("Revelation", 1, 0))
	assert.Equal(t, Verse(-1), KJVBooks.GetNextVerse("Revelation", 23, 0))
}

func TestBooks_GetPreviousVerse(t *testing.T) {
	// And to every beast of the earth, and to every fowl of the air, and to every thing that creepeth upon the earth, wherein there is life, I have given every green herb for meat: and it was so.
	assert.Equal(t, Verse(30), KJVBooks.GetPreviousVerse("Genesis", 1, 31))
	// And God saw every thing that he had made, and, behold, it was very good. And the evening and the morning were the sixth day.
	assert.Equal(t, Verse(31), KJVBooks.GetPreviousVerse("Genesis", 1, 1))
	// So Joseph died, being an hundred and ten years old: and they embalmed him, and he was put in a coffin in Egypt.
	assert.Equal(t, Verse(26), KJVBooks.GetPreviousVerse("Genesis", 50, 1))
	assert.Equal(t, Verse(-1), KJVBooks.GetPreviousVerse("Genesis", 0, 0))
	assert.Equal(t, Verse(-1), KJVBooks.GetPreviousVerse("genesis", 1, 1))
	assert.Equal(t, Verse(-1), KJVBooks.GetPreviousVerse("Genesis", 1, 0))
	assert.Equal(t, Verse(-1), KJVBooks.GetPreviousVerse("Genesis", 51, 1))

	// Write the things which thou hast seen, and the things which are, and the things which shall be hereafter;
	assert.Equal(t, Verse(19), KJVBooks.GetPreviousVerse("Revelation", 1, 20))
	// The mystery of the seven stars which thou sawest in my right hand, and the seven golden candlesticks. The seven stars are the angels of the seven churches: and the seven candlesticks which thou sawest are the seven churches.
	assert.Equal(t, Verse(20), KJVBooks.GetPreviousVerse("Revelation", 1, 1))
	// The grace of our Lord Jesus Christ be with you all. Amen.
	assert.Equal(t, Verse(21), KJVBooks.GetPreviousVerse("Revelation", 22, 1))
	assert.Equal(t, Verse(-1), KJVBooks.GetPreviousVerse("Revelation", 0, 0))
	assert.Equal(t, Verse(-1), KJVBooks.GetPreviousVerse("revelation", 1, 1))
	assert.Equal(t, Verse(-1), KJVBooks.GetPreviousVerse("Revelation", 1, 0))
	assert.Equal(t, Verse(-1), KJVBooks.GetPreviousVerse("Revelation", 23, 1))
}
