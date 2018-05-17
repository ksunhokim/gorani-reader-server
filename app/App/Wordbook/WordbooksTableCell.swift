import UIKit

class WordbooksTableCell: UITableViewCell {
    @IBOutlet weak var back: UIView!
    @IBOutlet weak var titleLabel: UILabel!
    @IBOutlet weak var countLabel: UILabel!
    @IBOutlet weak var accuracyLabel: UILabel!
    @IBOutlet var quizIcons: [UIButton]!
    
    override func setHighlighted(_ highlighted: Bool, animated: Bool) {
        if (highlighted) {
            self.back.backgroundColor = UIUtill.lightGray0
        } else {
            self.back.backgroundColor = UIUtill.white
        }
    }
}
