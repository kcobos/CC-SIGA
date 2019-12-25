# -*- coding: utf-8 -*-

def check_password_strength(password):
    if len(password) <  8: # length minimum 8 chars
        return False
    uppers = [x for x in password if x >= 'A' and x <= 'Z']
    if len(uppers) < 2: # must have, at least, 2 upper chars
        return False
    numbers = [x for x in password if x >= '0' and x <= '9']
    if len(numbers) < 2: # must have, at least, 2 numbers
        return False
    specials = [x for x in password if x in '~`!@#$%^&*()-_+={}[]|\\/:;"\'<>,.?']
    if len(specials) < 1: # must have, at least, 1 special char
        return False
    return True