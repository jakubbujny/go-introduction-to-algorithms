import math
output = ""
for i in range(1,8000):
    output += str(i) +"," + str(5.5*i*math.log(i)) + "\n"

with open("theoretical.csv", 'a') as out:
    out.write(output )
