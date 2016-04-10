#include <pthread.h>
#include <stdio.h>

int i = 0;

// Note the return type: void*
void* thread_1_function(){
// void *my_entry_function(void *param);
	int a = 0;
    printf("Hello from thread 1!\n");
    for (a = 0 ; a < 1000000 ; a++)
    {
    	i++;
    }
    printf("i fra thread_1: %d\n",i);
    return NULL;
}

void *thread_2_function(){
	int b = 0;
	printf("Hello from thread 2!\n");
	for (b = 0 ; b < 1000000 ; b++)
	{
		i--;
	}
	printf("i fra thread_2: %d\n",i);
	return NULL;
}



int main(void){

    pthread_t thread_1;
    pthread_t thread_2;

    pthread_create(&thread_1 , NULL , thread_1_function , NULL);
    // pthread_create(&thread0, NULL, my_entry_function, &parameter);
	// Arguments to a thread would be passed here ---------^
	pthread_create(&thread_2 , NULL , thread_2_function , NULL);

    
    pthread_join(thread_1 , NULL);
    pthread_join(thread_2 , NULL);
    printf("i fra main: %d\n",i);
    return 0;
}



// Forklaring
// http://timmurphy.org/2010/05/04/pthreads-in-c-a-minimal-working-example/




/*

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
*/


/*

#include <pthread.h>
#include <stdio.h>


void* someThreadFunction() {
	printf("Hello from a thread!\n");
	return NULL;
}




int main(void){
	pthread_t someThreadFunction;
	pthread_create(&someThread, NULL, someThreadFunction, NULL);
	//Arguments to a thread...

	pthread_join(someThread, NULL);
	printf("Hello from main!"\n)
	return 0;
	printf("Program utfort");
}

*/











// Forklaring
// http://timmurphy.org/2010/05/04/pthreads-in-c-a-minimal-working-example/




/*

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
*/


/*

#include <pthread.h>
#include <stdio.h>


void* someThreadFunction() {
	printf("Hello from a thread!\n");
	return NULL;
}




int main(void){
	pthread_t someThreadFunction;
	pthread_create(&someThread, NULL, someThreadFunction, NULL);
	//Arguments to a thread...

	pthread_join(someThread, NULL);
	printf("Hello from main!"\n)
	return 0;
	printf("Program utfort");
}

*/

