import threading
import os

def thread_function1():
    # os.system("./main")
    os.system("go run main.go")

def thread_function2():
    # os.system("cd app; ./main")
    os.system("cd app; go run main.go")


x = threading.Thread(target=thread_function1)
y = threading.Thread(target=thread_function2)
T = [x, y]

for t in T:
    t.start()
    
for t in T:
    t.join()
