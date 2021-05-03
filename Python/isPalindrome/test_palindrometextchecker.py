import unittest

from palindrometextchecker import is_palindrome

class TestPalindrome(unittest.TestCase):
    
    def setUp(self):
        self.palindrome_name = 'Anna'
        self.no_palindrome_name = 'Beto'

    def test_is_palindrome(self):
        self.assertEqual(is_palindrome(self.palindrome_name), True)
        self.assertEqual(is_palindrome(self.no_palindrome_name), False)


if __name__=='__main__':
    unittest.main()
