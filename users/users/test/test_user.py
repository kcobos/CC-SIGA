# -*- coding: utf-8 -*-

import os, sys, unittest
from datetime import date, timedelta
import hashlib

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)) + '/../')

from user import User
import settings
import errors 
import utils

class TestUser(unittest.TestCase):
    def setUp(self):
        self.u = User("test", "My#PasS123")

    def tearDown(self):
        self.u = None

    def test_init(self):
        u = None
        with self.assertRaises(errors.PasswordStrengthError):
            u = User("test", "1234")
        try:
            u = User("test", "1234")
        except errors.PasswordStrengthError as e:
            self.assertEqual(e.errors, 4001)
            self.assertEqual(e.message, "Password is too weak")
        self.assertIsNone(u)

        u = User("test", "My#PasS123")
        self.assertIsInstance(u, User)
        self.assertEqual(u.username, "test")
        self.assertEqual(len(u.password), 128)
        self.assertEqual(u.password, hashlib.sha512("My#PasS123".encode('utf-8')).hexdigest().upper())
        self.assertEqual(u.recent, [])
        self.assertEqual(u.favorites, [])
        self.assertIsNone(u.expiring_date)
        self.assertEqual(u.uid, 0)
    
    def test_check_password(self):
        self.assertTrue(self.u.check_password("My#PasS123"))
        self.assertFalse(self.u.check_password("M#PasS123"))
        self.assertFalse(self.u.check_password("MyPasS123"))

    def test_change_password(self):
        self.assertFalse(self.u.change_password("My#PasS123", "My#PasS123"))
        self.assertFalse(self.u.change_password("My#PasS13", "My#PasS123New"))
        self.assertTrue(self.u.change_password("My#PasS123", "My#PasS123New"))
        self.assertEqual(self.u.password, hashlib.sha512("My#PasS123New".encode('utf-8')).hexdigest().upper())

    def test_update_uid(self):
        self.assertTrue(self.u.update_uid("ff ff ff ff"))
        self.assertEqual(self.u.uid, int("ffffffff", 16))
        self.assertTrue(self.u.update_uid("ff ff ff ff ff ff ff"))
        self.assertEqual(self.u.uid, int("ffffffffffffff", 16))
        self.assertTrue(self.u.update_uid("ff ff ff ff ff ff ff ff ff ff"))
        self.assertEqual(self.u.uid, int("ffffffffffffffffffff", 16))

        self.assertFalse(self.u.update_uid("gf ff ff ff ff ff ff ff ff ff"))
        self.assertFalse(self.u.update_uid("af"))
        self.assertFalse(self.u.update_uid("af aa aa"))

        self.assertEqual(self.u.expiring_date, date.today()+settings.uid_duration)

        settings.uid_duration = timedelta(days=365*3)
        self.assertNotEqual(self.u.expiring_date, date.today()+settings.uid_duration)
        self.assertTrue(self.u.update_uid("ff ff ff ff"))
        self.assertEqual(self.u.expiring_date, date.today()+settings.uid_duration)

    def test_add_favorite(self):
        self.assertEqual(self.u.favorites, [])
        for x in range(20):
            self.assertTrue(self.u.add_favorite(x))
            self.assertIn(x, self.u.favorites)
            self.assertEqual(len(self.u.favorites),x+1)
        self.assertFalse(self.u.add_favorite(2))
        self.assertFalse(self.u.add_favorite(10))

    def test_remove_favorite(self):
        self.assertEqual(self.u.favorites, [])
        self.u.add_favorite(2)
        self.u.add_favorite(4)
        self.assertFalse(self.u.remove_favorite(1))
        self.assertTrue(self.u.remove_favorite(2))
        self.assertEqual(len(self.u.favorites), 1)
        self.assertFalse(self.u.remove_favorite(2))
        self.assertTrue(self.u.remove_favorite(4))
        self.assertEqual(len(self.u.favorites),0)

    def test_add_recent(self):
        self.assertEqual(self.u.recent, [])
        for x in range(5):
            self.u.add_recent(x)
            self.assertIn(x, self.u.recent)
            self.assertEqual(len(self.u.recent),x+1)
        self.u.add_recent(1)
        self.assertEqual([0,2,3,4,1], self.u.recent)
        self.u.add_recent(5)
        self.assertEqual(len(self.u.recent), settings.recent_size)
        self.assertIn(5, self.u.recent)

        settings.recent_size = 7
        self.u.add_recent(0)
        self.assertEqual(len(self.u.recent), 6)
        self.assertIn(0, self.u.recent)
        
        for x in range(10, 18):
            self.u.add_recent(x)
        self.assertEqual(len(self.u.recent), settings.recent_size)
        self.assertEqual(self.u.recent, [11,12,13,14,15,16,17])
        