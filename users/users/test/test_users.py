# -*- coding: utf-8 -*-

import os, sys, unittest

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)) + '/../')

from users import Users
from user import User
import errors
from datetime import date, timedelta

class TestUsers(unittest.TestCase):
    def setUp(self):
        self.u1 = User("test1", "My#PasS123")
        self.u2 = User("test2", "My#PasS123")
        self.u3 = User("test3", "My#PasS123")
        self.us = Users()

    def tearDown(self):
        self.us = None

    def test_init(self):
        self.assertIsInstance(self.us, Users)
        self.assertEqual(len(self.us.user_list), 0)
    
    def test_add(self):
        self.assertTrue(self.us.add(self.u1))
        self.assertFalse(self.us.add(self.u1))
        self.assertEqual(len(self.us.user_list),1)
        self.assertTrue(self.us.add(self.u2))
        self.assertEqual(len(self.us.user_list),2)

    def test_remove(self):
        self.us.add(self.u1)
        self.us.add(self.u2)
        self.assertTrue(self.us.remove(self.u2))
        with self.assertRaises(errors.UserNotExistsError):
            self.us.remove(self.u2)
        try:
            self.us.remove(self.u2)
        except errors.UserNotExistsError as e:
            self.assertEqual(e.errors, 4002)
            self.assertEqual(e.message, "User test2 not exists")

        with self.assertRaises(errors.UserNotExistsError):
            self.us.remove(self.u3)
        try:
            self.us.remove(self.u3)
        except errors.UserNotExistsError as e:
            self.assertEqual(e.errors, 4002)
            self.assertEqual(e.message, "User test3 not exists")

        self.assertEqual(len(self.us.user_list),1)

    def test_check_username(self):
        self.us.add(self.u1)
        self.us.add(self.u2)
        self.assertTrue(self.us.check_username("test1"))
        self.assertTrue(self.us.check_username("test2"))
        self.assertFalse(self.us.check_username("test3"))

    def test_check_uid(self):
        self.u1.update_uid("fa fa fa fa")
        self.us.add(self.u1)
        self.u2.update_uid("fe fe fe fe")
        self.us.add(self.u2)
        self.assertTrue(self.us.check_uid("fa fa fa fa"))
        self.assertTrue(self.us.check_uid("fe fe fe fe"))
        self.assertFalse(self.us.check_uid("af af af af"))
        self.us.remove(self.u2)
        self.assertFalse(self.us.check_uid("fe fe fe fe"))
        self.u2.expiring_date = date.today()
        self.us.add(self.u2)
        self.assertFalse(self.us.check_uid("fe fe fe fe"))

    def test_login(self):
        self.us.add(self.u1)
        self.us.add(self.u2)
        self.assertTrue(self.us.login("test1", "My#PasS123"))
        self.assertFalse(self.us.login("test1", "My#PasS12"))
        
        with self.assertRaises(errors.UserNotExistsError):
            self.us.login("test3", "My#PasS123")
        try:
            self.us.login("test3", "My#PasS123")
        except errors.UserNotExistsError as e:
            self.assertEqual(e.errors, 4002)
            self.assertEqual(e.message, "User test3 not exists")