//
//  Helper.swift
//  copyEpub
//
//  Created by sunho on 2018/05/03.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import UIKit

class Spinner: UIActivityIndicatorView {
    weak var target: UIView?

    override func startAnimating() {
        self.target?.alpha = 0
        super.startAnimating()
    }
    
    override func stopAnimating() {
        self.target?.alpha = 1
        super.stopAnimating()
    }
    
    init(target: UIView) {
        super.init(activityIndicatorStyle: .gray)
        self.hidesWhenStopped = true
        self.target = target
        self.center = target.center
    }
    
    required init(coder: NSCoder) {
        super.init(coder: coder)
    }
}
