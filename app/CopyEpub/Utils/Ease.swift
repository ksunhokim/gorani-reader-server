//
//  Util.swift
//  copyEpub
//
//  Created by sunho on 2018/05/03.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import UIKit

enum EaseType {
    case sineIn
    case sineOut
    case sineInOut
    
    case quadIn
    case quadOut
    case quadInOut
    
    case cubicIn
    case cubicOut
    case cubicInOut
    
    case quartIn
    case quartOut
    case quartInOut
    
    case quintIn
    case quintOut
    case quintInOut
    
    case expoIn
    case expoOut
    case expoInOut
    
    case circIn
    case circOut
    case circInOut
}

struct Ease {
    var timing: CAMediaTimingFunction
    
    // https://easings.net/
    init(_ type: EaseType) {
        switch type {
        case .sineIn:
            self.init(0.47,0,0.745,0.715)
        case .sineOut:
            self.init(0.39,0.575,0.565, 1)
        case .sineInOut:
            self.init(0.445, 0.05, 0.55, 0.95)
            
        case .quadIn:
            self.init(0.55, 0.085, 0.68, 0.53)
        case .quadOut:
            self.init(0.25, 0.46, 0.45, 0.94)
        case .quadInOut:
            self.init(0.455, 0.03, 0.515, 0.955)
            
        case .cubicIn:
            self.init(0.55, 0.055, 0.675, 0.19)
        case .cubicOut:
            self.init(0.215, 0.61, 0.355, 1)
        case .cubicInOut:
            self.init(0.645, 0.045, 0.355, 1)
            
        case .quartIn:
            self.init(0.895, 0.03, 0.685, 0.22)
        case .quartOut:
            self.init(0.165, 0.84, 0.44, 1)
        case .quartInOut:
            self.init(0.77, 0, 0.175, 1)
            
        case .quintIn:
            self.init(0.755, 0.05, 0.855, 0.06)
        case .quintOut:
            self.init(0.23, 1, 0.32, 1)
        case .quintInOut:
            self.init(0.86,0,0.07,1)
            
        case .expoIn:
            self.init(0.95, 0.05, 0.795, 0.035)
        case .expoOut:
            self.init(0.19, 1, 0.22, 1)
        case .expoInOut:
            self.init(1, 0, 0, 1)
            
        case .circIn:
            self.init(0.6, 0.04, 0.98, 0.335)
        case .circOut:
            self.init(0.075, 0.82, 0.165, 1)
        case .circInOut:
            self.init(0.785, 0.135, 0.15, 0.86)
        }
    }
    
    init(_ p1x: Float, _ p1y: Float, _ p2x: Float, _ p2y: Float) {
        self.timing = CAMediaTimingFunction(controlPoints: p1x, p1y, p2x, p2y)
    }
    
    static func begin(_ type: EaseType) {
        let ease = Ease(type)
        CATransaction.begin()
        CATransaction.setAnimationTimingFunction(ease.timing)
    }
    
    static func end() {
       CATransaction.commit()
    }
}

