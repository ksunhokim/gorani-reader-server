//
//  BookMainCell.swift
//  app
//
//  Created by sunho on 2018/05/04.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class BooksTableCell: UITableViewCell {
    @IBOutlet weak var titleLabel: UILabel!
    @IBOutlet weak var autorLabel: UILabel!
    @IBOutlet weak var coverImage: UIImageView!
    
    override func awakeFromNib() {
        super.awakeFromNib()
        // Initialization code
    }

    override func setSelected(_ selected: Bool, animated: Bool) {
        super.setSelected(selected, animated: animated)

        // Configure the view for the selected state
    }

}
