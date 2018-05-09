//
//  DictViewController.swift
//  app
//
//  Created by Sunho Kim on 09/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

fileprivate let MAXCHAR = 120

class DictViewController: UIViewController, UICollectionViewDelegate, UICollectionViewDataSource {
    var collectionView : UICollectionView!
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
        
        let line = UIView(frame: CGRect(x: 0, y: self.sentenceLabel.frame.height + self.sentenceLabel.frame.origin.y + 20, width: view.frame.width, height: 0.5))
        line.backgroundColor = UIUtill.gray
        self.view.addSubview(line)
        
        let layout: UICollectionViewFlowLayout = UICollectionViewFlowLayout()
        layout.sectionInset = UIEdgeInsets(top: 0, left: 14, bottom: 0, right: 14)
        layout.itemSize = CGSize(width: view.frame.width - 28, height: view.frame.height - 220)
        layout.minimumInteritemSpacing = 0
        layout.minimumLineSpacing = 28
        layout.scrollDirection = .horizontal
        
        self.collectionView = UICollectionView(frame: CGRect(x: 0, y: line.frame.origin.y + 0.5 , width: view.frame.width, height: view.frame.height - 190), collectionViewLayout: layout)
        self.collectionView.backgroundColor = UIColor.clear
        self.collectionView.register(DictCollectionViewCell.self, forCellWithReuseIdentifier: "DictCollectionViewCell")
        self.collectionView.delegate = self
        self.collectionView.dataSource = self
        self.collectionView.isPagingEnabled = true;
        self.collectionView.showsHorizontalScrollIndicator = false
        self.collectionView.backgroundColor = UIUtill.lightGray0
        self.view.addSubview(self.collectionView)
        
        let back = UIView(frame: CGRect(x: 0, y: self.collectionView.frame.height + self.collectionView.frame.origin.y, width: view.frame.width, height: view.frame.height))
        back.backgroundColor = UIUtill.lightGray0
        self.view.addSubview(back)
        
        self.cancelButton = UIButton(frame: CGRect(x: 14, y: view.frame.height - 70, width: view.frame.width - 28, height: 50))
        self.cancelButton.backgroundColor = UIUtill.blue
        self.cancelButton.setTitleColor(UIUtill.white, for: .normal)
        self.cancelButton.setTitle("Cancel", for: .normal)
        self.cancelButton.addTarget(self, action: #selector(onCacnelButton(_:)), for: .touchUpInside)
        UIUtill.roundView(self.cancelButton)
        self.view.addSubview(self.cancelButton)
    }
    
    func numberOfSections(in collectionView: UICollectionView) -> Int {
        return 1
    }
    
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return self.entries.count
    }
    
    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAtIndexPath indexPath: IndexPath) -> CGSize {
        return CGSize(width: UIScreen.main.bounds.width - 40, height: self.collectionView.frame.height - 40)
    }
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "DictCollectionViewCell", for: indexPath) as! DictCollectionViewCell
        let entry = self.entries[indexPath.row]
        cell.displayContent(entry: entry)
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
