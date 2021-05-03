"""
Source: https://en.wikipedia.org/wiki/Palindrome#Names
"""

def is_palindrome(text: str) -> bool:
    return text.lower() == text.lower()[::-1]


if __name__=='__main__':
    text = input('Give me the text to analyze: ')
    print(f'{text} Is Palindrome?: {is_palindrome(text)}')

