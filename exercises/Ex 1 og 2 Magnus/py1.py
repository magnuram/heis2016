from threading import Thread


def Thread_1_function():
	global i
	print 'Hei fra Thread_1\n'
	a = 0
	for a in range(0,1000000):
		i = i + 1
	print 'i fra thread1 = ' , i
	print '\n\n'
	return None

def Thread_2_function():
	global i
	print 'Hei fra Thread_2\n'
	a = 0
	for a in range(0,1000000):
		i = i - 1
	print 'i fra thread2 = ' , i
	print '\n\n'
	return None





i = 0

thread1 = Thread(target = Thread_1_function, args = (),)
thread2 = Thread(target = Thread_2_function, args = (),)

thread1.start()
thread2.start()


thread1.join()
thread2.join()
print 'Trader ferdige, i = ' , i







'''
if __name__ == '__main__':
	main()
'''


'''
# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread

i = 0

def someThreadFunction():
    print("Hello from a thread!")

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
    global i


def main():
    someThread = Thread(target = someThreadFunction, args = (),)
    someThread.start()
    
    someThread.join()
    print("Hello from main!")


main()
'''







'''

main:
    global shared int i = 0
    spawn thread_1
    spawn thread_2
    join all threads
    print i

thread_1:
    do 1_000_000 times:
        i++
thread_2:
    do 1_000_000 times:
        i--

'''