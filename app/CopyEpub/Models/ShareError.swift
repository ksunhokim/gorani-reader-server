//
//  ShareError.swift
//  copyEpub
//
//  Created by sunho on 2018/05/03.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation


enum ShareError: Error {
    case notURL
    case notNew
    case notProperEpub
    case system
}
