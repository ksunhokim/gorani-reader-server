//
//  WordbookDetailViewController.swift
//  app
//
//  Created by sunho on 2018/04/27.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class WordbookDetailViewController: UIViewController {

    @IBOutlet weak var bookImage: UIImageView!
    @IBOutlet weak var memorizeButton: UIButton!
    @IBOutlet weak var flashcardButton: UIButton!
    @IBOutlet weak var sentenceButton: UIButton!
    @IBOutlet weak var speakButton: UIButton!

    var item: Wordbook!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.layout()
    }
    
    func layout() {
        self.navigationItem.title = self.item.name
        roundView(self.bookImage)
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
    

}
