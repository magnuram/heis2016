#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

int i = 0;
pthread_mutex_t mtx;
//pthread_mutex_t mtx = PTHREAD_MUTEX_INITIALIZER;


// Note the return type: void*
void* Thread1(){
    pthread_mutex_lock(&mtx);

    for (int x = 0; x < 1000000; ++x)
	{
		i = i+1;
	}
    printf("i fra trad_1: %d\n", i );
    pthread_mutex_unlock(&mtx);

    return NULL;
}


void* Thread2(){
    pthread_mutex_lock(&mtx);

	for (int x = 0; x < 1000000; ++x)
	{
		i = i-1;
	}
    printf("i fra trad_2: %d\n", i );
    pthread_mutex_unlock(&mtx);

    return NULL;
}



int main(void){

    pthread_mutex_init(&mtx, NULL);

    pthread_t thread1;
    pthread_t thread2;

    pthread_create(&thread1, NULL, Thread1, NULL);
    pthread_create(&thread2, NULL, Thread2, NULL);

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    printf("\nIn main, threads done, i = %d\n\n",i);
    
    return 0;
    
}
//cc -o c program.c && ./c










/*

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
//#include <time.h>


int i = 0;
pthread_mutex_t mtx;
//pthread_mutex_t mtx = PTHREAD_MUTEX_INITIALIZER;


// Note the return type: void*
void* Thread1(){
    pthread_mutex_lock(&mtx);

    for (int x = 0; x < 1000000; ++x)
    {
        i = i+1;
    }
    printf("i fra trad_1: %d\n", i );
    pthread_mutex_unlock(&mtx);
    pthread_mutex_destroy(&mtx);


    return NULL;
}


void* Thread2(){
    pthread_mutex_lock(&mtx);

    for (int x = 0; x < 1000000; ++x)
    {
        i = i-1;
    }
    printf("i fra trad_2: %d\n", i );
    pthread_mutex_unlock(&mtx);
    pthread_mutex_destroy(&mtx);


    return NULL;
}



int main(void){

    pthread_mutex_init(&mtx, NULL);

    pthread_t thread1;
    pthread_t thread2;

    pthread_create(&thread1, NULL, Thread1, NULL);
    pthread_create(&thread2, NULL, Thread2, NULL);

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    
    return 0;
    
}
//cc -o c program.c && ./c

*/
