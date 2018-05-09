//
//  DictCollectionViewCell.swift
//  app
//
//  Created by Sunho Kim on 09/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class DictCollectionViewCell: UICollectionViewCell {
    var header: UIButton!
    
    var entry: DictEntry?

    override init(frame: CGRect) {
        super.init(frame: frame)
        self.initialization()
    }
    
    func displayContent(entry: DictEntry) {
        self.entry = entry
        self.header.setTitle(entry.word, for: .normal)
    }
    
    fileprivate func initialization() {
        self.backgroundColor = UIUtill.white
        self.contentView.backgroundColor = UIUtill.white
        self.header = UIButton(frame: CGRect(x: 0, y: 0, width: frame.width, height: 50))
        self.header.setTitle("Word", for: .normal)
        self.header.backgroundColor = UIColor.clear
        self.header.setTitleColor(UIUtill.black, for: .normal)
        self.contentView.addSubview(header)
    }

    required init?(coder aDecoder: NSCoder) {
        fatalError("storyboard is not good")
    }
    
}
