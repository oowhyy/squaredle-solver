import keyboard as kbd  # importing keyboard module
import time
string = "I love python! It is such an amazing language."


# Using readlines()
file1 = open('output.txt', 'r')
Lines = file1.readlines()

time.sleep(3)  # 3 second gap to avoid unwanted actions

for line in Lines:
    print(line.strip())
    kbd.write(line.strip(), 0.01)
    kbd.press('enter')
