import Foundation
import SQLite

fileprivate let examplesTable = Table("examples")
fileprivate let defIdField = Expression<Int64>("def_id")
fileprivate let firstField = Expression<String>("first")
fileprivate let secondField = Expression<String>("second")

class DictExample {
    var first: String
    var second: String
    
    init(first: String, second: String) {
        self.first = first
        self.second = second
    }

    class func fetch(def: DictDefinition) {
        let query = examplesTable.where(defIdField == def.id)
        guard let results = try? Dict.shared.connection.prepare(query) else {
            return
        }
        
        for result in results {
            do {
                def.examples.append(DictExample(first: try result.get(firstField), second: try result.get(secondField)))
            } catch {}
        }
    }
}
