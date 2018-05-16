import Foundation
import UIKit

class UIUtill {
    class var green: UIColor {
        return UIColor(rgba: "#4CD964")
    }
    
    class var black: UIColor {
        return UIColor.black
    }
    
    class var white: UIColor {
        return UIColor.white
    }
    
    class var tint: UIColor {
        return UIColor(rgba: "#9B5FED")
    }
    
    class var gray2: UIColor {
        return UIColor(rgba: "#484848")
    }
    
    class var gray1: UIColor {
        return UIColor(rgba: "#858787")
    }
    
    class var gray: UIColor {
        return UIColor(rgba: "#BFBFC3")
    }
    
    class var lightGray1: UIColor {
        return UIColor(rgba: "#E3E3E2")
    }
    
    class var lightGray0: UIColor {
        return UIColor(rgba: "#F0F0F0")
    }
    
    class var blue: UIColor {
        return UIColor(rgba: "#006FFF")
    }
    
    class func roundView(_ view: UIView, _ radius: CGFloat = 10) {
        view.layer.cornerRadius = radius
        view.clipsToBounds = true
    }
    
    class func dropShadow(_ view: UIView, offset: CGSize, radius: CGFloat) {
        view.layer.masksToBounds = false
        view.layer.shadowColor = UIColor.black.cgColor
        view.layer.shadowOpacity = 0.15
        view.layer.shadowOffset = offset
        view.layer.shadowRadius = radius
        
        view.layer.shadowPath = UIBezierPath(rect: view.bounds).cgPath
    }
}
