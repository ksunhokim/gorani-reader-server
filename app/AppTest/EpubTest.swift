import XCTest
@testable import app
import FolioReaderKit

class EpubTest: XCTestCase {
    var epubPath: URL!
    
    override func setUp() {
        super.setUp()
        let path = Bundle.main.path(forResource: "alice", ofType: "epub")!
        let book = try! FREpubParser().readEpub(epubPath: path, removeEpub: false, unzipPath: FileUtill.booksDir.path)
        self.epubPath = FileUtill.booksDir.appendingPathComponent(book.name!)
    }
    
    override func tearDown() {
        super.tearDown()
    }
    
    func testLoad() {
        let epub = try! Epub(bookBase: self.epubPath)
        XCTAssert(epub.title == "Alice's Adventures in Wonderland")
        XCTAssert(epub.cover == nil)
    }
 
    func testGetLocalBooks() {
        let epubs = Epub.getLocalBooks()
        var okay = false
        for epub in epubs {
            if epub.title == "Alice's Adventures in Wonderland" &&
                epub.cover == nil {
                okay = true
            }
        }
        XCTAssert(okay)
    }

}
