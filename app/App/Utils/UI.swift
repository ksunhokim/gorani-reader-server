//
//  Util.swift
//  app
//
//  Created by sunho on 2018/04/29.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import UIKit

func roundView(_ view: UIView, _ radius: CGFloat = 10) {
    view.layer.cornerRadius = radius
    view.clipsToBounds = true
}
