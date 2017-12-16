output = ""
for i in range(0,1000):
    output += str(i) +"," + str(i*i) + "\n"

with open("theoretical.csv", 'a') as out:
    out.write(output )
