import UIKit
import SwipeCellKit

class BooksTableCell: SwipeTableViewCell {
    @IBOutlet weak var back: UIView!
    @IBOutlet weak var titleLabel: UILabel!
    @IBOutlet weak var autorLabel: UILabel!
    @IBOutlet weak var coverImage: UIImageView!
    
    override func setHighlighted(_ highlighted: Bool, animated: Bool) {
        if (highlighted) {
            self.back.backgroundColor = UIUtill.lightGray0
        } else {
            self.back.backgroundColor = UIUtill.white
        }
    }
}
