# -*- coding: utf-8 -*-

from . import user, errors
from datetime import date

class Users:
    def __init__(self):
        self.user_list = dict()

    def add(self, user):
        if user.username not in self.user_list:
            self.user_list[user.username] = user
            return True
        return False
    
    def remove(self, user):
        if self.check_username(user.username):
            del(self.user_list[user.username])
            return True
        raise errors.UserNotExistsError("User %s not exists"%user.username, 4002)

    def check_username(self, username):
        if username in self.user_list:
            return True
        return False

    def check_uid(self, uid):
        uids = [val.uid for _, val in self.user_list.items() if val.expiring_date > date.today()]
        if int(uid.replace(" ", ""),16) in uids:
            return True
        return False

    def login(self, username, password):
        if self.check_username(username):
            user = self.user_list[username]
            if user.check_password(password):
                return True
        else:
            raise errors.UserNotExistsError("User %s not exists"%username, 4002)
        return False

