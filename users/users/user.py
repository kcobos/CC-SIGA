# -*- coding: utf-8 -*-

import utils, errors, settings
from datetime import date
import hashlib

class User:
    def __init__(self, username, password):
        self.username = username
        if not utils.check_password_strength(password):
            raise errors.PasswordStrengthError("Password is too weak", 4001)
        self.password = hashlib.sha512(password.encode('utf-8')).hexdigest().upper()
        self.uid = 0
        self.expiring_date = None
        self.favorites = list()
        self.recent = list()
        
    def check_password(self, password):
        sha = hashlib.sha512(password.encode('utf-8')).hexdigest().upper()
        if sha == self.password:
            return True
        return False
    
    def change_password(self, old, new):
        if old == new:
            return False
        sha_old = hashlib.sha512(old.encode('utf-8')).hexdigest().upper()
        if utils.check_password_strength(new) and sha_old == self.password:
            self.password = hashlib.sha512(new.encode('utf-8')).hexdigest().upper()
            return True
        return False

    def update_uid(self, uid):
        uid = uid.replace(" ", "")
        try:
            if len(uid) in (8,14,20): # 4,7,10 bytes
                if "0x"+uid == hex(int(uid,16)):
                    self.uid = int(uid,16)
                    self.expiring_date = date.today() + settings.uid_duration
                    return True
        except ValueError:
            return False
        return False

    def add_favorite(self, id_place):
        if id_place not in self.favorites:
            self.favorites.append(id_place)
            return True
        return False

    def remove_favorite(self, id_place):
        if id_place in self.favorites:
            self.favorites.remove(id_place)
            return True
        return False

    def add_recent(self, id_place):
        if id_place in self.recent:
            self.recent.remove(id_place)
        self.recent.append(id_place)
        if len(self.recent) > settings.recent_size:
            self.recent.pop(0)
