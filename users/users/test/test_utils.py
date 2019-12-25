# -*- coding: utf-8 -*-

import os, sys, unittest

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)) + '/../')

import utils

class TestUtils(unittest.TestCase):
    def test_check_password_strength(self):
        self.assertFalse(utils.check_password_strength("ACB@123"))
        self.assertFalse(utils.check_password_strength("hello123"))
        self.assertFalse(utils.check_password_strength("heLLO123"))
        self.assertFalse(utils.check_password_strength("hello@BYE"))
        self.assertTrue(utils.check_password_strength("My#PasS123"))