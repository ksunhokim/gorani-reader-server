//
//  WordbookDetailViewController.swift
//  app
//
//  Created by sunho on 2018/04/27.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class WordbookDetailViewController: UIViewController, UITableViewDelegate, UITableViewDataSource {

    @IBOutlet weak var memorizeButton: UIButton!
    @IBOutlet weak var flashcardButton: UIButton!
    @IBOutlet weak var sentenceButton: UIButton!
    @IBOutlet weak var speakButton: UIButton!
    @IBOutlet weak var wordsTable: UITableView!
    @IBOutlet weak var titleLabel: UILabel!
    @IBOutlet weak var headerView: UIView!
    
    var wordbook: Wordbook!
    
    private var headerY: CGFloat!
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
        let view = UIView()
        let label = UILabel()
        label.text = self.wordbook.name
        label.sizeToFit()
        label.alpha = 0
        view.addSubview(label)
        view.frame = label.frame
        
        self.navigationItem.titleView = view
    }

    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.wordsTable.dataSource = self
        self.wordsTable.delegate = self
        

        self.titleLabel.text = self.wordbook.name

        self.layout()
    }
    
    func layout() {
        self.headerY = self.headerView.frame.minY
        self.wordsTable.contentInset = UIEdgeInsetsMake(self.headerView.frame.height, 0, 0, 0)
        roundView(self.memorizeButton)
        roundView(self.flashcardButton)
        roundView(self.speakButton)
        roundView(self.sentenceButton)
    }
    
    // header location
    func scrollViewDidScroll(_ scrollView: UIScrollView) {
        let frame = self.headerView.frame
        let y = scrollView.contentOffset.y + frame.height
        self.headerView.frame = CGRect(x: frame.minX, y: self.headerY - y, width: frame.width, height: frame.height)
        
        if let titleView = self.navigationItem.titleView {
            let textView = titleView.subviews[0]
            if y > titleLabel.frame.minY + titleLabel.frame.height {
                UIView.animate(withDuration: 0.2, animations: {
                    textView.alpha = 1
                }, completion: nil)
            } else {
                UIView.animate(withDuration: 0.2, animations: {
                    textView.alpha = 0
                }, completion: nil)
            }
        }
    }
    
    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return UITableViewAutomaticDimension
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.wordbook.words.count
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "WordsTableCell")
        let item = self.wordbook.words[indexPath.row]
        
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
