import UIKit

class DictViewTableCell: UITableViewCell {
    @IBOutlet weak var backView: UIView!
    @IBOutlet weak var label: UILabel!

    override func awakeFromNib() {
        super.awakeFromNib()
        self.label.numberOfLines = 0
        UIUtill.roundView(self.backView)
    }
}
