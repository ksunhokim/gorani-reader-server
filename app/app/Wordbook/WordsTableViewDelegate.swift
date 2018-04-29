//
//  WordsTableViewDelegate.swift
//  app
//
//  Created by sunho on 2018/04/30.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class WordsTableViewDelegate: NSObject, UITableViewDelegate, UITableViewDataSource {
    var words: [Word]
    var maximumItem: Int = 0
    init(words: [Word]) {
        self.words = words
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return maximumItem == 0 ? self.words.count : min(self.words.count, self.maximumItem)
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "WordsTableCell")
        let item = self.words[indexPath.row]
        
        cell!.textLabel!.text = item.word
        
        let correct = item.correct
        if correct > 0 {
            cell!.detailTextLabel!.textColor = UIColor(red: 0, green: 255, blue: 0, alpha: 255)
            cell!.detailTextLabel!.text = "+\(correct)"
        } else if correct < 0 {
            cell!.detailTextLabel!.textColor = UIColor(red: 255, green: 0, blue: 0, alpha: 255)
            cell!.detailTextLabel!.text = String(correct)
        }
        
        return cell!
    }
}
