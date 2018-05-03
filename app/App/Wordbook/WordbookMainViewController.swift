//
//  WordbookMainViewController.swift
//  app
//
//  Created by sunho on 2018/04/26.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class WordbookMainViewController: UIViewController, UITableViewDataSource, UITableViewDelegate {
    @IBOutlet var tableView: UITableView!
    
    var wordbooks: [Wordbook] = [
        Wordbook(name: "asdf", words: [
            Word(word:"hello", definition: "hoi", pronunciation: "hoi", correct: 10, sourceBook: "hoi", sourceSentence: "hoi", addedDate: Date(), seenDate: Date()),
            Word(word:"hello2", definition: "hoi", pronunciation: "hoi", correct: -10, sourceBook: "hoi", sourceSentence: "hoi", addedDate: Date(), seenDate: Date())
            ],
                 addedDate: Date(), seenDate: Date(),
                 new: true
        ),
        Wordbook(name: "asdf2", words: [],
                 addedDate: Date(), seenDate: Date(),
                 new: false
                 )
    ]
    
    
    override func viewDidLoad() {
        super.viewDidLoad()

        self.tableView.tableFooterView = UIView()
        self.tableView.delegate = self
        self.tableView.dataSource = self
    }

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.wordbooks.count
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "WordbooksTableCell") as! WordbooksTableCell
        
        let item = self.wordbooks[indexPath.row]
        cell.textLabel!.text = item.name
        cell.detailTextLabel!.text = String(item.words.count)
        
        if item.new {
            cell.badge.image = UIImage(named: "circle")
        } else {
            cell.badge.image = UIImage()
        }
        
        return cell
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?)
    {
        if segue.destination is WordbookDetailViewController
        {
            let vc = segue.destination as? WordbookDetailViewController
            
            let row = self.tableView.indexPathForSelectedRow!.row
            let item = self.wordbooks[row]
            vc?.wordbook = item
        }
    }
}
