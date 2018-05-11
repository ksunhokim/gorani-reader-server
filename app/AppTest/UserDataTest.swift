import XCTest
@testable import app

class UserDataTest: XCTestCase {
    override class func setUp() {
        super.setUp()
        try? FileManager.default.removeItem(at: FileUtill.userDataURL)
    }
    
    override func tearDown() {
        super.tearDown()
    }
    
    func testKnownWord() {
        let word = KnownWord.get(word: "142sdf089hyxcsv")
        XCTAssert(word == nil)
        
        let word2 = KnownWord(word: "hello")
        try! word2.add()
        
        let word3 = KnownWord.get(word: "hello")!
        XCTAssert(word3.word == "hello")
        
        try! word3.delete()
        let word4 = KnownWord.get(word: "hello")
        XCTAssert(word4 == nil)
    }
    
    func testKnownWordFromHTML() {
        let html = """
            <html>
                <body>
                    <p>helloo <chunk>from</chunk> the other side</p>
                </body>
            </html>
        """
        try! KnownWord.add(html: html)
        let word = KnownWord.get(word: "helloo")
        XCTAssertNotNil(word)
        
        let word2 = KnownWord.get(word: "from")
        XCTAssertNil(word2)

        let word3 = KnownWord.get(word: "side")
        XCTAssertNotNil(word3)
    }

    func testKnownWordFromSpecialHTML() {
        let html = """
            <html>
                <body>
                    <p>helloo, <chunk>from.</chunk> !the other-side.</p><p>...</p>
                </body>
            </html>
        """
        try! KnownWord.add(html: html)
        let word = KnownWord.get(word: "helloo")
        XCTAssertNotNil(word)
        
        let word2 = KnownWord.get(word: "from")
        XCTAssertNil(word2)
        
        let word3 = KnownWord.get(word: "other-side")
        XCTAssertNotNil(word3)
        
    }
    
    func testWordbook() {
        let wordbook = Wordbook(name: "test")
        try! wordbook.add()
        let list = Wordbook.get()
        var pass = false
        for book in list {
            if book.name == "test" {
                pass = true
            }
        }
        XCTAssertTrue(pass)
    }
    
    func testUnknownWord() {
        let wordbook = Wordbook.unknown

        let word = WordbookEntry(sentence: "I have a pen.", index: 3, word: "pen", def: "hohoho", book: "hohoho2")
        try! wordbook.addEntry(word)

        let word2 = WordbookEntry(sentence: "I have a pen2.", index: 3, word: "pen2", def: "hohoho", book: "hohoho")
        try! wordbook.addEntry(word2)

        let wordF = wordbook.entries
        XCTAssertEqual(wordF[0].word, "pen2")
        XCTAssertEqual(wordF[1].word, "pen")
        
        let wordF2 = wordbook.filteredEntries([.book("hohoho2")])
        XCTAssertEqual(wordF2[0].word, "pen")

        let wordF3 = wordbook.filteredEntries([.day(Date().dayRange)])
        XCTAssertEqual(wordF3[0].word, "pen2")

        let wordF4 = wordbook.filteredEntries([.day(Date().yesterday.dayRange)])
        XCTAssertEqual(wordF4.count, 0)

        let wordF5 = wordbook.filteredEntries([.word("pen")])
        XCTAssertEqual(wordF5[0].word, "pen2")
        XCTAssertEqual(wordF5[1].word, "pen")
        
        try! wordbook.deleteEntry(at: 0)
        
        XCTAssertEqual(wordbook.count, 1)
        XCTAssertEqual(wordbook.entries[0].word, "pen")
    }
    
}
