
import Foundation
import SQLite



class Dict {
    fileprivate static let dictURL: URL = {
        let path = Bundle.main.path(forResource: "dict", ofType: "db")!
        return URL(string: path)!
    }()
    
    static let shared = Dict(url: dictURL)
    
    let connection: Connection
    
    typealias DefSortPolicy = (_ word: String, _ entries: [DictDefinition], _ pos: POS?) -> [DictDefinition]
    var defSortPolicy: DefSortPolicy?

    private init(url: URL) {
        self.connection = try! Connection(url.path)
    }
    
    func get(word: String, firstDefPos: POS? = nil) -> DictEntry? {
        return DictEntry.get(word: word, firstDefPos: firstDefPos, policy: self.defSortPolicy)
    }
    
    func search(word: String, firstWordType: VerbType? = nil, firstDefPos: POS? = nil) -> [DictEntry] {
        return DictEntry.search(word: word, firstWordType: firstWordType, firstDefPos: firstDefPos, policy: self.defSortPolicy)
    }
}
