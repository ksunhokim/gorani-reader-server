//
//  DictViewController.swift
//  app
//
//  Created by Sunho Kim on 09/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

fileprivate let MAXCHAR = 120

class DictViewController: UIViewController, UITableViewDelegate, UITableViewDataSource {
    var tableView: UITableView!
    var sentenceLabel: UILabel!
    var cancelButton: UIButton!
    
    var word: String
    var sentence: String
    var index: Int
    
    var entries: [DictEntry]
    
    override var prefersStatusBarHidden: Bool {
        return true
    }
    
    init(word: String, sentence: String, index: Int) {
        self.word = word
        self.sentence = sentence
        self.index = index
        self.entries = Dict.shared.search(word: word)
        super.init(nibName: nil, bundle: Bundle.main)
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("storyboard is not good" )
    }

    @objc func onCacnelButton(_ sender: Any? = nil) {
        dismiss(animated: true)
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        
        self.view.backgroundColor = UIColor.white
        let attrs = [NSAttributedStringKey.font : UIFont.boldSystemFont(ofSize: 20)]

        let (front, middle, end) = self.getFrontMiddleEnd()
        
        let frontString = NSMutableAttributedString(string: front)
        let middleString = NSMutableAttributedString(string: middle, attributes:attrs)
        let endString = NSMutableAttributedString(string: end)
        frontString.append(middleString)
        frontString.append(endString)
        
        self.sentenceLabel = UILabel(frame: CGRect(x: 14, y: 20, width: view.frame.width - 28, height: 70))
        self.sentenceLabel.attributedText = frontString
        self.sentenceLabel.textAlignment = .center
        self.sentenceLabel.numberOfLines = 0
        self.view.addSubview(self.sentenceLabel)
        
        let line = UIView(frame: CGRect(x: 0, y: self.sentenceLabel.frame.height + self.sentenceLabel.frame.origin.y + 20, width: view.frame.width, height: 0.7))
        line.backgroundColor = UIUtill.gray1
        self.view.addSubview(line)

        self.tableView = UITableView(frame: CGRect(x: 0, y: line.frame.origin.y + 0.7 , width: view.frame.width, height: view.frame.height - 200))
        self.tableView.delegate = self
        self.tableView.dataSource = self
        self.tableView.backgroundColor = UIUtill.lightGray1
        self.tableView.separatorStyle = .none
        self.tableView.rowHeight = UITableViewAutomaticDimension;
        self.tableView.estimatedRowHeight = 100;
        self.tableView.showsVerticalScrollIndicator = false
        self.tableView.register(UINib(nibName: "DictViewTableCell", bundle: nil), forCellReuseIdentifier: "DictViewTableCell")
        self.tableView.contentInset = UIEdgeInsets(top: 0, left: 0, bottom: 10, right: 0)
        self.view.addSubview(self.tableView)
        

        self.cancelButton = UIButton(frame: CGRect(x: 14, y: view.frame.height - 70, width: view.frame.width - 28, height: 50))
        self.cancelButton.backgroundColor = UIUtill.blue
        self.cancelButton.setTitleColor(UIUtill.white, for: .normal)
        self.cancelButton.setTitle("Cancel", for: .normal)
        self.cancelButton.addTarget(self, action: #selector(onCacnelButton(_:)), for: .touchUpInside)
        UIUtill.roundView(self.cancelButton)
        self.view.addSubview(self.cancelButton)
    }
    
    func numberOfSections(in tableView: UITableView) -> Int {
        return self.entries.count
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.entries[section].defs.count
    }
    
    func tableView(_ tableView: UITableView, heightForHeaderInSection section: Int) -> CGFloat {
        return 50
    }
    
    func tableView(_ tableView: UITableView, viewForHeaderInSection section: Int) -> UIView? {
        let view = UIView(frame: CGRect(x: 8, y: 0, width: self.tableView.bounds.width - 16, height: 50))
        view.backgroundColor = UIUtill.lightGray0
        let label = UILabel()
        label.font = UIFont.boldSystemFont(ofSize: 20)
        label.frame.origin.x = 28
        label.textColor = UIColor.black
        let entry = self.entries[section]
        label.text = entry.word
        label.sizeToFit()
        label.frame = CGRect(origin: label.frame.origin, size: CGSize(width: label.frame.width, height: 50))
        view.addSubview(label)
        
        let pronButton = UIButton()
        pronButton.setTitle(entry.pron, for: .normal)
        pronButton.contentEdgeInsets = UIEdgeInsets(top: 0, left: 20, bottom: 0, right: 20)
        pronButton.backgroundColor = UIUtill.blue
        pronButton.setTitleColor(UIUtill.white, for: .normal)
        pronButton.titleLabel?.font = UIFont.systemFont(ofSize: 10)
        pronButton.titleLabel?.baselineAdjustment = .alignCenters
        pronButton.contentVerticalAlignment = .center
        pronButton.titleLabel?.sizeToFit()
        pronButton.sizeToFit()
        pronButton.frame = CGRect(x: pronButton.frame.origin.x + pronButton.frame.width + 20, y: 5, width: pronButton.frame.width, height: 40)
        view.addSubview(pronButton)
        return view
    }
    
    fileprivate func getDictEntryColor(entry: DictEntry) -> UIColor {
        if entry is DictEntryRedirect {
            return UIUtill.green
        }
        return UIUtill.lightGray0
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = self.tableView.dequeueReusableCell(withIdentifier: "DictViewTableCell", for: indexPath) as! DictViewTableCell
        let entry = self.entries[indexPath.section].defs[indexPath.row]
        cell.backgroundColor = UIColor.clear
        cell.label.text = entry.def
        return cell
    }

    fileprivate func getFrontMiddleEnd() -> (String, String, String) {
        let arr = self.sentence.components(separatedBy: " ")
        var front = ""
        var end = ""
        if arr.count > index {
            let frontarr = arr[...(index - 1)]
            front = frontarr.joined(separator: " ")
            let endarr = arr[(index + 1)...]
            end = endarr.joined(separator: " ")
        }
        
        let candidate1 = min(front.count, MAXCHAR / 2)
        let candidate2 = min(end.count, MAXCHAR / 2)
        var frontLength = 0
        if candidate1 < candidate2 {
            frontLength = candidate1
        } else {
            frontLength = MAXCHAR - candidate2
        }
        
        front = String(front.suffix(frontLength))
        end = String(end.prefix(MAXCHAR - frontLength))
        return (front, " \(arr[index]) ", end)
    }
}
