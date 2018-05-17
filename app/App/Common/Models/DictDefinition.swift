import Foundation
import SQLite

fileprivate let defsTable = Table("defs")
fileprivate let idField = Expression<Int64>("id")
fileprivate let wordIdField = Expression<Int64>("word_id")
fileprivate let posField = Expression<String?>("pos")
fileprivate let defField = Expression<String>("def")

class DictDefinition {
    var id: Int64
    var pos: POS?
    var def: String
    var examples: [DictExample] = []
    
    init(id: Int64, word: DictEntry, pos: POS?, def: String) {
        self.id = id
        self.pos = pos
        self.def = def
    }
    
    class func fetch(entry: DictEntry, firstPos pos2: POS?, policy: Dict.DefSortPolicy?) {
        let query = defsTable.where(wordIdField == entry.id)
            .order(posField, idField)
        guard let results = try? Dict.shared.connection.prepare(query) else {
            return
        }
        
        var defs: [DictDefinition] = []
        
        for result in results {
            do {
                let defi = DictDefinition(id: try result.get(idField), word: entry, pos: POS(rawValue: try result.get(posField) ?? ""), def: try result.get(defField))
                
                DictExample.fetch(def: defi)
                
                if pos2 != nil && pos2 == defi.pos {
                    defs.insert(defi, at: 0)
                } else {
                    defs.append(defi)
                }
                
            } catch{}
        }
        
        if let policy = policy {
            defs = policy(entry.word, defs, pos2)
        }
        
        entry.defs = defs
    }
}
