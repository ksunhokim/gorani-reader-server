
import Foundation
import SQLite

fileprivate let wordsTable = Table("words")
fileprivate let idField = Expression<Int64>("id")
fileprivate let wordField = Expression<String>("word")
fileprivate let pronField = Expression<String?>("pron")

class DictEntry {
    var id: Int64
    var word: String
    var pron: String
    var defs: [DictDefinition] = []
    
    init(id: Int64, word: String, pron: String) {
        self.id = id
        self.word = word
        self.pron = pron.unstressed
    }
    
    class func get(word wordstr: String, firstDefPos: POS?, policy: Dict.DefSortPolicy?) -> DictEntry? {
        let query = wordsTable.where(wordField.collate(.nocase) == wordstr)
        
        do {
            if let entry = try Dict.shared.connection.pluck(query) {
                let entry = DictEntry(id: try entry.get(idField), word: try entry.get(wordField), pron: try entry.get(pronField) ?? "")
                
                DictDefinition.fetch(entry: entry, firstPos: firstDefPos, policy: policy)
                return entry
            }
        } catch {}
        
        return nil
    }
    
    
    class func search(word: String, firstWordType: VerbType?, firstDefPos: POS?, policy: Dict.DefSortPolicy?) -> [DictEntry] {
        if word == "" {
            return []
        }
        
        var entries: [DictEntry] = []
        
        let candidates = VerbType.candidates(word: word)
        for candidate in candidates {
            if let entry = DictEntry.get(word: candidate.0, firstDefPos: firstDefPos, policy: policy) {
                let entry = DictEntryRedirect(entry: entry, type: candidate.1)
                
                if candidate.1 == firstWordType
         {
                    entries.insert(entry, at: 0)
                } else {
                    entries.append(entry)
                }
            }
        }

        if let entry = DictEntry.get(word: word, firstDefPos: firstDefPos, policy: policy) {
            entries.append(entry)
        }
        
        return entries
    }
}

class DictEntryRedirect: DictEntry {
    var verbType: VerbType?
    
    convenience init(entry: DictEntry, type: VerbType?) {
        self.init(id: entry.id, word: entry.word, pron: entry.pron, type: type)
        self.defs = entry.defs
    }
    
    init(id: Int64, word: String, pron: String, type: VerbType?) {
        self.verbType = type
        super.init(id: id, word: word, pron: pron)
    }
}


