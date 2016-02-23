
from threading import Thread, Lock

i = 0
mtx = Lock()

def thread_1():
    mtx.acquire()
    global i
    for x in range (0,1000000):
        i= i+1
    mtx.release()

def thread_2():
    mtx.acquire()
    global i
    for x in range(0,1000000):
        i = i - 1
    mtx.release()
def main():
    global i
    Thread1 = Thread(target= thread_1, args= (),)
    Thread2 = Thread(target= thread_2,args= (),)

    Thread1.start()
    Thread2.start()

    Thread1.join()
    Thread2.join()
    print "\n Value of i after locking and unlocking threads: " , i , "\n"


main()