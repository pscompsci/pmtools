import random
import string

initials = string.ascii_uppercase
first_names = ['Steve', 'Virat', 'Babar', 'Joe', 'Harry', 'Ron', 'Hermione']
last_names = ['Smith', 'Kohli', 'Azam', 'Root', 'Potter', 'Weasley', 'Granger']

def rand_name(a, b):
    return f'{random.choice(a)} {random.choice(b)}'