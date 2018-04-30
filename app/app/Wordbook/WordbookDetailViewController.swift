//
//  WordbookDetailViewController.swift
//  app
//
//  Created by sunho on 2018/04/27.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class WordbookDetailViewController: UIViewController {

    @IBOutlet weak var memorizeButton: UIButton!
    @IBOutlet weak var flashcardButton: UIButton!
    @IBOutlet weak var sentenceButton: UIButton!
    @IBOutlet weak var speakButton: UIButton!
    @IBOutlet weak var wordsTable: UITableView!

    var wordbook: Wordbook!
    
    private var wordsTableDelegate: WordsTableViewDelegate!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.wordsTableDelegate = WordsTableViewDelegate(words: self.wordbook.words)
        self.wordsTable.dataSource = self.wordsTableDelegate
        self.wordsTable.delegate = self.wordsTableDelegate
        self.layout()
    }
    
    func layout() {
        self.wordsTableDelegate.maximumItem = 1
        self.navigationItem.title = self.wordbook.name
        roundView(self.memorizeButton)
        roundView(self.flashcardButton)
        roundView(self.speakButton)
        roundView(self.sentenceButton)
    }

    @IBAction func openActionSheet(_ sender: Any) {
        let optionMenu = UIAlertController(title: nil, message: "Choose Option", preferredStyle: .actionSheet)
        
        let deleteAction = UIAlertAction(title: "Delete", style: .default, handler: {
            (alert: UIAlertAction!) -> Void in
        })
        
        let saveAction = UIAlertAction(title: "Save", style: .default, handler: {
            (alert: UIAlertAction!) -> Void in
        })
        
        let cancelAction = UIAlertAction(title: "Cancel", style: .cancel, handler: {
            (alert: UIAlertAction!) -> Void in
        })

        optionMenu.addAction(deleteAction)
        optionMenu.addAction(saveAction)
        optionMenu.addAction(cancelAction)
        
        self.present(optionMenu, animated: true, completion: nil)
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?)
    {
        if segue.destination is WordsViewController
        {
            let vc = segue.destination as? WordsViewController
            vc?.words = self.wordbook.words
        }
    }
}
