# -*- coding: utf-8 -*-

class PasswordStrengthError(Exception):
    def __init__(self, message, errors):
        self.message = message
        self.errors = errors

class UserNotExistsError(Exception):
    def __init__(self, message, errors):
        self.message = message
        self.errors = errors