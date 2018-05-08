//
//  DictViewController.swift
//  app
//
//  Created by Sunho Kim on 09/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class DictViewController: UIViewController {
    var sentenceLabel: UILabel!
    var cancelButton: UIButton!
    
    var word: String
    var sentence: String
    var index: Int
    init(word: String, sentence: String, index: Int) {
        self.word = word
        self.sentence = sentence
        self.index = index
        super.init(nibName: nil, bundle: Bundle.main)
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("storyboard is not good" )
    }

    private func getFrontMiddleEnd() -> (String, String, String) {
        let arr = self.sentence.components(separatedBy: " ")
        var front = ""
        var end = ""
        if arr.count > index {
            let frontarr = arr[...(index - 1)]
            front = frontarr.joined(separator: " ")
            let endarr = arr[(index + 1)...]
            end = endarr.joined(separator: " ")
        }
        front = String(front.suffix(60))
        end = String(end.prefix(60))
        return (front, " \(arr[index]) ", end)
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
        
        self.sentenceLabel = UILabel(frame: CGRect(x: 20, y: 40, width: view.frame.width - 40, height: 70))
        self.sentenceLabel.attributedText = frontString
        self.sentenceLabel.textAlignment = .center
        self.sentenceLabel.numberOfLines = 0
        self.view.addSubview(self.sentenceLabel)
        
        self.cancelButton = UIButton(frame: CGRect(x: 20, y: view.frame.height - 80, width: view.frame.width - 40, height: 60))
        self.cancelButton.backgroundColor = UIColor(rgba: "#007AFF")
        self.cancelButton.tintColor = UIColor.white
        self.cancelButton.setTitle("Cancel", for: .normal)
        self.cancelButton.addTarget(self, action: #selector(onCacnelButton(_:)), for: .touchUpInside)
        roundView(self.cancelButton)
        self.view.addSubview(self.cancelButton)

    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
}
