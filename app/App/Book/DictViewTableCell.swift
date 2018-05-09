//
//  DictViewTableCell.swift
//  app
//
//  Created by Sunho Kim on 09/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class DictViewTableCell: UITableViewCell {
    @IBOutlet weak var backView: UIView!
    @IBOutlet weak var label: UILabel!

    override func awakeFromNib() {
        super.awakeFromNib()
        self.label.numberOfLines = 0
        UIUtill.roundView(self.backView)
    }

    override func setSelected(_ selected: Bool, animated: Bool) {
        super.setSelected(selected, animated: animated)

        // Configure the view for the selected state
    }

}
