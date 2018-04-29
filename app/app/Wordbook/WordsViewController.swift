//
//  WordsViewController.swift
//  app
//
//  Created by sunho on 2018/04/30.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class WordsViewController: UIViewController {
    @IBOutlet weak var tableView: UITableView!
    
    var words: [Word]!

    private var wordsTableDelegate: WordsTableViewDelegate!
    
    override func viewDidLoad() {
        super.viewDidLoad()
 
        self.wordsTableDelegate = WordsTableViewDelegate(words: self.words)
        self.tableView.delegate = self.wordsTableDelegate
        self.tableView.dataSource = self.wordsTableDelegate
    }
}
