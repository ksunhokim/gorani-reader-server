import Foundation
import UIKit
import FolioReaderKit

enum EpubError: Error{
    case notProper
}

class Epub {
    var book: FRBook?
    var title: String = ""
    var cover: UIImage?
    
    convenience init(bookBase: URL) throws {
        try self.init(path: bookBase.path)
    }
    
    init(path: String) throws {
        self.book = try FREpubParser().readEpub(bookBasePath: path)
        try self.parse()
    }
    
    func parse() throws {
        if let image = self.book!.coverImage {
            self.cover = UIImage(contentsOfFile: image.fullHref)
        }
        
        guard let title = self.book!.title else {
            throw EpubError.notProper
        }
        
        self.title = title
    }
    
    static func getLocalBooks() -> [Epub] {
        guard let paths = FileUtill.contentsOfDirectory(path: FileUtill.booksDir.path) else {
            assert(true)
            return []
        }
        
        var books = [Epub]()
        for path in paths{
            if let book = try? Epub(path: path) {
                books.append(book)
            } else {
                assert(true)
            }
        }
        
        return books
    }
    
    // should only be used in inheritance
    init() {
    }
}
