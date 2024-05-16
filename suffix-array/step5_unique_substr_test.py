import string
import random
import os
# initializing size of string
N = 100000
 
# using random.choices()
# generating random strings

letters = string.ascii_lowercase

while True: 
    res = ''.join(random.choice(letters) for i in range(N))
    # print(f"RES={res}")
    os.system(f"echo \"{res}\" | ./main.bin")
    input()